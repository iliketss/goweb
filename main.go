package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"goweb/services"
	"log"
	"net"
)

func main() {
	creds, err2 := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
	if err2 != nil {
		log.Fatalln(err2)
	}

	g := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(g, &services.ProductService)

	list, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
	err = g.Serve(list)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("grpc启动成功")
}
