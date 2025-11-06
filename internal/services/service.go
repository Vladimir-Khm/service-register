package services

import (
	"service-register/internal/models"
)

type UserService interface {
	CreateService(model *models.ServiceModel) error
	CreateMethod(model *models.Method) error
	CreateArgument(model *models.Argument) error

	UpdateService(ID uint, model *models.UpdateServiceDTO) error
	UpdateMethod(ID uint, model *models.UpdateMethodDTO) error
	UpdateArgument(ID uint, model *models.UpdateArgumentDTO) error

	DeleteService(ID uint) error
	DeleteMethod(ID uint) error
	DeleteArgument(ID uint) error

	GetAllServices() ([]models.ServiceModel, error)
	GetServiceByID(id uint) (*models.ServiceModel, error)
}

type Service struct {
	User UserService
}
