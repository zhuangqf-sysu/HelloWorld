package dao

import (
	"HelloWorld/internal/config"
)

type Dao struct {
	// 依赖组件
	// db
	// cache
	conf *config.Config
}

// 根据conf初始化依赖组件
func NewDao(conf *config.Config) (*Dao, func(), error) {
	dao := &Dao{conf: conf}
	return dao, dao.Close, nil
}

// 关闭
func (dao *Dao) Close() {
	//if dao.db != nil {
	//	dao.db.Close
	//}
}
