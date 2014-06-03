package image

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/rishihahs/qrmatrix/datatype"
)

// Margin is the number of margin modules
const Margin = 4

var white = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
var black = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}

// CreateImage creates the QR Code Matrix image
func CreateImage(moduleSize, codesPerRow int, location string) {
	// Create a new image
	size := datatype.Size()
	singlewidth := datatype.Code().Width*moduleSize + Margin*2*moduleSize
	totalwidth := singlewidth * codesPerRow
	totalheight := singlewidth * (size/codesPerRow + size%2)
	img := image.NewRGBA(image.Rect(0, 0, totalwidth, totalheight))

	// Fill the image with white
	fill := uint8(255)
	for i := 0; i < len(img.Pix); i++ {
		img.Pix[i] = fill
	}

	width := datatype.Code().Width

	for count := 0; datatype.Code() != nil; count++ {

		for n := 0; n < codesPerRow && datatype.Code() != nil; n++ {
			data := datatype.Code().Data

			for i := 0; i < len(data); i++ {
				var colour color.RGBA
				if data[i]&1 == 0 {
					colour = white
					continue
				} else {
					colour = black
				}

				for y := 0; y < moduleSize; y++ {
					for x := 0; x < moduleSize; x++ {
						img.SetRGBA(singlewidth*n+Margin*moduleSize+(i%width)*moduleSize+x, singlewidth*count+Margin*moduleSize+(i/width)*moduleSize+y, colour)
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

	fo, err := os.Create(location)
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
