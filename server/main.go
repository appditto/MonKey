package main

import (
	"flag"
	"fmt"

	"github.com/appditto/monKey/server/controller"
	"github.com/appditto/monKey/server/utils"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "https://natricon.com"
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
	// Parse server options
	loadFiles := flag.Bool("load-files", false, "Print assets as GO arrays")

	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	if *loadFiles {
		LoadAssetsToArray()
		return
	}

	// Setup router
	router := gin.Default()
	router.Use(CorsMiddleware())

	// Setup natricon controller
	monkeyController := controller.MonkeyController{
		Seed: seed,
	}

	// V1 API
	router.GET("/api/v1/random", monkeyController.GetRandomSvg)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
