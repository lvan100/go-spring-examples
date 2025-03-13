package grpcsvr

import (
	"context"
	"log"
	"net"

	"simple-server/idl/grpc/proto"

	"github.com/go-spring/spring-core/gs"
	"google.golang.org/grpc"
)

func init() {
	gs.Object(NewGrpcServer()).AsServer()
}

type GrpcServer struct {
	Addr string                  `value:"${grpc.server.addr:=0.0.0.0:9494}"`
	Echo proto.EchoServiceServer `autowire:""`
	svr  *grpc.Server
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		svr: grpc.NewServer(),
	}
}

func (s *GrpcServer) ListenAndServe(sig gs.ReadySignal) error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	proto.RegisterEchoServiceServer(s.svr, s.Echo)
	<-sig.TriggerAndWait()
	return s.svr.Serve(listener)
}

func (s *GrpcServer) Shutdown(ctx context.Context) error {
	s.svr.GracefulStop()
	return nil
}
