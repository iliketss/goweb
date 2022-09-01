package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"goweb/services"
	"log"
)

func main() {

	creds, err2 := credentials.NewClientTLSFromFile("cert/server.pem", "*mszlu.com")
	if err2 != nil {
		log.Fatalln(err2)
	}

	coon, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer coon.Close()
	client := services.NewProdServiceClient(coon)
	request := &services.ProductRequest{ProdId: 123}
	stock, err := client.GetProductStock(context.Background(), request)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Check success", stock.ProdStock)
}
