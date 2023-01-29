package handler

import (
	"github.com/gin-gonic/gin"
	"lidget/domain/pkg/logger"
	"lidget/domain/repository"
	"net/http"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoriesHandler struct {
	repos repository.Repositories
}

func NewCategoriesHandler(repos repository.Repositories) RequestsHandler {
	return RequestsHandler{
		repos: repos,
	}
}

func (h *CategoriesHandler) GetAll(c *gin.Context) {
	dbCats, err := h.repos.Categories.GetAll(c)
	if err != nil {
		logger.Error(err)
		c.Error(err)
	}

	result := make([]Category, len(dbCats))
	for i, item := range dbCats {
		result[i] = Category{
			Id:   item.Id,
			Name: item.Name,
		}
	}

	c.JSON(http.StatusOK, result)
}
