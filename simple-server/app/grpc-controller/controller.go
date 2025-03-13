package grpc_controller

import (
	"simple-server/idl/grpc/proto"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&Controller{}).Export(gs.As[proto.EchoServiceServer]())
}

var _ proto.EchoServiceServer = (*Controller)(nil)

type Controller struct {
	proto.UnimplementedEchoServiceServer
}
