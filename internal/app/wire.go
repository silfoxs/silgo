//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/pkg/database"
	"github.com/silfoxs/silgo/internal/pkg/http"
	"github.com/silfoxs/silgo/internal/pkg/logger"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		logger.New,
		database.New,
		http.NewHttp,
		InjectorSet,
	)

	return new(Injector), nil, nil
}
