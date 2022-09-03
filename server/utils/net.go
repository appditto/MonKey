package utils

import (
	"github.com/gin-gonic/gin"
)

func IPAddress(c *gin.Context) string {
	IPAddress := c.GetHeader("CF-Connecting-IP")
	if IPAddress == "" {
		IPAddress = c.ClientIP()
	}
	return IPAddress
}
