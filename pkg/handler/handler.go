package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handler "http_server/pkg/handler/nats-streaming"
	"http_server/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		//api.POST("/", h.createOrder)
		api.GET("/:id", h.getOrderByID)
	}

	msg, err := handler.Subscription()
	if err != nil {
	}

	if order, err := validation(msg); err == nil {
		fmt.Println("hehe")
		h.createOrder(order)
	}
	fmt.Println()
	return router
}
