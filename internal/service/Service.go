package service

import (
	"HelloWorld/internal/model"
	"HelloWorld/internal/repo"
	"context"
)

type Service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) (*Service, func(), error) {
	service := &Service{userRepo: userRepo}
	return service, service.Close, nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *Service) Close() {

}
