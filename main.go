package main

import (
	"flag"
	"log"
	"net"

	api "bitbucket.org/agrostar/grpc_test/proto"
	"bitbucket.org/agrostar/grpc_test/server"
	"google.golang.org/grpc"
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ss := server.SomeStruct{}
	var s api.SomeStructServer
	s = &ss
	grpcServer := grpc.NewServer()
	api.RegisterSomeStructServer(grpcServer, s)
	grpcServer.Serve(lis)
}
