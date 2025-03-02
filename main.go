package main

import (
	"backend/configs"
	_ "backend/docs"
	"backend/storage"

	// "backend/mq"
	"backend/routes"
)

func main() {
	configs.InitEnv()   // init env
	// mq.Init()           // init rabbitmq connection
	storage.GetStorageInstance()	// init db
	routes.InitRoutes() // init controller routes
}
