package thrift_controller

import (
	"simple-server/idl/thrift/proto"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&Controller{}).Export(gs.As[proto.EchoService]())
}

var _ proto.EchoService = (*Controller)(nil)

type Controller struct {
	EchoController
}
