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
)

var white = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
var black = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}

// CreateImage creates the QR Code Matrix image
func CreateImage() {
	// Create a new image
	size := datatype.Size()
	singlewidth := datatype.Code().Width*ModuleSize + Margin*2*ModuleSize
	totalwidth := singlewidth * CodesPerRow
	totalheight := singlewidth * (size/CodesPerRow + size%2)
	img := image.NewRGBA(image.Rect(0, 0, totalwidth, totalheight))

	// Fill the image with white
	fill := uint8(255)
	for i := 0; i < len(img.Pix); i++ {
		img.Pix[i] = fill
	}

	width := datatype.Code().Width

	for count := 0; datatype.Code() != nil; count++ {

		for n := 0; n < CodesPerRow && datatype.Code() != nil; n++ {
			data := datatype.Code().Data

			for i := 0; i < len(data); i++ {
				var colour color.RGBA
				if data[i]&1 == 0 {
					colour = white
					continue
				} else {
					colour = black
				}

				for y := 0; y < ModuleSize; y++ {
					for x := 0; x < ModuleSize; x++ {
						img.SetRGBA(singlewidth*n+Margin*ModuleSize+(i%width)*ModuleSize+x, singlewidth*count+Margin*ModuleSize+(i/width)*ModuleSize+y, colour)
					}
				}
			}

			datatype.Next()
		}

		if datatype.Code() != nil {
			datatype.Next()
		}
	}

	datatype.Free()

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
