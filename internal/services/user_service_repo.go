package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
)

type userServiceWithRepo struct {
	repo interfaces.UserRepository
}

func NewUserServiceWithRepo(repo interfaces.UserRepository) interfaces.UserService {
	return &userServiceWithRepo{repo: repo}
}

func (s *userServiceWithRepo) CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) (*models.User, error) {

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

func (s *userServiceWithRepo) ListUsers() []*models.User {
	users, _ := s.repo.List(context.Background())
	return users
}

func (s *userServiceWithRepo) FilterUsersByRole(role string) []*models.User {
	out, _ := s.repo.FilterByRole(context.Background(), role)
	return out
}

func (s *userServiceWithRepo) UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL string) error {
	if newEmail != "" {
		if u, _ := s.repo.GetByEmail(context.Background(), newEmail); u != nil && u.ID != id {
			return ErrEmailExists
		}
		temp := &models.User{Email: newEmail}
		if ok, err := temp.IsEmailValid(); err != nil || !ok {
			if err != nil {
				return err
			}
			return ErrInvalidEmail
		}
	}
	if err := s.repo.UpdateCore(context.Background(), id, newUsername, newEmail, newPasswordHash, newRole); err != nil {
		return err
	}
	return s.repo.UpdateProfile(context.Background(), id, models.Profile{
		FirstName: newFirstName,
		LastName:  newLastName,
		AvatarURL: newAvatarURL,
	})
}

func (s *userServiceWithRepo) DeleteUser(id string) error {
	return s.repo.Delete(context.Background(), id)
}

func (s *userServiceWithRepo) GetUser(id string) (*models.User, error) {
	u, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}
