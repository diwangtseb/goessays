package main

import (
	"context"
	"fmt"
	"sgrpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	myString := make([]byte, 0, 4*1024*1024*2)
	appendedString := append(myString, []byte("H")...)
	client := pb.NewGreeterClient(conn)
	r, err := client.SayHello(context.TODO(), &pb.HelloRequest{
		Name: string(appendedString),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", r)
}
