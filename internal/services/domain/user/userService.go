package user

import (
	"fmt"
	"log/slog"
	"service-register/internal/config"
	"service-register/internal/models"
	"service-register/internal/repositories"
)

type UserService struct {
	config         *config.Config
	logger         slog.Logger
	userRepository repositories.UserRepository
}

func (s *UserService) GetAllServices() ([]models.ServiceModel, error) {
	return s.userRepository.GetAllServices()
}

func (s *UserService) GetServiceByID(id uint) (*models.ServiceModel, error) {
	return s.userRepository.GetServiceByID(id)
}

func (s *UserService) DeleteService(ID uint) error {
	return s.userRepository.DeleteService(ID)
}

func (s *UserService) DeleteMethod(ID uint) error {
	return s.userRepository.DeleteMethod(ID)
}

func (s *UserService) DeleteArgument(ID uint) error {
	return s.userRepository.DeleteArgument(ID)
}

func (s *UserService) UpdateService(ID uint, req *models.UpdateServiceDTO) error {
	updates := make(map[string]interface{})
	if req.ServiceName != nil {
		updates["service_name"] = *req.ServiceName
	}

	if len(updates) == 0 {
		return fmt.Errorf("there is no fields to update")
	}

	return s.userRepository.UpdateService(ID, updates)
}

func (s *UserService) UpdateMethod(ID uint, req *models.UpdateMethodDTO) error {
	updates := make(map[string]interface{})
	if req.MethodName != nil {
		updates["method_name"] = *req.MethodName
	}
	if req.IsPrivate != nil {
		updates["is_private"] = *req.IsPrivate
	}
	if req.Price != nil {
		updates["price"] = *req.Price
	}

	if len(updates) == 0 {
		return fmt.Errorf("there is no fields to update")
	}

	return s.userRepository.UpdateMethod(ID, updates)
}

func (s *UserService) UpdateArgument(ID uint, req *models.UpdateArgumentDTO) error {
	updates := make(map[string]interface{})
	if req.ArgumentNumber != nil {
		updates["argument_number"] = *req.ArgumentNumber
	}
	if req.ArgumentName != nil {
		updates["argument_name"] = *req.ArgumentName
	}
	if req.ArgumentType != nil {
		updates["argument_type"] = *req.ArgumentType
	}
	if req.IsRequired != nil {
		updates["is_required"] = *req.IsRequired
	}

	if len(updates) == 0 {
		return fmt.Errorf("there is no fields to update")
	}

	return s.userRepository.UpdateArgument(ID, updates)
}

func (s *UserService) CreateMethod(method *models.Method) error {
	return s.userRepository.CreateMethod(method)
}

func (s *UserService) CreateArgument(argument *models.Argument) error {
	return s.userRepository.CreateArgument(argument)
}

func (s *UserService) CreateService(serviceModel *models.ServiceModel) error {
	return s.userRepository.CreateService(serviceModel)
}

func CreateUserService(config *config.Config, logger slog.Logger, userRepository repositories.UserRepository) (s *UserService) {
	s = &UserService{}
	s.config = config
	s.logger = *logger.With()
	s.userRepository = userRepository
	return
}
