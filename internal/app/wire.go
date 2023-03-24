//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/pkg/database"
	"github.com/silfoxs/silgo/internal/pkg/logger"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		logger.New,
		database.New,
		InjectorSet,
	)

	return new(Injector), nil, nil
}
