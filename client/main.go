/*
@Time : 2022/10/27 13:54
@Author : lianyz
@Description :
*/

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/lianyz/workflow/pei"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	engine := pb.NewProcessEngineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := engine.GetProcessDefine(ctx,
		&pb.GetProcessDefineRequest{
			ProcessDefineID: "1",
		})
	if err != nil {
		log.Fatalf("get process define failed: %v", err)
	}

	log.Printf("get process define : %v successfuly", res)

	instance := pb.ProcessInstance{}
	reply, err := engine.CreateProcessInstance(ctx, &instance)
	if err != nil {
		log.Fatalf("create process instance failed.: %v", err)
	}
	log.Printf("Create process instance reply: %v", reply)
}
