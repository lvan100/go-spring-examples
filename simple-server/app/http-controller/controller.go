package http_controller

import (
	"simple-server/idl/http/proto"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&Controller{}).Export(gs.As[proto.Controller]())
}

var _ proto.Controller = (*Controller)(nil)

type Controller struct {
	BookController
}
