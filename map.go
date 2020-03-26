package map12381

/*
#cgo LDFLAGS: -L./map-rs/lib -lkyberrs
#include <stdlib.h>
#include "./map-rs/lib/kyberrs.h"
*/
import "C"
import "unsafe"

// BLS_SIG_BLS12381G2-SHA256-SSWU-RO-_NUL_
var csuite = []byte{66, 76, 83, 95, 83, 73, 71, 95, 66, 76, 83, 49, 50, 51, 56, 49, 71, 50, 45, 83, 72, 65, 50, 53, 54, 45, 83, 83, 87, 85, 45, 82, 79, 45, 95, 78, 85, 76, 95}

func MapToG2(msg []byte, domain []byte) []byte {
	dst := make([]byte, 96)
	pointPtr := (*C.char)(unsafe.Pointer(&dst[0]))
	if domain == nil {
		domain = csuite
	}
	domainPtr := (*C.char)(unsafe.Pointer(&domain[0]))
	domainLen := C.size_t(len(domain))
	msgPtr := (*C.char)(unsafe.Pointer(&msg[0]))
	msgLen := C.size_t(len(msg))
	C.maptopointg2(msgPtr, msgLen, pointPtr, domainPtr, domainLen)
	return dst
}

func MapToG1(msg []byte, domain []byte) []byte {
	dst := make([]byte, 48)
	pointPtr := (*C.char)(unsafe.Pointer(&dst[0]))
	if domain == nil {
		domain = csuite
	}
	domainPtr := (*C.char)(unsafe.Pointer(&domain[0]))
	domainLen := C.size_t(len(domain))
	msgPtr := (*C.char)(unsafe.Pointer(&msg[0]))
	msgLen := C.size_t(len(msg))
	C.maptopointg1(msgPtr, msgLen, pointPtr, domainPtr, domainLen)
	return dst
}
