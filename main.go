package main

import (
	"github.com/fajryhamzah/mkpay-user/controllers"
	db "github.com/fajryhamzah/mkpay-user/db"
	server "github.com/fajryhamzah/mkpay-user/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env var")
	}

	dbConnection := db.GetInstance()
	dbConnection.GetConnection()

	//define controller
	controllers := controllers.New()

	//define web server
	webserver := server.New(
		controllers,
	)

	webserver.InitServer()

	defer dbConnection.Close()
}
