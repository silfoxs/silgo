//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/action"
	"github.com/silfoxs/silgo/internal/app/repository"
	"github.com/silfoxs/silgo/internal/app/router"
	"github.com/silfoxs/silgo/internal/app/service"
	"github.com/silfoxs/silgo/internal/pkg/database"
	"github.com/silfoxs/silgo/internal/pkg/http"
	"github.com/silfoxs/silgo/internal/pkg/logger"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		logger.New,
		database.New,
		http.NewHttp,
		router.RouterSet,
		action.ActionSet,
		service.ServiceSet,
		repository.RepositorySet,
		InjectorSet,
	)

	return new(Injector), nil, nil
}
