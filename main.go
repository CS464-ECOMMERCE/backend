package main

import (
	"backend/configs"
	"backend/grpc"
	"backend/routes"
	"backend/storage"
)

func main() {
	configs.InitEnv()            // init env
	storage.GetStorageInstance() // init db
	grpc.Init()
	routes.InitRoutes() // init controller routes
}
