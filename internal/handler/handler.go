package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/middleware"
	"server/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	router.Use(middleware.Cors())

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
	}

	api := router.Group("/api", h.AuthMiddleware)
	{
		logs := api.Group("/logs")
		{
			logs.POST("/", h.createLog)
			logs.GET("/:date", h.getAllLogs)
			logs.GET("/log/:id", h.getLogById)
			logs.PUT("/:id", h.updateLog)
			logs.DELETE("/:id", h.deleteLog)
		}
		food := api.Group("/food")
		{
			food.GET("/", h.getAllFood)
		}
	}

	return router
}