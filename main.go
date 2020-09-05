package main

import (
	db "github.com/fajryhamzah/mkpay-user/db"
	server "github.com/fajryhamzah/mkpay-user/server"
	"github.com/joho/godotenv"
)

func main() {
	var dbConnection db.DbInterface

	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env var")
	}

	dbConnection = db.GetInstance()
	dbConnection.GetConnection()

	server.InitServer()
}
