package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
)

type UserManager struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {

	return &UserManager{repo}
}

func (s *UserManager) CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) (*models.User, error) {
	if existing, _ := s.repo.GetByEmail(context.Background(), email); existing != nil {
		return nil, ErrEmailExists
	}

	user, err := models.NewUser(id, username, email, passwordHash, role, models.Profile{
		FirstName: firstName, LastName: lastName, AvatarURL: avatarURL,
	})
	if err != nil {
		return nil, err
	}
	ok, err := user.IsEmailValid()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrInvalidEmail
	}
	_, err = s.repo.Create(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserManager) ListUsers() []*models.User {
	list, err := s.repo.List(context.Background())
	if err != nil {
		return nil
	}
	return list
}

func (s *UserManager) FilterUsersByRole(role string) []*models.User {
	out, _ := s.repo.FilterByRole(context.Background(), role)
	return out
}

func (s *UserManager) UpdateUser(
	id string,
	newUsername, newEmail, newPasswordHash, newRole,
	newFirstName, newLastName, newAvatarURL *string,
) error {

	v := func(p *string) string {
		if p == nil {
			return ""
		}
		return *p
	}

	if newEmail != nil && *newEmail != "" {
		if u, _ := s.repo.GetByEmail(context.Background(), *newEmail); u != nil && u.ID != id {

			return ErrEmailExists
		}
		temp := &models.User{Email: *newEmail}
		if ok, err := temp.IsEmailValid(); err != nil || !ok {
			if err != nil {
				return err
			}
			return ErrInvalidEmail
		}
	}
	if err := s.repo.UpdateCore(context.Background(), id, v(newUsername), v(newEmail), v(newPasswordHash), v(newRole)); err != nil {
		return err
	}
	return s.repo.UpdateProfile(context.Background(), id, models.Profile{
		FirstName: v(newFirstName),
		LastName:  v(newLastName),
		AvatarURL: v(newAvatarURL),
	})
}

func (s *UserManager) DeleteUser(id string) error {
	return s.repo.Delete(context.Background(), id)
}

func (s *UserManager) GetUser(id string) (*models.User, error) {
	u, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}

var _ interfaces.UserService = (*UserManager)(nil)
