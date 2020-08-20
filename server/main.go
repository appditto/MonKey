package main

import (
	"flag"
	"fmt"

	"github.com/appditto/monKey/server/controller"
	"github.com/appditto/monKey/server/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "https://testmonkey.appditto.com"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With, ResponseType")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	// Get seed from env
	seed := utils.GetEnv("MONKEY_SEED", "1234567890")

	// Server options
	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	// Setup imagemagick
	// Setup magickwand
	imagick.Initialize()
	defer imagick.Terminate()

	// Setup router
	router := gin.Default()
	router.Use(CorsMiddleware())

	// Setup natricon controller
	monkeyController := controller.MonkeyController{
		Seed: seed,
	}

	// V1 API
	router.GET("/api/v1/:address", monkeyController.GetBanano)
	router.GET("/api/random", monkeyController.GetRandomSvg)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
