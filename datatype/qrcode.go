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
	Version, Width int
	Data           []byte
}

// CreateQRCode creates a Go QR Code from a C QR Code
func CreateQRCode(cqrcode *C.QRcode) *QRcode {
	if cqrcode == nil {
		return nil
	}

	version := int(cqrcode.version)
	width := int(cqrcode.width)
	data := C.GoBytes(unsafe.Pointer(cqrcode.data), C.int(width*width))

	return &QRcode{Version: version, Width: width, Data: data}
}
