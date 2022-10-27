/*
@Time : 2022/10/27 14:08
@Author : lianyz
@Description :
*/

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/lianyz/workflow/pei"
)

const (
	port = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed listen: %v", err)
		return
	}

	rpcServer := grpc.NewServer()
	pb.RegisterProcessEngineServer(rpcServer, &processEngine{})

	log.Printf("Sttarting gRPC listener on port " + port)
	if err := rpcServer.Serve(listener); err != nil {
		log.Fatalf("failed tot serve: %v", err)
	}
}
