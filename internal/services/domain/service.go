package domain

import (
	"log/slog"

	"service-register/internal/config"
	"service-register/internal/repositories"
	"service-register/internal/services"
	"service-register/internal/services/domain/user"
)

func CreateService(config *config.Config, logger slog.Logger, repositories *repositories.Repository) *services.Service {
	userService := user.CreateUserService(config, logger, repositories.User)
	return &services.Service{
		User: userService,
	}
}
