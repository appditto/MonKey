package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/appditto/MonKey/server/image"
	"github.com/appditto/MonKey/server/utils"
	"github.com/gofiber/fiber/v2"
)

const defaultRasterSize = 128 // Default size of PNG/WEBP images
const minConvertedSize = 100  // Minimum size of PNG/WEBP converted output
const maxConvertedSize = 1000 // Maximum size of PNG/WEBP converted outpu

type MonkeyController struct {
	Seed string
	// StatsChannel *chan *gin.Context
}

// Return monKey for given address
func (mc MonkeyController) GetBanano(c *fiber.Ctx) error {
	address := c.Params("address")

	valid := utils.ValidateAddress(address)
	if !valid {
		return c.Status(http.StatusBadRequest).SendString("Invalid address")
	}

	// Parse stats
	//*mc.StatsChannel <- c

	// See if this is a vanity∂
	vanity := image.GetAssets().GetVanityAsset(address)
	if vanity != nil {
		return generateVanityAsset(vanity, c)
	}

	pubKey := utils.AddressToPub(address)
	sha256 := utils.Sha256(pubKey, mc.Seed)

	generateIcon(&sha256, c)
	return nil
}

// Testing APIs
func (mc MonkeyController) GetRandomSvg(c *fiber.Ctx) error {
	address := utils.GenerateAddress()
	sha256 := utils.Sha256(address, mc.Seed)

	accessories, err := image.GetAccessoriesForHash(sha256, false)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	svg, err := image.CombineSVG(accessories)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error occured")
	}
	c.Set("Content-Type", "image/svg+xml; charset=utf-8")
	return c.Status(http.StatusOK).SendStream(bytes.NewReader(svg))
}

// Generate monKey with given hash
func generateIcon(hash *string, c *fiber.Ctx) error {
	var err error

	format := strings.ToLower(c.Query("format"))
	size := 0
	if format == "" || format == "svg" {
		format = "svg"
	} else if format != "png" && format != "webp" {
		return c.Status(http.StatusBadRequest).SendString("Valid formats are 'svg', 'png', or 'webp'")
	} else {
		sizeStr := c.Query("size")
		if sizeStr == "" {
			size = defaultRasterSize
		} else {
			size, err = strconv.Atoi(c.Query("size"))
			if err != nil || size < minConvertedSize || size > maxConvertedSize {
				return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("size must be an integer between %d and %d", minConvertedSize, maxConvertedSize))
			}
		}
	}

	withBackground := strings.ToLower(c.Query("background")) == "true"

	accessories, err := image.GetAccessoriesForHash(*hash, withBackground)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	svg, err := image.CombineSVG(accessories)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error occured")
	}
	if format != "svg" {
		// Convert
		var converted []byte
		converted, err = image.ConvertSvgToBinary(svg, image.ImageFormat(format), uint(size))
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error occured")
		}
		c.Set("Content-Type", fmt.Sprintf("image/%s", format))
		return c.Status(http.StatusOK).SendStream(bytes.NewReader(converted))
	}
	c.Set("Content-Type", "image/svg+xml; charset=utf-8")
	return c.Status(http.StatusOK).SendStream(bytes.NewReader(svg))
}

// Return vanity with given options
func generateVanityAsset(vanity *image.Asset, c *fiber.Ctx) error {
	var err error

	format := strings.ToLower(c.Query("format"))
	size := 0
	if format == "" || format == "svg" {
		format = "svg"
	} else if format != "png" && format != "webp" {
		return c.Status(http.StatusBadRequest).SendString("Valid formats are 'svg', 'png', or 'webp'")
	} else {
		sizeStr := c.Query("size")
		if sizeStr == "" {
			size = defaultRasterSize
		} else {
			size, err = strconv.Atoi(c.Query("size"))
			if err != nil || size < minConvertedSize || size > maxConvertedSize {
				return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("size must be an integer between %d and %d", minConvertedSize, maxConvertedSize))
			}
		}
	}

	withBackground := strings.ToLower(c.Query("background")) == "true"

	svg, err := image.PureSVG(vanity, withBackground)

	if format != "svg" {
		// Convert
		var converted []byte
		converted, err = image.ConvertSvgToBinary(svg, image.ImageFormat(format), uint(size))
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error occured")
		}
		c.Set("Content-Type", fmt.Sprintf("image/%s", format))
		return c.Status(http.StatusOK).SendStream(bytes.NewReader(converted))
	}
	c.Set("Content-Type", "image/svg+xml; charset=utf-8")
	return c.Status(http.StatusOK).SendStream(bytes.NewReader(svg))
}

type MonkeyStatsRequest struct {
	Addresses []string `json:"addresses"`
}

type MonkeyStatsResponseItem map[string]map[string]string

// Info about a MonKey
func (mc MonkeyController) MonkeyStats(c *fiber.Ctx) error {
	var reqJson MonkeyStatsRequest

	ret := make(MonkeyStatsResponseItem)

	for _, address := range reqJson.Addresses {
		if !utils.ValidateAddress(address) {
			return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Invalid address in address list %s", address))
		}
		// Get monkey info
		pubKey := utils.AddressToPub(address)
		sha256 := utils.Sha256(pubKey, mc.Seed)
		accessories, _ := image.GetAccessoriesForHash(sha256, true)

		ret[address] = make(map[string]string)
		ret[address]["background_color"] = accessories.BGColor
		if accessories.GlassesAsset != nil {
			ret[address]["glasses"] = accessories.GlassesAsset.FileName
		} else {
			ret[address]["glasses"] = "none"
		}
		if accessories.HatAsset != nil {
			ret[address]["hat"] = accessories.HatAsset.FileName
		} else {
			ret[address]["hat"] = "none"
		}
		if accessories.MiscAsset != nil {
			ret[address]["misc"] = accessories.MiscAsset.FileName
		} else {
			ret[address]["misc"] = "none"
		}
		if accessories.MouthAsset != nil {
			ret[address]["mouth"] = accessories.MouthAsset.FileName
		} else {
			ret[address]["mouth"] = "none"
		}
		if accessories.ShirtPantsAsset != nil {
			ret[address]["shirt_pants"] = accessories.ShirtPantsAsset.FileName
		} else {
			ret[address]["shirt_pants"] = "none"
		}
		if accessories.ShoeAsset != nil {
			ret[address]["shoes"] = accessories.ShoeAsset.FileName
		} else {
			ret[address]["shoes"] = "none"
		}
		if accessories.TailAccessory != nil {
			ret[address]["tail_accessory"] = accessories.TailAccessory.FileName
		} else {
			ret[address]["tail_accessory"] = "none"
		}
		for k, v := range accessories.AccessoryColors {
			ret[address][fmt.Sprintf("color_%s", k)] = v.ToHTML(true)
		}
	}

	return c.Status(http.StatusOK).JSON(ret)
}
