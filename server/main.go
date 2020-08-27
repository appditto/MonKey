package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/appditto/monKey/server/controller"
	"github.com/appditto/monKey/server/image"
	"github.com/appditto/monKey/server/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "https://monkey.banano.cc"
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

// Generate random files
func RandFiles(count int, seed string) {
	if _, err := os.Stat("randsvg"); os.IsNotExist(err) {
		os.Mkdir("randsvg", os.FileMode(0755))
	}
	for i := 0; i < count; i++ {
		address := utils.GenerateAddress()
		sha256 := utils.Sha256(address, seed)

		accessories, _ := image.GetAccessoriesForHash(sha256)
		svg, _ := image.CombineSVG(accessories)
		ioutil.WriteFile(fmt.Sprintf("randsvg/%s.svg", address), svg, os.FileMode(0644))
	}
}

func main() {
	// Get seed from env
	seed := utils.GetEnv("MONKEY_SEED", "1234567890")

	// Server options
	serverHost := flag.String("host", "127.0.0.1", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	testAccessoryDistribution := flag.Bool("test-ad", false, "Test accessory distribution")
	randomFiles := flag.Int("rand-files", -1, "Generate this many random SVGs and output to randsvg folder")
	flag.Parse()

	if *testAccessoryDistribution {
		controller.TestAccessoryDistribution(seed)
		return
	} else if *randomFiles > 0 {
		fmt.Printf("Generating %d files in ./randsvg", *randomFiles)
		RandFiles(*randomFiles, seed)
		return
	}

	// Setup channel for stats processing job
	statsChan := make(chan *gin.Context, 100)

	// Setup imagemagick
	// Setup magickwand
	imagick.Initialize()
	defer imagick.Terminate()

	// Setup router
	router := gin.Default()
	router.Use(CorsMiddleware())

	// Setup natricon controller
	monkeyController := controller.MonkeyController{
		Seed:         seed,
		StatsChannel: &statsChan,
	}

	// V1 API
	router.GET("/api/v1/:address", monkeyController.GetBanano)
	// Stats
	router.GET("/api/stats", controller.Stats)
	router.GET("/api/stats/monthly", controller.StatsMonthly)
	router.GET("/api/random", monkeyController.GetRandomSvg)

	// Start stats worker
	go controller.StatsWorker(statsChan)

	// Run on 8080
	router.Run(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
