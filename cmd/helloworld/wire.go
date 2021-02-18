// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"loveservice/internal/biz"
	"loveservice/internal/conf"
	"loveservice/internal/data"
	"loveservice/internal/server"
	"loveservice/internal/service"
	"github.com/go-kratos/kratos"
	"github.com/go-kratos/kratos/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
