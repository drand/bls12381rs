package bls12381

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
