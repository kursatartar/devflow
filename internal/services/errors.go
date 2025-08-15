package services

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrEmailExists       = errors.New("email already exists")

	ErrProjectNotFound = errors.New("project not found")
	ErrProjectExists   = errors.New("project already exists")

	ErrTaskNotFound   = errors.New("task not found")
	ErrTaskExists     = errors.New("task already exists")
	ErrInvalidDueDate = errors.New("invalid due date format")

	ErrTeamNotFound = errors.New("team not found")
	ErrTeamExists   = errors.New("team already exists")
)
