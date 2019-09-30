package write_controller

import (
	"fmt"
	"golang.org/x/net/context"
	"write/app/proto/write"
	"write/app/util"
	"write/app/util/grpc_client"
	"write/app/proto/listen"
)

type WriteController struct{}

func (s *WriteController) WriteData(ctx context.Context, in *write.Request) (*write.Response, error) {

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
	return &write.Response{Message : msg}, nil
}
