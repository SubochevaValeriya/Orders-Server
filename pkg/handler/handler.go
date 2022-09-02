package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"http_server/pkg/service"
)

type Handler struct {
	services      *service.Service
	natsStreaming stan.Conn
}

func NewHandler(services *service.Service, natsStreaming stan.Conn) *Handler {
	return &Handler{
		services:      services,
		natsStreaming: natsStreaming}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//Loading HTML Template
	router.LoadHTMLGlob("./templates/*.html")
	search := router.GET("/search", h.searchHandler)

	//Adding CSS template
	search.Static("/templates", "./templates/")

	//Adding API endpoint
	api := router.Group("/api")
	{
		api.GET("", h.getOrderByID)
	}

	//Subscribe to a channel
	h.natsStreaming.Subscribe(viper.GetString("nats-streaming.channel_name"),
		func(m *stan.Msg) {
			if order, err := validation(m.Data); err == nil {
				h.createOrder(order)
			}
			m.Ack()
		},
		stan.StartWithLastReceived())

	return router
}
