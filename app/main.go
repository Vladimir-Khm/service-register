package main

import (
	"log"

	"service-register/internal/config"
	"service-register/internal/repositories/postgres"
	"service-register/internal/server/rest"
)

// @title Task and Contest Management API
// @version 1.0.0
// @description API for managing user profiles, tasks, and contests with TON Proof authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host ton-reputation-backend-development-cd-808855530615.europe-central2.run.app
// @BasePath /

// @securityDefinitions.bearerAuth Bearer
// @in header
// @name Authorization

// @tag.name Authorization
// @tag.description Operations about authorization

// @securityDefinitions.bearerAuth Bearer
// @in header
// @name Authorization
// @type apiKey
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Config loading error: " + err.Error())
	}

	dbContext, err := postgres.ConnectDB(config)
	if err != nil {
		log.Fatalln("DbContext creating error: " + err.Error())
	}

	server := rest.CreateServer(config, dbContext)

	log.Fatalln(server.Run(":8080"))
}
