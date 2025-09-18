package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"strings"
	"time"
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
	now := time.Now()
	u := &models.User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Profile:      models.Profile{FirstName: firstName, LastName: lastName, AvatarURL: avatarURL},
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if !strings.Contains(u.Email, "@") {
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
	return s.repo.GetByID(context.Background(), id)
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
