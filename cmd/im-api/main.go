package main

import (
	"context"
	"reason-im/internal/api"
	"reason-im/internal/utils"
)

func main() {
	err := start()
	if err != nil {
		panic(err)
	}
}

func start() error {
	router := api.NewGinRouter()

	err := router.Run("0.0.0.0:1080")
	if err != nil {
		utils.Warn(context.Background(), "api run failed", err)
		return err
	}
	return nil
}