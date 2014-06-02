package datatype

/*
#cgo pkg-config: libqrencode
#include <stdlib.h>
#include <qrencode.h>
*/
import "C"

import "unsafe"

// QRcode struct represents a QR Code
type QRcode struct {
	version, width int
	data           []byte
}

// CreateQRCode creates a Go QR Code from a C QR Code
func CreateQRCode(cqrcode *C.QRcode) *QRcode {
	version := int(cqrcode.version)
	width := int(cqrcode.width)
	data := C.GoBytes(unsafe.Pointer(cqrcode.data), C.int(width*width))

	return &QRcode{version: version, width: width, data: data}
}
