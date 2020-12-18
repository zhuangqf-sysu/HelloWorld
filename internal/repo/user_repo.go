package repo

import (
	"HelloWorld/internal/dao"
	"HelloWorld/internal/model"
	"context"
)

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	AddUser(ctx context.Context, user *model.User) error
	DelUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, user *model.User) error
	BatchGetUser(ctx context.Context, ids []int64) ([]*model.User, error)

	Close()
}

func NewUserRepo(dao *dao.Dao) (UserRepo, error) {
	return &UserRepoImpl{dao: dao}, nil
}

type UserRepoImpl struct {
	dao *dao.Dao
}

func (u UserRepoImpl) Close() {
}

func (u UserRepoImpl) GetUser(ctx context.Context, id int64) (*model.User, error) {
	//屏蔽数据存储细节（读取：cache -> db， 回写：db -> cache)
	return u.dao.GetUserById(ctx, id)
}

func (u UserRepoImpl) AddUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (u UserRepoImpl) DelUser(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u UserRepoImpl) UpdateUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (u UserRepoImpl) BatchGetUser(ctx context.Context, ids []int64) ([]*model.User, error) {
	panic("implement me")
}
