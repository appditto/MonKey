package image

import (
	"errors"

	"github.com/h2non/bimg"
)

type ImageFormat string

func ConvertSvgToBinary(svgData []byte, format bimg.ImageType, size uint) ([]byte, error) {
	inputImage := bimg.NewImage(svgData)
	if inputImage == nil {
		return nil, errors.New("Unable to load svg for rasterization")
	}

	op := bimg.Options{
		Width:         int(size),
		Height:        int(size),
		Type:          format,
		StripMetadata: true,
	}
	processed, err := inputImage.Process(op)
	if err != nil {
		return nil, err
	}
	return processed, nil
}
