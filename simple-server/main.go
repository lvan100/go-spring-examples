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

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	grpcclt "simple-server/idl/grpc/proto"
	thriftclt "simple-server/idl/thrift/proto"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/sysconf"
	"github.com/go-spring/spring-core/util/syslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "simple-server/app"
	_ "simple-server/biz"

	_ "github.com/go-spring/spring-core/gs/monitor"
)

func init() {
	gs.SetActiveProfiles("online")
	gs.EnableSimpleHttpServer(true)
	_ = sysconf.Set("spring.monitor.enable", "true")
}

func main() {
	_ = os.Unsetenv("_")
	_ = os.Unsetenv("TERM")
	_ = os.Unsetenv("TERM_SESSION_ID")
	go func() {
		time.Sleep(time.Millisecond * 500)
		runTest()
	}()
	if err := gs.Run(); err != nil {
		syslog.Errorf("app run failed: %s", err.Error())
	}
}

func runTest() {

	// http
	{
		resp, err := http.Get("http://127.0.0.1:9090/books")
		if err != nil {
			log.Fatalf("Failed to get: %v", err)
		}
		fmt.Println("Response from server:", resp.Status)
	}

	// grpc
	{
		conn, err := grpc.NewClient(":9494", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		defer conn.Close()

		client := grpcclt.NewEchoServiceClient(conn)

		response, err := client.Echo(context.Background(), &grpcclt.EchoRequest{Message: "Hello, gRPC!"})
		if err != nil {
			log.Fatalf("Error calling Echo: %v", err)
		}

		fmt.Println("Response from server:", response.Message)
	}

	// thrift
	{
		transport := thrift.NewTSocketConf(":9292", nil)
		defer transport.Close()

		protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)
		client := thriftclt.NewEchoServiceClientFactory(transport, protocolFactory)

		if err := transport.Open(); err != nil {
			log.Fatalf("Error opening transport: %v", err)
		}

		response, err := client.Echo(context.Background(), &thriftclt.EchoRequest{Message: "Hello, Thrift!"})
		if err != nil {
			log.Fatalf("Error calling Echo: %v", err)
		}

		fmt.Println("Response from server:", response.Message)
	}

	gs.ShutDown()
}
