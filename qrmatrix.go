package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rishihahs/qrmatrix/datatype"
	"github.com/rishihahs/qrmatrix/image"
)

var codesPerRow int
var moduleSize int
var output string

func init() {
	flag.IntVar(&codesPerRow, "codes-per-row", 3, "[optional] number of qr codes per row")
	flag.IntVar(&moduleSize, "size", 2, "[optional] width and height (in pixels) of each module (square block of qr code)")
	flag.StringVar(&output, "output", "", "[optional] The file to output the image to (e.g. codes.png). Will output to STDOUT by default")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nData sent through STDIN will be encoded.\nE.g.\n\t%s --size 3 < fileToEncode\n\techo \"Hello World\" | %s > output.png\n", os.Args[0], os.Args[0])
	}
}

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	datatype.CreateQRCodes(bytes)
	image.CreateImage(moduleSize, codesPerRow, output)
}
