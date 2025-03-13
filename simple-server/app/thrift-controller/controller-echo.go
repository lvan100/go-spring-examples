package thrift_controller

import (
	"context"

	"simple-server/idl/thrift/proto"
)

type EchoController struct {
}

func (c *EchoController) Echo(ctx context.Context, req *proto.EchoRequest) (r *proto.EchoResponse, err error) {
	return &proto.EchoResponse{Message: req.Message}, nil
}
