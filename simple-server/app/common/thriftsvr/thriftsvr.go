package thriftsvr

import (
	"context"
	"time"

	"simple-server/idl/thrift/proto"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&ThriftServer{}).AsServer()
}

type ThriftServer struct {
	Addr string            `value:"${thrift.server.addr:=0.0.0.0:9292}"`
	Echo proto.EchoService `autowire:""`
	svr  *thrift.TSimpleServer
}

func (s *ThriftServer) ListenAndServe(sig gs.ReadySignal) error {
	transport, err := thrift.NewTServerSocket(s.Addr)
	if err != nil {
		return err
	}
	processor := proto.NewEchoServiceProcessor(s.Echo)
	s.svr = thrift.NewTSimpleServer2(processor, transport)
	<-sig.TriggerAndWait()
	return s.svr.Serve()
}

func (s *ThriftServer) Shutdown(ctx context.Context) error {
	thrift.ServerStopTimeout = time.Second
	return s.svr.Stop()
}
