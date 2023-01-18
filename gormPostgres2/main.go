package main

import (
	"github.com/joho/godotenv"
	"postgres/gormPostgres2/connection"
	"postgres/gormPostgres2/handlers"
)

func main() {
	godotenv.Load("config.env")
	connection.DatabaseConnection()
	handlers.Handler()

}
