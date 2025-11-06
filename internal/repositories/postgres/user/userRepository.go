package user

import (
	"fmt"
	"gorm.io/gorm"
	"service-register/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) (r *UserRepository) {
	r = &UserRepository{}
	r.db = db
	return
}

func (r *UserRepository) CreateService(serviceModel *models.ServiceModel) error {
	result := r.db.Create(serviceModel)
	return result.Error
}

func (r *UserRepository) CreateMethod(method *models.Method) error {
	var service models.ServiceModel
	if err := r.db.First(&service, method.ServiceModelID).Error; err != nil {
		return fmt.Errorf("service with ID %d not found: %w", method.ServiceModelID, err)
	}

	result := r.db.Create(method)
	return result.Error
}

func (r *UserRepository) CreateArgument(argument *models.Argument) error {
	var method models.Method
	if err := r.db.First(&method, argument.MethodID).Error; err != nil {
		return fmt.Errorf("method with ID %d not found: %w", argument.MethodID, err)
	}

	result := r.db.Create(argument)
	return result.Error
}

func (r *UserRepository) UpdateService(serviceID uint, updates map[string]interface{}) error {
	result := r.db.Model(&models.ServiceModel{}).Where("id = ?", serviceID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no service found with ID %d to update", serviceID)
	}
	return nil
}

func (r *UserRepository) UpdateMethod(methodID uint, updates map[string]interface{}) error {
	result := r.db.Model(&models.Method{}).Where("id = ?", methodID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no method found with ID %d to update", methodID)
	}
	return nil
}

func (r *UserRepository) UpdateArgument(argumentID uint, updates map[string]interface{}) error {
	result := r.db.Model(&models.Argument{}).Where("id = ?", argumentID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no argument found with ID %d to update", argumentID)
	}
	return nil
}

func (r *UserRepository) DeleteService(serviceID uint) error {
	result := r.db.Delete(&models.ServiceModel{}, serviceID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no service found with ID %d to delete", serviceID)
	}
	return nil
}

func (r *UserRepository) DeleteMethod(methodID uint) error {
	result := r.db.Delete(&models.Method{}, methodID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no method found with ID %d to delete", methodID)
	}
	return nil
}

func (r *UserRepository) DeleteArgument(argumentID uint) error {
	result := r.db.Delete(&models.Argument{}, argumentID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no argument found with ID %d to delete", argumentID)
	}
	return nil
}

func (r *UserRepository) GetAllServices() ([]models.ServiceModel, error) {
	var services []models.ServiceModel

	err := r.db.Preload("Methods.Arguments").Find(&services).Error
	return services, err
}

func (r *UserRepository) GetServiceByID(serviceID uint) (*models.ServiceModel, error) {
	var service models.ServiceModel

	err := r.db.Preload("Methods.Arguments").First(&service, serviceID).Error
	return &service, err
}
