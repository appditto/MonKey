package controller

import (
	"net/http"

	"github.com/appditto/monKey/server/image"
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

	accessories, err := image.GetAccessoriesForHash(sha256)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	svg, err := image.CombineSVG(accessories)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error occured")
		return
	}
	c.Data(200, "image/svg+xml; charset=utf-8", svg)
}
