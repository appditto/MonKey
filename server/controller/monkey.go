package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/appditto/monKey/server/image"
	"github.com/appditto/monKey/server/utils"
	"github.com/gin-gonic/gin"
)

const defaultRasterSize = 128 // Default size of PNG/WEBP images
const minConvertedSize = 100  // Minimum size of PNG/WEBP converted output
const maxConvertedSize = 1000 // Maximum size of PNG/WEBP converted outpu

type MonkeyController struct {
	Seed string
}

// Return monKey for given address
func (mc MonkeyController) GetBanano(c *gin.Context) {
	address := c.Query("address")

	valid := utils.ValidateAddress(address)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid address")
		return
	}

	// See if this is a vanity
	vanity := image.GetAssets().GetVanityAsset(address)
	if vanity != nil {
		generateVanityAsset(vanity, c)
		return
	}

	pubKey := utils.AddressToPub(address)
	sha256 := utils.Sha256(pubKey, mc.Seed)

	generateIcon(&sha256, c)
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

// Generate monKey with given hash
func generateIcon(hash *string, c *gin.Context) {
	var err error

	format := strings.ToLower(c.Query("format"))
	size := 0
	if format == "" || format == "svg" {
		format = "svg"
	} else if format != "png" && format != "webp" {
		c.String(http.StatusBadRequest, "%s", "Valid formats are 'svg', 'png', or 'webp'")
		return
	} else {
		sizeStr := c.Query("size")
		if sizeStr == "" {
			size = defaultRasterSize
		} else {
			size, err = strconv.Atoi(c.Query("size"))
			if err != nil || size < minConvertedSize || size > maxConvertedSize {
				c.String(http.StatusBadRequest, "%s", fmt.Sprintf("size must be an integer between %d and %d", minConvertedSize, maxConvertedSize))
				return
			}
		}
	}

	accessories, err := image.GetAccessoriesForHash(*hash)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	svg, err := image.CombineSVG(accessories)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error occured")
		return
	}
	if format != "svg" {
		// Convert
		var converted []byte
		converted, err = image.ConvertSvgToBinary(svg, image.ImageFormat(format), uint(size))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error occured")
			return
		}
		c.Data(200, fmt.Sprintf("image/%s", format), converted)
		return
	}
	c.Data(200, "image/svg+xml; charset=utf-8", svg)
}

// Return vanity with given options
func generateVanityAsset(vanity []byte, c *gin.Context) {
	var err error

	format := strings.ToLower(c.Query("format"))
	size := 0
	if format == "" || format == "svg" {
		format = "svg"
	} else if format != "png" && format != "webp" {
		c.String(http.StatusBadRequest, "%s", "Valid formats are 'svg', 'png', or 'webp'")
		return
	} else {
		sizeStr := c.Query("size")
		if sizeStr == "" {
			size = defaultRasterSize
		} else {
			size, err = strconv.Atoi(c.Query("size"))
			if err != nil || size < minConvertedSize || size > maxConvertedSize {
				c.String(http.StatusBadRequest, "%s", fmt.Sprintf("size must be an integer between %d and %d", minConvertedSize, maxConvertedSize))
				return
			}
		}
	}

	svg, err := image.PureSVG(vanity)

	if format != "svg" {
		// Convert
		var converted []byte
		converted, err = image.ConvertSvgToBinary(svg, image.ImageFormat(format), uint(size))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error occured")
			return
		}
		c.Data(200, fmt.Sprintf("image/%s", format), converted)
		return
	}
	c.Data(200, "image/svg+xml; charset=utf-8", svg)
}
