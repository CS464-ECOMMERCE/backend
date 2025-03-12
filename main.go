package main

import (
	"backend/configs"
	_ "backend/docs"
	"backend/grpc"
	"backend/routes"
	// "backend/mq"
)

func main() {
	configs.InitEnv() // init env
	// mq.Init()           // init rabbitmq connection
	// storage.GetStorageInstance() // init db
	grpc.Init()
	routes.InitRoutes() // init controller routes
}
