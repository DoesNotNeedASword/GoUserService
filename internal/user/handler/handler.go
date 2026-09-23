package handler

import (
	"Test2/internal/user/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service Service
}

func New(service Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := h.service.GetUsers(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}
	response := UsersListResponse{ID: user.ID, Name: user.Name}

	c.IndentedJSON(http.StatusOK, response)
}

type Service interface {
	GetUsers(context.Context) (model.User, error)
}
