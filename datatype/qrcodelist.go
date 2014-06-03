package datatype

/*
#cgo pkg-config: libqrencode
#include <stdlib.h>
#include <qrencode.h>
*/
import "C"

import (
	"syscall"
	"unsafe"
)

var origList *C.QRcode_List
var qrcodeList *C.QRcode_List
var qrcode *QRcode

// CreateQRCodes creates multiple QR Codes using structured append
func CreateQRCodes(data []byte) {
	a, err := C.QRcode_encodeDataStructured(C.int(len(data)), (*C.uchar)(unsafe.Pointer(&data[0])), C.int(40), C.QR_ECLEVEL_L)

	// If data too large
	if err == syscall.ERANGE {
		split(data)
	} else {
		qrcodeList = a
		origList = qrcodeList
	}
}

// Split list and recursively build onto a new one
func split(data []byte) {
	size := len(data) / 2
	var mainList *C.QRcode_List
	CreateQRCodes(data[:size])
	mainList = qrcodeList
	tempList := mainList

	CreateQRCodes(data[size-1 : len(data)])
	for tempList.next != nil {
		tempList = (*C.QRcode_List)(tempList.next)
	}
	for qrcodeList != nil {
		tempList.next = qrcodeList.next
		tempList = (*C.QRcode_List)(tempList.next)
		qrcodeList = (*C.QRcode_List)(qrcodeList.next)
	}
	qrcodeList = mainList
	origList = mainList
}

// Next increments QRcode_List pointer to the next element
func Next() {
	qrcodeList = (*C.QRcode_List)(qrcodeList.next)
	qrcode = nil
	Code()
}

// Code returns a QRcode struct of the QR Code
func Code() *QRcode {
	if qrcodeList == nil {
		return nil
	}

	if qrcode == nil {
		qrcode = CreateQRCode(qrcodeList.code)
	}

	return qrcode
}

// Size returns the size of the QR Code List
func Size() int {
	return int(C.QRcode_List_size(qrcodeList))
}

// Free frees the QR Code List
func Free() {
	C.QRcode_List_free(origList)
}
