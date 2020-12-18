package service

import (
	"HelloWorld/internal/model"
	"HelloWorld/internal/repo"
	"context"
)

type Service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) (*Service, error) {
	return &Service{userRepo: userRepo}, nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetUser(ctx, id)
}
