package interfaces

import "devflow/internal/models"

type UserService interface {
	CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) (*models.User, error)
	ListUsers() []*models.User
	UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL string) error
	DeleteUser(id string) error
	FilterUsersByRole(role string) []*models.User
}

type ProjectService interface {
	CreateProject(id, name, description, ownerID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error)
	UpdateProject(id, name, description, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error)
	DeleteProject(id string) error
	ListProjects() []*models.Project
	FilterProjectsByOwner(ownerID string) []*models.Project
}

type TaskService interface {
	CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) (*models.Task, error)
	UpdateTask(id string, title, description, status, priority, dueDate *string, labels *[]string, estimated, logged *float64) (*models.Task, error)
	DeleteTask(id string) error
	ListTasks() []*models.Task
	FilterTasksByProject(projectID string) []*models.Task
}

type TeamService interface {
	CreateTeam(id, name, description, ownerID string, members []models.TeamMember, settings models.TeamSettings) (*models.Team, error)
	GetTeam(id string) (*models.Team, error)
	UpdateTeam(id string, name, description *string, settings *models.TeamSettings) (*models.Team, error)
	AddMember(teamID, userID, role string) (*models.Team, error)
	RemoveMember(teamID, userID string) (*models.Team, error)
	ChangeMemberRole(teamID, userID, role string) (*models.Team, error)
	ListTeams() []*models.Team
	FilterTeamsByOwner(ownerID string) []*models.Team
	DeleteTeam(id string) error
}
