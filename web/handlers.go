package web

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "Hello")
}

func ApiV1Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "0.1.0",
		"routes":  []string{"fill me in"},
	})
}
