package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vision/product-service/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		Service: service.NewService(),
	}
}

func (h *Handler) GetProducts(c *gin.Context) {
	products := h.Service.GetProducts()
	c.JSON(http.StatusOK, products)
}
