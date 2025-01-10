package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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

	// doRequestResponse(ctx, client)
	doServerStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
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

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	primesReq := &proto.PrimesRequest{
		Start: 2,
		End:   100,
	}
	clientStream, err := client.GeneratePrimes(ctx, primesReq)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All generated prime numbers have been received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No : %d\n", res.GetPrimeNo())
	}
}
