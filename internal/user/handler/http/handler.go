package http

import (
	"Test2/internal/user/model"
	"context"
	"net/http"
	"strconv"

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
	users, err := h.service.GetUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	response := UsersListResponse{Users: users}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ctx := c.Request.Context()
	user, err := h.service.GetUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	response := UserResponse{User: *user}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := req.ToModel()
	user, err := h.service.CreateUser(ctx, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusCreated, user)
}

type Service interface {
	GetUsers(context.Context) ([]model.User, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, u model.CreateUser) (*model.User, error)
}
