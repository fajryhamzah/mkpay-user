package main

import (
	"github.com/fajryhamzah/mkpay-user/controllers"
	"github.com/fajryhamzah/mkpay-user/controllers/middlewares"
	db "github.com/fajryhamzah/mkpay-user/db"
	authHandlers "github.com/fajryhamzah/mkpay-user/handlers/auth"
	userHandlers "github.com/fajryhamzah/mkpay-user/handlers/user"
	server "github.com/fajryhamzah/mkpay-user/server"
	"github.com/fajryhamzah/mkpay-user/src/user"
	"github.com/fajryhamzah/mkpay-user/utils/cache"
	"github.com/joho/godotenv"
)

func main() {
	cache.Get()

	var userRepo user.RepoInterface
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env var")
	}

	dbConnection := db.GetInstance()
	conn := dbConnection.GetConnection()

	//define repository
	userRepo = user.New(conn)

	//define handlers
	authHandler := authHandlers.NewAuthHandler(userRepo)
	userHandler := userHandlers.NewUserHandler(userRepo)

	//define controller
	controllers := controllers.New(authHandler, userHandler)

	//define middleware
	middlewares := middlewares.NewAuthMiddleware(userRepo)

	//define web server
	webserver := server.New(
		controllers,
		middlewares,
	)

	webserver.InitServer()

	defer dbConnection.Close()
}
