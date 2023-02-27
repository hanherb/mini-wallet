package main

import (
	"github.com/hanherb/mini-wallet/src/config"
	"github.com/hanherb/mini-wallet/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Initialization Env
	godotenv.Load(".env")
	config.VariablesInitialization()

	// Initialization DB
	config.MysqlInitialization()

	// Start REST Server on main thread
	router := routes.StartRoute()
	router.Run(":3000")
}
