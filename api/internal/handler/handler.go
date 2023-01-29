package handler

import (
	"github.com/gin-gonic/gin"
	"lidget/domain/repository"
)

type Handlers struct {
	Requests RequestsHandler
}

func NewHandlers(repos repository.Repositories) Handlers {
	return Handlers{
		Requests: NewRequestsHandler(repos),
	}
}

func (h *Handlers) BuildRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Status": "Running"})
	})

	api := router.Group("api")
	{
		requests := api.Group("requests")
		{
			requests.GET("/all", h.Requests.GetAll)
		}

		categories := api.Group("categories")
		{
			categories.GET("/all", h.Requests.GetAll)
		}
	}
}
