package handlers

import "devflow/internal/interfaces"

var taskService interfaces.TaskService

func InitTaskService(s interfaces.TaskService) {
	taskService = s
}
