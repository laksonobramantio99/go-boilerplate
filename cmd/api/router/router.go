package router

import (
	"go-boilerplate/cmd/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	r.GET("/panic", handler.Panic)
}
