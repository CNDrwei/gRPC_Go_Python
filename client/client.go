package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "test/protobuf"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "wei"})
	if err != nil {
		fmt.Printf("could not greet %v\n", err)
	}
	fmt.Printf("greeting : %s", r.Message)

}
