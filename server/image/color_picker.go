package image

import (
	"math"
	"strconv"

	"github.com/appditto/monKey/server/color"
)

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
	) * 1000
	upperBound := math.Min(
		math.Sqrt(
			math.Max(
				(MaxPerceivedBrightness255*MaxPerceivedBrightness255-color.RedPBMultiplier*outRGB.R*outRGB.R-color.GreenPBMultiplier*outRGB.G*outRGB.G)/color.BluePBMultiplier,
				0.0,
			),
		),
		255.0,
	) * 1000

	delta := upperBound - lowerBound

	outRGB.B = float64(math.Mod(float64(blueAsInt), delta)) + lowerBound

	return outRGB, nil
}
