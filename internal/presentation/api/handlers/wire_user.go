package handlers

import "devflow/internal/interfaces"

var userService interfaces.UserService

func InitUserService(s interfaces.UserService) {
	userService = s
}
