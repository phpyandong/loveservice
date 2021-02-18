package data

import (
	"loveservice/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper("data/greeter", logger),
	}
}
//MySQL 等的具体实现。无需new redis 及mysql
func (r *greeterRepo) CreateGreeter(g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(g *biz.Greeter) error {
	return nil
}
