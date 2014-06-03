package image

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/rishihahs/qrmatrix/datatype"
)

const (
	// ModuleSize is the size of QR Code module
	ModuleSize = 1

	// Margin is the number of margin modules
	Margin = 4

	// CodesPerRow is the the number of QR Codes per row
	CodesPerRow = 2

	// DPI Dots per inch
	DPI = 72
)

var white = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
var black = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}

// CreateImage creates the QR Code Matrix image
func CreateImage() {
	img := image.NewRGBA(image.Rect(0, 0, datatype.Code().Width*ModuleSize+Margin*2*ModuleSize, datatype.Code().Width*ModuleSize+Margin*2*ModuleSize))
	fill := uint8(255)
	for i := 0; i < len(img.Pix); i++ {
		img.Pix[i] = fill
	}

	data := datatype.Code().Data
	width := datatype.Code().Width
	for i := 0; i < len(data); i++ {
		var colour color.RGBA
		if data[i]&1 == 0 {
			colour = white
		} else {
			colour = black
		}

		for y := 0; y < ModuleSize; y++ {
			for x := 0; x < ModuleSize; x++ {
				img.SetRGBA(Margin*ModuleSize+(i%width)*ModuleSize+x, Margin*ModuleSize+(i/width)*ModuleSize+y, colour)
			}
		}
	}

	fo, err := os.Create("/Users/shah/Desktop/image.png")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	writer := bufio.NewWriter(fo)
	png.Encode(writer, img)

	if err = writer.Flush(); err != nil {
		panic(err)
	}
}
