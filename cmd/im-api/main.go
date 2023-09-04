//go:generate go run main.go

//go:network
package main

import (
	"context"
	"github.com/alibaba/ioc-golang"
	"reason-im/internal/api"
	"reason-im/internal/utils/logger"
)

func main() {
	err := start()
	if err != nil {
		panic(err)
	}
}

func start() error {
	if err := ioc.Load(); err != nil {
		logger.Warn(context.Background(), "ioc load failed", err)
		return err
	}
	router := api.NewGinRouter()
	err := router.Run("0.0.0.0:1080")
	if err != nil {
		logger.Warn(context.Background(), "api run failed", err)
		return err
	}
	return nil
}
