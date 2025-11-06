package user

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"service-register/internal/models"
	"service-register/internal/services"
	"strconv"
)

type UserController struct {
	userService services.UserService

	logger slog.Logger
}

func CreateUserController(logger slog.Logger, userService services.UserService) (c *UserController) {
	c = &UserController{}
	c.logger = *logger.With()
	c.userService = userService
	return
}

func (controller *UserController) CreateService(c *gin.Context) {
	var serverModel models.ServiceModel
	err := c.ShouldBindJSON(&serverModel)
	if err != nil {
		controller.logger.Error("Failed to bind ServerModel", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Failed to bind ServerModel"})
		return
	}

	err = controller.userService.CreateService(&serverModel)
	if err != nil {
		controller.logger.Error("Can't create service", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Can't create service: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})

}

func (controller *UserController) CreateMethod(c *gin.Context) {
	var method models.Method
	err := c.ShouldBindJSON(&method)
	if err != nil {
		controller.logger.Error("Failed to bind Method", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Failed to bind Method"})
		return
	}

	err = controller.userService.CreateMethod(&method)
	if err != nil {
		controller.logger.Error("Can't create method", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Can't create method: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})

}

func (controller *UserController) CreateArgument(c *gin.Context) {
	var argument models.Argument
	err := c.ShouldBindJSON(&argument)
	if err != nil {
		controller.logger.Error("Failed to bind Argument", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Failed to bind Argument"})
		return
	}

	err = controller.userService.CreateArgument(&argument)
	if err != nil {
		controller.logger.Error("Can't create argument", slog.String("error", err.Error()))
		c.JSON(400, gin.H{"error": "Can't create argument: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})

}

func (controller *UserController) UpdateService(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid service ID"})
		return
	}

	var req models.UpdateServiceDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind updates: " + err.Error()})
		return
	}

	if err = controller.userService.UpdateService(uint(id), &req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update service: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *UserController) UpdateMethod(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid method ID"})
		return
	}

	var req models.UpdateMethodDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind updates: " + err.Error()})
		return
	}

	if err = controller.userService.UpdateMethod(uint(id), &req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update method: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *UserController) UpdateArgument(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid argument ID"})
		return
	}

	var req models.UpdateArgumentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "failed to bind updates: " + err.Error()})
		return
	}

	if err = controller.userService.UpdateArgument(uint(id), &req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update argument: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *UserController) DeleteService(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid service ID"})
		return
	}

	if err = controller.userService.DeleteService(uint(id)); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete service: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *UserController) DeleteMethod(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid method ID"})
		return
	}

	if err = controller.userService.DeleteMethod(uint(id)); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete method: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *UserController) DeleteArgument(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid argument ID"})
		return
	}

	if err = controller.userService.DeleteArgument(uint(id)); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete argument: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *UserController) GetAllServices(c *gin.Context) {
	allServ, err := controller.userService.GetAllServices()

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get services: " + err.Error()})
		return
	}
	c.JSON(200, allServ)
}

func (controller *UserController) GetServiceByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("ID"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid service ID"})
		return
	}

	serv, err := controller.userService.GetServiceByID(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get service: " + err.Error()})
		return
	}
	c.JSON(200, serv)
}
