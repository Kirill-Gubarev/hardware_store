package service

import (
	"github.com/gin-gonic/gin"
)

func respondError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}

func respondSuccess(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
	})
}
