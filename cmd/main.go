package main

import (
	"fmt"
	"net/http"
	"src/internal/configs"
	"src/internal/persistence/database"
	"src/internal/router"
	"src/internal/setup"
	"src/internal/usecase"
)

func main() {

	config, configErr := configs.LoadConfig()
	if configErr != nil {
		fmt.Println("Erro ao carregar as configurações:", configErr)
		return
	}

	db, connectErr := database.Connect()
	if connectErr != nil {
		fmt.Println("Erro ao conectar no banco de dados:", connectErr)
		return
	}

	database.RunMigration(db)
	database.Seed(db)
	otherSetup := setup.NewOtherSetup()
	repositorySetup := setup.NewRepositorySetup(db)
	useCaseSetup := setup.NewUseCaseSetup(repositorySetup)
	handlerSetup := setup.NewHandlerSetup(useCaseSetup, otherSetup)

	pool := usecase.NewPool()
	go pool.Start()

	route := router.NewRoute(handlerSetup, pool)

	server := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: route,
	}

	fmt.Println("Servidor iniciado na porta", config.ServerPort)
	serverErr := server.ListenAndServe()

	if serverErr != nil {
		fmt.Println("Erro no servidor:", serverErr)
	}
}
