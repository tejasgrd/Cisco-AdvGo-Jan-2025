package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"

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
