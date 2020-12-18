package server

import (
	"context"
	"log"

	"HelloWorld/api"
	"HelloWorld/internal/service"
)

type GrpcServerImpl struct {
	service *service.Service
}

func (g GrpcServerImpl) GetUser(ctx context.Context, req *api.IDReq) (*api.UserReply, error) {
	// 参数校验

	// 依赖service层
	user, err := g.service.GetUser(ctx, req.Id)
	if err != nil {
		// log
		log.Printf("GetUser(%v) err(%v)\n", req.Id, err)
		return nil, err
	}
	return user.ToReply(), nil
}

func (g GrpcServerImpl) ListUser(ctx context.Context, query *api.UserQuery) (*api.UsersReply, error) {
	panic("implement me")
}
