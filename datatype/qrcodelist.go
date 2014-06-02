package datatype

/*
#cgo pkg-config: libqrencode
#include <stdlib.h>
#include <qrencode.h>
*/
import "C"

import "unsafe"

var qrcodeList *C.QRcode_List

// CreateQRCodes creates multiple QR Codes using structured append
func CreateQRCodes(data []byte) {
	qrcodeList = C.QRcode_encodeDataStructured(C.int(len(data)), (*C.uchar)(unsafe.Pointer(&data[0])), C.int(40), C.QR_ECLEVEL_L)
}

// Next increments QRcode_List pointer to the next element
func Next() {
	qrcodeList = (*C.QRcode_List)(qrcodeList.next)
}

// Code returns a QRcode struct of the QR Code
func Code() *QRcode {
	return CreateQRCode(qrcodeList.code)
}

// Free frees the QR Code List
func Free() {
	C.QRcode_List_free(qrcodeList)
}
