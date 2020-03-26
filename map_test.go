package main

/*
#cgo LDFLAGS: -L ${SRCDIR}/map-rs/lib -lkyberrs
#include <stdlib.h>
#include "./map-rs/lib/kyberrs.h"
*/

import "fmt"
import ffi "github.com/filecoin-project/filecoin-ffi"
import "testing"

func TestMapG2Filecoin(t *testing.T) {
	msg := []byte("hello world")
	msgFIL := ffi.Message(msg)
	digestFIL := ffi.Hash(msgFIL)
	res := MapToG2(msg, nil)
	fmt.Println(digestFIL)
	fmt.Println(res)
}
