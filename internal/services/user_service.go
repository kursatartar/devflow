package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserManager struct {
	repo interfaces.UserRepository
}

func (s *UserManager) FilterUsersByRole(role string) []*models.User {
	users, _ := s.repo.FilterByRole(context.Background(), role)
	return users
}

func NewUserService(repo interfaces.UserRepository) *UserManager {
	return &UserManager{repo: repo}
}

func (s *UserManager) CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) (*models.User, error) {
	profile := models.Profile{FirstName: firstName, LastName: lastName, AvatarURL: avatarURL}
	u, err := models.NewUser(id, username, email, passwordHash, role, profile)
	if err != nil {
		return nil, err
	}
	if ok, err := u.IsEmailValid(); err != nil || !ok {
		if err != nil {
			return nil, err
		}
		return nil, ErrInvalidEmail
	}
	if _, err := s.repo.Create(context.Background(), u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserManager) ListUsers() []*models.User {
	us, _ := s.repo.List(context.Background())
	return us
}

func (s *UserManager) GetUser(id string) (*models.User, error) {
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserManager) UpdateUser(
	id string,
	username, email, passwordHash, role,
	firstName, lastName, avatarURL *string,
) error {
	v := func(p *string) string {
		if p == nil {
			return ""
		}
		return *p
	}
	if email != nil && *email != "" {
		tmp := &models.User{Email: *email}
		if ok, err := tmp.IsEmailValid(); err != nil || !ok {
			if err != nil {
				return err
			}
			return ErrInvalidEmail
		}
	}
	if err := s.repo.UpdateCore(context.Background(), id, v(username), v(email), v(passwordHash), v(role)); err != nil {
		return err
	}
	return s.repo.UpdateProfile(context.Background(), id, models.Profile{
		FirstName: v(firstName), LastName: v(lastName), AvatarURL: v(avatarURL),
	})
}

func (s *UserManager) DeleteUser(id string) error {
	return s.repo.Delete(context.Background(), id)
}

// Auth
func (s *UserManager) Register(username, email, password, firstName, lastName, avatarURL string) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return s.CreateUser("", username, email, string(hash), "user", firstName, lastName, avatarURL)
}

func (s *UserManager) Authenticate(identifier, password string) (*models.User, error) {
	var u *models.User
	var err error
	if strings.Contains(identifier, "@") {
		u, err = s.repo.GetByEmail(context.Background(), identifier)
	} else {
		u, err = s.repo.GetByUsername(context.Background(), identifier)
	}
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return u, nil
}
