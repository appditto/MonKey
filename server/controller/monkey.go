package controller

import (
	"github.com/appditto/monKey/server/utils"
	"github.com/gin-gonic/gin"
)

type MonkeyController struct {
	Seed string
}

// Testing APIs
func (mc MonkeyController) GetRandomSvg(c *gin.Context) {
	address := utils.GenerateAddress()
	sha256 := utils.Sha256(address, mc.Seed)

	c.JSON(200, gin.H{
		"hello": sha256,
	})
}
