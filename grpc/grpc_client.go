package grpc

import (
	"backend/configs"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiServer struct {
	ProductServiceConn *grpc.ClientConn
}

func mustConnGRPC(conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	log.Printf("grpc: connecting to %s", addr)
	
	if err != nil {
		log.Printf("grpc: failed to connect %s", addr)
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}

var ApiServerInstance *ApiServer

func Init() {
	ApiServerInstance = &ApiServer{}
	mustConnGRPC(&ApiServerInstance.ProductServiceConn, configs.PRODUCT_SERVICE_ADDR)
}
