package router

import (
	"go-boilerplate/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.MainHandler) {
	r.GET("/ping", handler.Ping)
	r.GET("/panic", handler.Panic)

	api := r.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")

	h.Book.Mount(v1)
	h.Book.Mount(v2)
}
