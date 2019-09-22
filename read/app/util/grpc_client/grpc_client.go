package grpc_client

import (
	"fmt"
	"read/app/util/jaeger_service"
	"google.golang.org/grpc"
)

func CreateServiceListenConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9901")
}

func CreateServiceSpeakConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9902")
}

func CreateServiceReadConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9903")
}

func CreateServiceWriteConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9904")
}

func createGrpcClient(serviceAddress string ) *grpc.ClientConn {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithUnaryInterceptor(jaeger_service.ClientInterceptor(jaeger_service.Tracer, jaeger_service.ParentContext)))
	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
