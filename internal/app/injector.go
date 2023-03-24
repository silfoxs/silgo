package app

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Logger *logger.Logger
	Db     *gorm.DB
}
