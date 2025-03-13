/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc_controller

import (
	"simple-server/idl/grpc/proto"
	"simple-server/util"

	"github.com/go-spring/spring-core/gs"
	"google.golang.org/grpc"
)

func init() {
	gs.Object(&Controller{})
	gs.Provide(func(c *Controller) util.GrpcServerConfiger {
		return func(svr *grpc.Server) {
			proto.RegisterEchoServiceServer(svr, c)
		}
	})
}

var _ proto.EchoServiceServer = (*Controller)(nil)

type Controller struct {
	proto.UnimplementedEchoServiceServer
}
