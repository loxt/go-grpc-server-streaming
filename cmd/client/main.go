package main

import (
	"context"
	pb "github.com/loxt/go-grpc-server-streaming/pb/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	req := &pb.SayHelloRequest{Name: "Loxt"}

	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Message)
}
