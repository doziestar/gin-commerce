package main

import (
	"gin-commerce/pkg/config"
	"gin-commerce/pkg/database"
	"gin-commerce/pkg/logger"
	"gin-commerce/routers"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.Connection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	db := database.GetDB()
	router := routers.Routes(db)

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
