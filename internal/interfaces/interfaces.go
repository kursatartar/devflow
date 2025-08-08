package interfaces

import "devflow/internal/models"

type UserService interface {
	CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) error
	ListUsers()
	UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL string) error
	DeleteUser(id string)
	FilterUsersByRole(role string) []*models.User
}

type ProjectService interface {
	CreateProject(id, name, description, ownerID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) error
	UpdateProject(id, name, description, status string, teamMembers []string, isPrivate bool, taskWorkflow []string)
	DeleteProject(id string)
	ListProjects()
	FilterProjectsByOwner(ownerID string) []*models.Project
}

type TaskService interface {
	CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) error
	UpdateTask(id, newTitle, newDescription, newStatus, newPriority, newDueDate string, newLabels []string, newEstimated, newLogged float64)
	DeleteTask(id string)
	ListTasks()
	FilterTasksByProject(projectID string) []*models.Task
}
