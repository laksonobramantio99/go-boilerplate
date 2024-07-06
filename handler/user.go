package handler

import (
	"go-boilerplate/util/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	logger.Info(ctx, "Hit API Get user")
	c.JSON(http.StatusOK, gin.H{
		"message": "test done",
	})
}
