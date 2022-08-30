package handler

import (
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

	router.LoadHTMLGlob("/home/valeriya/Документы/GitHub/http-server/pkg/handler/templates/*.html")
	search := router.GET("/search", h.searchHandler)
	api := router.Group("/api")
	{
		//api.POST("/", h.createOrder)
		api.GET("/:id", h.getOrderByID)
	}

	search.Static("/templates", "/home/valeriya/Документы/GitHub/http-server/pkg/handler/templates/")

	msg, err := handler.Subscription()
	if err != nil {
	}

	if order, err := validation(msg); err == nil {
		h.createOrder(order)
	}
	return router
}
