package datatype

/*
#cgo pkg-config: libqrencode
#include <stdlib.h>
#include <qrencode.h>
*/
import "C"

import "unsafe"

var origList *C.QRcode_List
var qrcodeList *C.QRcode_List
var qrcode *QRcode

// CreateQRCodes creates multiple QR Codes using structured append
func CreateQRCodes(data []byte) {
	qrcodeList = C.QRcode_encodeDataStructured(C.int(len(data)), (*C.uchar)(unsafe.Pointer(&data[0])), C.int(40), C.QR_ECLEVEL_L)
	origList = qrcodeList
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
