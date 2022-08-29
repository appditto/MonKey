package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/appditto/MonKey/server/controller"
	"github.com/appditto/MonKey/server/image"
	"github.com/appditto/MonKey/server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gopkg.in/gographics/imagick.v3/imagick"
	"k8s.io/klog/v2"
)

// Generate random files
func RandFiles(count int, seed string, bg bool) {
	if _, err := os.Stat("randsvg"); os.IsNotExist(err) {
		os.Mkdir("randsvg", os.FileMode(0755))
	}
	for i := 0; i < count; i++ {
		address := utils.GenerateAddress()
		sha256 := utils.Sha256(address, seed)

		accessories, _ := image.GetAccessoriesForHash(sha256, bg)
		svg, _ := image.CombineSVG(accessories)
		os.WriteFile(fmt.Sprintf("randsvg/%s.svg", address), svg, os.FileMode(0644))
	}
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// Get seed from env
	seed := utils.GetEnv("MONKEY_SEED", "1234567890")

	// Server options
	flag.Usage = usage
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	if utils.GetEnv("ENVIRONMENT", "development") == "development" {
		flag.Set("stderrthreshold", "INFO")
		flag.Set("v", "3")
	}
	serverHost := flag.String("host", "localhost", "Host to listen on")
	serverPort := flag.Int("port", 8080, "Port to listen on")
	testAccessoryDistribution := flag.Bool("test-ad", false, "Test accessory distribution")
	randomFiles := flag.Int("rand-files", -1, "Generate this many random SVGs and output to randsvg folder")
	randomFilesBG := flag.Int("rand-files-bg", -1, "Generate random SVGs with background and output to randsvg folder")
	flag.Parse()

	if *testAccessoryDistribution {
		controller.TestAccessoryDistribution(seed)
		return
	} else if *randomFiles > 0 {
		fmt.Printf("Generating %d files in ./randsvg", *randomFiles)
		RandFiles(*randomFiles, seed, false)
		return
	} else if *randomFilesBG > 0 {
		fmt.Printf("Generating %d files in ./randsvg", *randomFilesBG)
		RandFiles(*randomFilesBG, seed, true)
		return
	}

	// Setup channel for stats processing job
	// statsChan := make(chan *gin.Context, 100)

	// Setup imagemagick
	// Setup magickwand
	imagick.Initialize()
	defer imagick.Terminate()

	// Setup router
	router := fiber.New()
	// pprof.Register(router)
	router.Use(cors.New(cors.ConfigDefault))

	// Setup natricon controller
	monkeyController := controller.MonkeyController{
		Seed: seed,
		// StatsChannel: &statsChan,
	}

	// V1 API
	apiGroup := router.Group("/api/v1")
	// Stats
	apiGroup.Get("/stats", controller.Stats)
	apiGroup.Get("/stats/monthly", controller.StatsMonthly)
	// Address
	apiGroup.Get("/monkey/:address", monkeyController.GetBanano)
	apiGroup.Post("/monkey/dtl", monkeyController.MonkeyStats)
	// Testing
	if utils.GetEnv("ENVIRONMENT", "development") == "development" {
		apiGroup.Get("/random", monkeyController.GetRandomSvg)

	}

	// Start stats worker
	//go controller.StatsWorker(statsChan)

	// Run on 8080
	router.Listen(fmt.Sprintf("%s:%d", *serverHost, *serverPort))
}
