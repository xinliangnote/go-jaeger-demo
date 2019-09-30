package read_controller

import (
	"fmt"
	"golang.org/x/net/context"
	"read/app/proto/read"
	"read/app/util"
	"read/app/proto/listen"
	"read/app/util/grpc_client"
)

type ReadController struct{}

func (s *ReadController) ReadData(ctx context.Context, in *read.Request) (*read.Response, error) {

	// 调用 gRPC 服务
	grpcListenClient := listen.NewListenClient(grpc_client.CreateServiceListenConn(ctx))
	resListen, _ := grpcListenClient.ListenData(context.Background(), &listen.Request{Name: "listen"})

	// 调用 HTTP 服务
	resHttpGet := ""
	_, err := util.HttpGet("http://localhost:9905/sing", ctx)
	if err == nil {
		resHttpGet = "[HttpGetOk]"
	}

	msg := "[" + fmt.Sprintf("%s", in.Name) + "-" +
		   resListen.Message + "-" +
		   resHttpGet +
		   "]"
	return &read.Response{Message : msg}, nil
}
