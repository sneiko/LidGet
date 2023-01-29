package api

import (
	cors "github.com/itsjamie/gin-cors"
	"lidget/api/internal"
	"lidget/api/internal/handler"
	"lidget/domain/config"
	"lidget/domain/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(cfg *config.Config, db *pgxpool.Pool) {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	repositories := repository.NewRepositories(db)

	handlers := handler.NewHandlers(repositories)
	handlers.BuildRoutes(router)

	if err := internal.RunServer(cfg, router); err != nil {
		panic(err)
	}
}
