package main

import (
	gw "Bookstore/pb"
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var addr = "localhost:50050"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dialOps := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()
	err := gw.RegisterBookServiceHandlerFromEndpoint(ctx, mux, addr, dialOps)
	if err != nil {
		log.Fatalln("register handler error:", err.Error())
	}

	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalln("http err:", err.Error())
	}
}
