package rest

import (
	"log/slog"
	"os"
	"time"

	"service-register/internal/middlewares"
	"service-register/internal/repositories/postgres"
	"service-register/internal/server/rest/handlers/user"
	"service-register/internal/services"

	"gorm.io/gorm"
	"service-register/internal/services/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"service-register/internal/config"
)

type Server struct {
	router *gin.Engine
}

type controller struct {
	user *user.UserController
}

func CreateServer(config *config.Config, db *gorm.DB) *Server {
	s := &Server{}

	logger := createLogger(config)
	repository := postgres.CreateRepository(db)
	service := domain.CreateService(config, *logger, repository)
	controller := createController(config, logger, service)

	authMiddleware := middlewares.CreateAuthMiddleware(config)

	s.initRoutes(config, authMiddleware, controller)

	return s
}

func createController(config *config.Config, logger *slog.Logger, domainServices *services.Service) *controller {
	return &controller{
		user: user.CreateUserController(*logger, domainServices.User),
	}
}

func (s *Server) initRoutes(config *config.Config, authMiddleware *middlewares.AuthMiddleware, controller *controller) {
	s.router = gin.Default()

	allowedOrigins := []string{"*"}
	allowedHeaders := []string{"*"}
	maxAge := 30 * time.Minute

	corsConfig := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     allowedHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	}

	s.router.Use(cors.New(corsConfig))
	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "Oh yeah")
	})

	s.router.POST("/service", controller.user.CreateService)
	s.router.POST("/method", controller.user.CreateMethod)
	s.router.POST("/argument", controller.user.CreateArgument)

	s.router.PATCH("/service", controller.user.UpdateService)
	s.router.PATCH("/method", controller.user.UpdateMethod)
	s.router.PATCH("/argument", controller.user.UpdateArgument)

	s.router.DELETE("/service", controller.user.DeleteService)
	s.router.DELETE("/method", controller.user.DeleteMethod)
	s.router.DELETE("/argument", controller.user.DeleteArgument)

	s.router.GET("/allServices", controller.user.GetAllServices)
	s.router.GET("/serviceByID", controller.user.GetServiceByID)

}

func createLogger(config *config.Config) *slog.Logger {
	var logLevel slog.Leveler

	if config.LogLevel == "debug" {
		logLevel = slog.LevelDebug
	} else if config.LogLevel == "info" {
		logLevel = slog.LevelInfo
	} else if config.LogLevel == "warn" {
		logLevel = slog.LevelWarn
	} else if config.LogLevel == "error" {
		logLevel = slog.LevelError
	} else {
		logLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	return logger
}

func (s *Server) Run(addr ...string) error {
	return s.router.Run(addr...)
}
