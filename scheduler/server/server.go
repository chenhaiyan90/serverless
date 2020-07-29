package server

import (
	"com/aliyun/serverless/scheduler/core"
	pb "com/aliyun/serverless/scheduler/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Server struct {
}

func (s Server) AcquireContainer(ctx context.Context, req *pb.AcquireContainerRequest) (*pb.AcquireContainerReply, error) {
	startTime := time.Now().UnixNano()
	str, _ := json.Marshal(req)
	fmt.Println(startTime, string(str))
	if req.AccountId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "account ID cannot be empty")
	}

	if req.FunctionConfig == nil {
		return nil, status.Errorf(codes.InvalidArgument, "function config cannot be nil")
	}

	reply, err := core.AcquireContainer(req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s Server) ReturnContainer(ctx context.Context, req *pb.ReturnContainerRequest) (*pb.ReturnContainerReply, error) {
	startTime := time.Now().UnixNano()
	str, _ := json.Marshal(req)
	fmt.Println(startTime, string(str))
	reply, err := core.ReturnContainer(req)
	return reply, err
}
