//+build wireinject

package main

import (
	"HelloWorld/internal/config"
	"HelloWorld/internal/dao"
	"HelloWorld/internal/repo"
	"HelloWorld/internal/server"
	"HelloWorld/internal/service"

	"github.com/google/wire"
)

func InitServer(configPath string) (*server.Server, func(), error) {
	// 每个领域对象都有一个repo，用set组装
	//repoSet := wire.NewSet(repo.NewUserRepo)
	wire.Build(server.NewSever, service.NewService, repo.NewUserRepo, dao.NewDao, config.NewConfig)
	return &server.Server{}, nil, nil
}
