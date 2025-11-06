package repositories

import "service-register/internal/models"

type Repository struct {
	User UserRepository
}

type UserRepository interface {
	CreateService(model *models.ServiceModel) error
	CreateMethod(model *models.Method) error
	CreateArgument(model *models.Argument) error

	UpdateService(ID uint, updates map[string]interface{}) error
	UpdateMethod(ID uint, updates map[string]interface{}) error
	UpdateArgument(ID uint, updates map[string]interface{}) error

	DeleteService(ID uint) error
	DeleteMethod(ID uint) error
	DeleteArgument(ID uint) error

	GetAllServices() ([]models.ServiceModel, error)
	GetServiceByID(ID uint) (*models.ServiceModel, error)
}
