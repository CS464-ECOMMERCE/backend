package main

import (
	"backend/configs"
	_ "backend/docs"
	"backend/storage"

	// "backend/mq"
	"backend/routes"
	// "backend/services"
)

func main() {
	configs.InitEnv()   // init env
	// services.Init()     // init s3
	// mq.Init()           // init rabbitmq connection
	storage.GetStorageInstance()	// init db
	routes.InitRoutes() // init controller routes
}
