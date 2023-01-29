package main

import (
	"context"
	"lidget/api/pkg/api"
	appDb "lidget/api/pkg/db"
	"lidget/domain/config"
	"lidget/tgClient"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := appDb.Connect(cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		client := tgClient.NewTgClient(&cfg.Tg, db)

		err = client.Run(context.Background())
		if err != nil {
			return
		}
	}()

	api.Run(cfg, db)
}
