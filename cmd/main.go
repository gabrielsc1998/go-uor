package main

import (
	"fmt"

	usecases "github.com/gabrielsc98/go-uow/internal/application/use-cases"
	"github.com/gabrielsc98/go-uow/internal/infra/controllers"
	"github.com/gabrielsc98/go-uow/internal/infra/db"
	"github.com/gabrielsc98/go-uow/internal/infra/repositories"
	"github.com/gabrielsc98/go-uow/internal/infra/server"
	"github.com/gabrielsc98/go-uow/internal/infra/uow"
)

func main() {
	database := db.NewDB()
	err := database.Connect(db.DBCredentials{
		Host:     "localhost",
		Port:     "3307",
		User:     "root",
		Password: "root",
		Database: "test",
	})
	defer database.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	uow := uow.NewUnitOfWork(database)
	uow.RegisterRepository(
		"users", repositories.NewUserRepository(database),
	)
	uow.RegisterRepository(
		"companies", repositories.NewCompanyRepository(database),
	)

	createCompanyUseCase := usecases.NewCreateCompany(uow)
	createCompanyController := controllers.NewCreateCompanyController(createCompanyUseCase)

	server := server.NewServer("8080")
	server.AddRoute("POST", "/companies", createCompanyController.Handle)
	server.Start()
}
