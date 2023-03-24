package app

import (
	"context"
	"os"

	"time"

	"fmt"

	"github.com/silfoxs/silgo/pkg/shutdown"
)

func Run() {
	pid := os.Getpid()
	injector, _, err := BuildInjector()
	if err != nil {
		fmt.Printf("pid:%d build injector error: %s", pid, err.Error())
	}
	go func() {
		if err := injector.Server.ListenAndServe(); err != nil {
			injector.Logger.Errorf("pid:%d %s", pid, err.Error())
		}
	}()
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := injector.Server.Shutdown(ctx); err != nil {
				injector.Logger.Errorf("pid:%d server close err %s", pid, err.Error())
			} else {
				injector.Logger.Infof("pid:%d server close success", pid)
			}
		},
	)
}
