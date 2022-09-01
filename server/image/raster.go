package image

import (
	"errors"
	"sync"

	"github.com/h2non/bimg"
)

type ImageFormat string

type ImageConverter struct {
	mutex sync.Mutex
}

func (c *ImageConverter) ConvertSvgToBinary(svgData []byte, format bimg.ImageType, size uint) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
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
