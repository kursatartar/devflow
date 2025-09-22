package handlers

import "devflow/internal/services"

var authService *services.AuthService

func InitAuthService(s *services.AuthService) {
	authService = s
}
