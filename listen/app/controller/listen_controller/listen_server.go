package listen_controller

import (
	"fmt"
	"golang.org/x/net/context"
	"listen/app/proto/listen"
)

type ListenController struct{}

func (l *ListenController) ListenData(ctx context.Context, in *listen.Request) (*listen.Response, error) {
	return &listen.Response{Message : fmt.Sprintf("[%s]", in.Name)}, nil
}
