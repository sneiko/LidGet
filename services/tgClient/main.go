package main

import (
	"context"
	"fmt"
	"lidget/domain/config"
	db2 "lidget/domain/pkg/db"
	"lidget/domain/pkg/logger"
	"lidget/tgClient/internal/tgClient"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := db2.Connect(cfg)
	if err != nil {
		panic(err)
	}

	client := tgClient.NewTgClient(&cfg.Tg, db)

	err = client.Run(ctx)
	if err != nil {
		return
	}

	_, err = fmt.Scanln()
	if err != nil {
		logger.Error(err)
	}
}
