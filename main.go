package main

import (
	"log"

	"github.com/oliveira-a-rafael/mycareer-api/app"
	"github.com/oliveira-a-rafael/mycareer-api/config"
	"github.com/oliveira-a-rafael/mycareer-api/controllers"
	"github.com/oliveira-a-rafael/mycareer-api/infra/databasex"
	"github.com/oliveira-a-rafael/mycareer-api/infra/repositories"
	"github.com/oliveira-a-rafael/mycareer-api/services"
)

func main() {
	log.Println("Starting services..")
	buildApp()
}

func buildApp() {

	// @TODO renomear para database e excluir o antigo
	db := databasex.GetInstancePostgres()

	accountRepository := &repositories.AccountRepository{DB: db}

	// services
	accountService := &services.AccountService{AccountRepository: *accountRepository}

	// Inject controllers
	accountController := &controllers.AccountController{AccountService: *accountService}

	// Start server
	server := &app.Server{
		Config: app.ServerConfig{
			Port:         config.APP_PORT,
			AllowedHosts: config.ALLOWED_HOSTS,
		},
		AccountController: accountController,
	}

	server.Run()
}
