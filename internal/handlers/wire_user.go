package handlers

import "devflow/internal/interfaces"

func InitUserService(s interfaces.UserService) {
	userService = s
}
