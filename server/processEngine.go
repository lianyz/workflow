/*
@Time : 2022/10/27 14:11
@Author : lianyz
@Description :
*/

package main

import (
	"context"
	pb "github.com/lianyz/workflow/pei"
)

type processEngine struct {
}

func (p processEngine) GetProcessDefine(ctx context.Context, request *pb.GetProcessDefineRequest) (*pb.ProcessDefine, error) {
	//TODO implement me
	processDefine := pb.ProcessDefine{
		Id:   "1",
		Name: "测试流程",
	}
	return &processDefine, nil
}

func (p processEngine) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	//TODO implement me
	return nil, nil
}

func (p processEngine) CreateProcessInstance(ctx context.Context, instance *pb.ProcessInstance) (*pb.CreateProcessInstanceReply, error) {
	//TODO implement me
	return nil, nil
}

func (p processEngine) GetInstance(ctx context.Context, request *pb.GetInstanceRequest) (*pb.ProcessInstance, error) {
	//TODO implement me
	return nil, nil
}

func (p processEngine) GetTasks(ctx context.Context, request *pb.GetTasksRequest) (*pb.GetTasksReply, error) {
	//TODO implement me
	return nil, nil
}

func (p processEngine) CompleteTask(ctx context.Context, request *pb.CompleteTaskRequest) (*pb.CompleteTaskReply, error) {
	//TODO implement me
	return nil, nil
}
