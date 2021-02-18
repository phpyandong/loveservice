package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}
//依赖的是接口，而不是具体实现
type GreeterRepo interface {
	CreateGreeter(*Greeter) error
	UpdateGreeter(*Greeter) error
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}
//控制翻转的思想
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper("usecase/greeter", logger)}
}

func (uc *GreeterUsecase) Create(g *Greeter) error {
	return uc.repo.CreateGreeter(g)
}

func (uc *GreeterUsecase) Update(g *Greeter) error {
	return uc.repo.UpdateGreeter(g)
}
