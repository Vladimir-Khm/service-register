package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"service-register/internal/config"
	"service-register/internal/models"
	"service-register/internal/repositories"
	"service-register/internal/repositories/postgres/user"
)

func CreateRepository(db *gorm.DB) *repositories.Repository {
	return &repositories.Repository{
		User: user.CreateUserRepository(db),
	}
}

func ConnectDB(config *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName,
		config.DbPort)

	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		),
	}
	db, err := gorm.Open(postgres.Open(connectionString), gormConfig)
	if err != nil {
		return nil, err
	}

	// --- Configure the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if config.DbMaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(config.DbMaxOpenConns)
	}

	if config.DbMaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(config.DbMaxIdleConns)
	}

	if config.DbConnMaxLifetimeSec != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(config.DbConnMaxLifetimeSec) * time.Second)
	}

	if config.DbConnMaxIdleTimeSec != 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(config.DbConnMaxIdleTimeSec) * time.Second)
	}

	err = migrate(db)
	if err != nil {
		fmt.Println("WRONG MIGRATION")
		return nil, err
	}
	fmt.Println("SUCCESS MIGRATION")

	return db, nil
}

func migrate(db *gorm.DB) error {

	if err := db.AutoMigrate(
		&models.ServiceModel{},
		&models.Method{},
		&models.Argument{},
	); err != nil {
		return fmt.Errorf("database migration error: %w", err)
	}

	return nil
}
