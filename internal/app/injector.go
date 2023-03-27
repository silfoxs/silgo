package app

import (
	"net/http"

	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/pkg/logger"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Logger *logger.Logger
	Server *http.Server
}
