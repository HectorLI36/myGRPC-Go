package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "myGRPC/pb"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial error: %v", err)
	}
	defer conn.Close()
	client := pb.NewSliceServerClient(conn)
	reply, err := client.GetSlice(context.Background(), &pb.SliceRequest{
		BookId:       1,
		Offset:       2,
		SlicePerPage: 3,
	})
	if err != nil {
		fmt.Printf("GetSlice RPC error: %v\n", err)
	}
	fmt.Printf("GRPC Server Response is \n")
	fmt.Print(reply)
}
