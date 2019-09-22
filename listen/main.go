package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"listen/app/controller/listen_controller"
	"listen/app/proto/listen"
	"listen/app/util/jaeger_service"
	"log"
	"net"
	"os"
)

const (
	ServiceName     = "gRPC-Service-Listen"
	ServiceHostPort = "0.0.0.0:9901"

	JaegerHostPort  = "127.0.0.1:6831"
)

func main() {

	var serviceOpts []grpc.ServerOption

	tracer, _, err := jaeger_service.NewJaegerTracer(ServiceName, JaegerHostPort)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		serviceOpts = append(serviceOpts, jaeger_service.ServerOption(tracer))
	}

	l, err := net.Listen("tcp", ServiceHostPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(serviceOpts...)

	// 服务注册
	listen.RegisterListenServer(s, &listen_controller.ListenController{})

	log.Println("Listen on " + ServiceHostPort)
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
