package main

import (
	"context"
	pb "github.com/loxt/go-grpc-server-streaming/pb/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
	"time"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}

func (s *Server) BeatsPerMinute(in *pb.BeatsPerMinuteRequest, stream pb.HelloService_BeatsPerMinuteServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return status.Errorf(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)

			value := 30 + rand.Int31n(80)

			err := stream.SendMsg(&pb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})

			if err != nil {
				return status.Errorf(codes.Canceled, "Stream has ended")
			}
		}
	}

	return nil
}

func main() {
	println("Running gRPC server...")

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
