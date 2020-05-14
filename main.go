package main

import (
	echo2 "github.com/labstack/echo/v4"
	"golang_api/databases"
	"golang_api/handlers"
	"golang_api/helpers"
	"golang_api/repositories/repo_impl"
	"golang_api/routes"
)

func main() {
	sql := &databases.PostgresDB{
		Host:     "localhost",
		Port:     "5432",
		User:     "ntc",
		Password: "1",
		DBName:   "go_db",
	}

	sql.Connect()
	defer sql.Close()

	app := echo2.New()
	structValidate := helpers.NewStructValidator()
	structValidate.RegisterValidate()
	app.Validator = structValidate

	userHandler := handlers.HandlerUser{UserRepo:repo_impl.NewUserRepo(sql)}

	api := routes.API{
		Echo: app,
		UserHandler:userHandler,
	}

	api.SetupRouter()

	app.Logger.Fatal(app.Start(":3000"))
}
