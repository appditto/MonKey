package image

import (
	"fmt"
	"math"
	"strconv"

	"github.com/appditto/monKey/server/color"
)

// Min and max shadow opacity for fur
const MinShadowOpacityFur = 0.12
const MaxShadowOpacityFur = 0.4

// Min and max shadow opacity for fur (dark)
const MinShadowOpacityFurDark = 0.15
const MaxShadowOpacityFurDark = 0.5

// Min and max shadow opacity for eyes
const MinShadowOpacityIris = 0.075
const MaxShadowOpacityIris = 0.25

// Min and max perceivedBrightness values (between 0 and 100)
const MinPerceivedBrightness = 18.0
const MaxPerceivedBrightness = 95.0

// Min and max perceivedBrightness values (between 0 and 255)
const MinPerceivedBrightness255 = MinPerceivedBrightness / 100 * 255
const MaxPerceivedBrightness255 = MaxPerceivedBrightness / 100 * 255

// GetColor - Get color with given entropy respecting perceived brightness boundaries
func GetColor(redSeed string, greenSeed string, blueSeed string) (color.RGB, error) {
	// Parse hex scales as int
	redAsInt, err := strconv.ParseInt(redSeed, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	greenAsInt, err := strconv.ParseInt(greenSeed, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	blueAsInt, err := strconv.ParseInt(blueSeed, 16, 64)
	if err != nil {
		return color.RGB{}, err
	}
	outRGB := color.RGB{}

	// Determine Red and Green values, 0-255
	outRGB.R = math.Mod(float64(redAsInt), 255.0)
	outRGB.G = math.Mod(float64(greenAsInt), 255.0)

	// Generate Blue satisfying perceved brightness requirements
	lowerBound := math.Max(
		math.Sqrt(
			math.Max(
				(MinPerceivedBrightness255*MinPerceivedBrightness255-color.RedPBMultiplier*outRGB.R*outRGB.R-color.GreenPBMultiplier*outRGB.G*outRGB.G)/color.BluePBMultiplier,
				0.0,
			),
		),
		0.0,
	)
	upperBound := math.Min(
		math.Sqrt(
			math.Max(
				(MaxPerceivedBrightness255*MaxPerceivedBrightness255-color.RedPBMultiplier*outRGB.R*outRGB.R-color.GreenPBMultiplier*outRGB.G*outRGB.G)/color.BluePBMultiplier,
				0.0,
			),
		),
		255.0,
	)

	maxBlueStr := ""
	for i := 0; i < len(blueSeed); i++ {
		maxBlueStr += "f"
	}
	maxBlue, _ := strconv.ParseInt(maxBlueStr, 16, 64)

	outRGB.B = lowerBound + (1.0-float64(blueAsInt)/float64(maxBlue))*(upperBound-lowerBound)
	if outRGB.B < lowerBound || outRGB.B > upperBound {
		fmt.Printf("\n\nBLUE OUT OF RANGE\nLOWER BOUND %f\bUPPER BOUND %f\nACTUA BOUNDD %f\nINPUTS: %d %f %f", lowerBound, upperBound, outRGB.B, blueAsInt, outRGB.R, outRGB.G)
	}

	return outRGB, nil
}

func GetShadowOpacityFur(clr color.RGB) float64 {
	return math.Round((MinShadowOpacityFur+(1-clr.PerceivedBrightness()/100)*(MaxShadowOpacityFur-MinShadowOpacityFur))*100) / 100
}

func GetShadowOpacityFurDark(clr color.RGB) float64 {
	return math.Round((MinShadowOpacityFurDark+(1-clr.PerceivedBrightness()/100)*(MaxShadowOpacityFurDark-MinShadowOpacityFurDark))*100) / 100
}

func GetShadowOpacityIris(clr color.RGB) float64 {
	return math.Round((MinShadowOpacityIris+(1-clr.PerceivedBrightness()/100)*(MaxShadowOpacityIris-MinShadowOpacityIris))*100) / 100
}
