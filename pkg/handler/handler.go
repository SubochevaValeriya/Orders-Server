package handler

import (
	"github.com/gin-gonic/gin"
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
		api.POST("/", h.createUser)
		api.GET("/:id", h.getBalanceByID)
	}

	// но сейчас у меня всё по таблице balance, а можно ещё добавить по транзакциям, будет ещё один сет апи

	return router
}
