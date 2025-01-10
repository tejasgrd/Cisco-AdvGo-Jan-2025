package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)

	ctx := context.Background()
	addReq := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addRes, err := client.Add(ctx, addReq)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Add Result :", addRes.GetResult())
}
