package handlers

import "devflow/internal/interfaces"

var projectService interfaces.ProjectService

func InitProjectService(s interfaces.ProjectService) {
	projectService = s
}
