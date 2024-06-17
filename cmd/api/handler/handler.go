package handler

import (
	"go-boilerplate/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	ctx := c.Request.Context()

	logger.Info(ctx, "Hit API Ping")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Panic(c *gin.Context) {
	panic("panic example")
}
