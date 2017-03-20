package main

import (
	pb "Bookstore/pb"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type bookServer struct{}

func newServer() pb.BookServiceServer {
	return new(bookServer)
}

func (s *bookServer) ListBooks(ctx context.Context, e *empty.Empty) (*pb.BookResp, error) {
	b1 := new(pb.Book)
	b1.Id = 1
	b1.Title = "A"

	b2 := &pb.Book{
		Id:    2,
		Title: "B",
	}

	bs := &pb.BookResp{}
	bs.Books = append(bs.Books, b1)
	bs.Books = append(bs.Books, b2)

	return bs, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalln("listen failed: ", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterBookServiceServer(server, newServer())
	err = server.Serve(listen)
	if err != nil {
		log.Fatalln("server failed: ", err.Error())
	}
}
