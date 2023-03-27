package app

import (
	"net/http"

	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/router"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Logger *logger.Logger
	Db     *gorm.DB
	Server *http.Server
	router *router.Router
}
