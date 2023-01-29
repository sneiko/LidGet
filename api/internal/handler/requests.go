package handler

import (
	"github.com/gin-gonic/gin"
	"lidget/api/internal/models/requests"
	"lidget/domain/pkg/logger"
	"lidget/domain/repository"
	"net/http"
	"time"
)

type Request struct {
	Id         int       `json:"id"`
	SenderTgId string    `json:"senderTgId"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"createdAt"`
}

type RequestsHandler struct {
	repos repository.Repositories
}

func NewRequestsHandler(repos repository.Repositories) RequestsHandler {
	return RequestsHandler{
		repos: repos,
	}
}

func (h *RequestsHandler) GetAll(c *gin.Context) {
	var model requests.GetAllFromTo
	err := c.BindQuery(&model)
	if err != nil {
		logger.Error(err)
	}

	dbReqs, err := h.repos.Requests.GetByRange(c, model.From, model.To)
	if err != nil {
		logger.Error(err)
		logger.Error(c.Error(err))
	}

	result := make([]Request, len(dbReqs))
	for i, item := range dbReqs {
		result[i] = Request{
			Id:         item.Id,
			SenderTgId: item.SenderTgId,
			Text:       item.Text,
			CreatedAt:  item.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, result)
}
