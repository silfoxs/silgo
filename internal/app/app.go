package app

import (
	"os"

	"fmt"

	"github.com/silfoxs/silgo/pkg/shutdown"
)

func Run() {
	pid := os.Getpid()
	injector, clearup, err := BuildInjector()
	if err != nil {
		panic(
			fmt.Sprintf("pid:%d build injector error: %s", pid, err.Error()),
		)
	}
	go func() {
		if err := injector.Server.ListenAndServe(); err != nil {
			injector.Logger.Errorf("pid:%d %s", pid, err.Error())
		}
	}()
	shutdown.NewHook().Close(
		clearup,
	)
}
