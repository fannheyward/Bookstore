package main

import (
	pb "Bookstore/pb"
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

var (
	addr = "localhost:50050"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("dial failed:", err.Error())
	}
	defer conn.Close()

	cli := pb.NewBookServiceClient(conn)
	e := new(empty.Empty)
	bs, err := cli.ListBooks(context.Background(), e)
	if err != nil {
		log.Fatalln("list error:", err.Error())
	}

	log.Println("bs", bs)
}
