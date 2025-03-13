package grpc_controller

import (
	"context"

	"simple-server/idl/grpc/proto"
)

func (c *Controller) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	return &proto.EchoResponse{Message: req.Message}, nil
}
