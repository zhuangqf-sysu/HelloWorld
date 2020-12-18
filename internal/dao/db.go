package dao

import (
	"context"

	"HelloWorld/internal/model"
)

// GetUserById
func (dao *Dao) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	// 从数据库中取数据
	return &model.User{Id: id}, nil
}
