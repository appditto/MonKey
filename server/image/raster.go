package image

import (
	"github.com/davidbyttow/govips/v2/vips"
)

type ImageFormat string

func ConvertSvgToBinary(svgData []byte, format ImageFormat, size uint) ([]byte, error) {
	defer vips.ClearCache()
	inputImage, err := vips.NewImageFromBuffer(svgData)
	if err != nil {
		return nil, err
	}
	defer inputImage.Close()

	inputImage.Resize(float64(size)/float64(DefaultSize), vips.KernelAuto)

	if format == "png" {
		ep := vips.NewPngExportParams()
		ep.StripMetadata = true
		imageBytes, _, err := inputImage.ExportPng(ep)
		if err != nil {
			return nil, err
		}
		return imageBytes, nil
	}
	ep := vips.NewWebpExportParams()
	ep.StripMetadata = true
	imageBytes, _, err := inputImage.ExportWebp(ep)
	if err != nil {
		return nil, err
	}
	return imageBytes, nil
}
