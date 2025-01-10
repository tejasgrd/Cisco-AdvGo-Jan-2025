package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

// proto.AppServiceServer interface implementation
func (appService *AppServiceImpl) Add(ctx context.Context, r *proto.AddRequest) (*proto.AddResponse, error) {
	x := r.GetX()
	y := r.GetY()

	log.Printf("Processing [Add], x = %d and y = %d\n", x, y)
	result := x + y

	response := &proto.AddResponse{
		Result: result,
	}

	return response, nil
}

func (appService *AppServiceImpl) GeneratePrimes(r *proto.PrimesRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := r.GetStart()
	end := r.GetEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Println("[GeneratePrimes] sending prime no :", no)
			res := &proto.PrimesResponse{
				PrimeNo: no,
			}
			if err := serverStream.Send(res); err != nil {
				log.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	asi := &AppServiceImpl{}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAppServiceServer(grpcServer, asi)

	grpcServer.Serve(listener)
}
