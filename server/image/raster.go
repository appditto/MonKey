package image

import (
	"runtime"
	"strings"
	"sync"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type ImageFormat string

type ImageConverter struct {
	mutex sync.Mutex
}

func (c *ImageConverter) ConvertSvgToBinary(svgData []byte, format ImageFormat, size uint) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	mw.SetImageFormat("SVG")
	pixelWand := imagick.NewPixelWand()
	defer pixelWand.Clear()
	pixelWand.SetColor("none")
	mw.SetBackgroundColor(pixelWand)
	mw.SetImageUnits(imagick.RESOLUTION_PIXELS_PER_INCH)
	density := 96.0 * float64(size) / float64(DefaultSize)
	mw.SetResolution(density, density)
	err := mw.ReadImageBlob(svgData)
	if err != nil {
		return nil, err
	}
	mw.SetImageCompression(imagick.COMPRESSION_NO)
	mw.SetImageCompressionQuality(100)
	//mw.SetAntialias(true)
	mw.SetImageFormat(strings.ToUpper(string(format)))
	return mw.GetImageBlob(), nil
}
