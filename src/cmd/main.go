package main

import (
	"fmt"
	"net/http"
	"src/internal/configs"
	"src/internal/persistence/database"
	"src/internal/router"
	"src/internal/setup"
)

func main() {
    config, configErr := configs.LoadConfig()

	if configErr != nil {
		fmt.Print("Erro ao carregar as configurações", configErr)
	}

	_, connectErr := database.Connect()
     
	if connectErr != nil {
		fmt.Print("Erro ao conectar no banco de dados", connectErr)
	}

	database.RunMigration()
	database.Seed()

	repositorySetup := setup.NewRepositorySetup()
    useCaseSetup := setup.NewUseCaseSetup(repositorySetup)
	handlerSetup := setup.NewHandlerSetup(useCaseSetup)

	route := router.NewRoute(handlerSetup)

	server := &http.Server{
		Addr: ":"+ config.ServerPort,
		Handler: route,
	}

	serverErr  := server.ListenAndServe()

	if serverErr != nil {
		fmt.Print(serverErr)
	}
}