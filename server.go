package main

import (
	"context"
	pb "myGRPC/pb"
	"strconv"
)

type server struct {
}

func (s *server) GetSlice(ctx context.Context, in *pb.SliceRequest) (*pb.SliceResponse, error) {
	return &pb.SliceResponse{
		Content:      "hi",
		Title:        "hi" + strconv.Itoa(int(in.BookId)),
		SliceNumber:  0,
		SlicePerPage: 0,
		Offset:       0,
		BookId:       0,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterSliceServerServer(s, &server{})
	reflection.Register(s)
	//在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
	}

	// --------------------------
	//lis, err := net.Listen("tcp", ":8972")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//s := grpc.NewServer()
	//pb.RegisterSliceServerServer(s, &server{})
	//err = s.Serve(lis)
	//if err != nil {
	//	log.Fatal("failed to serve: %v", err)
	//}

}
