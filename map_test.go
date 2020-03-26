package bls12381

import ffi "github.com/filecoin-project/filecoin-ffi"

func TestMapG2Filecoin(t *testing.T) {
	msg := []byte("hello world")
	msgFIL := ffi.Message(msg)
	digestFIL := ffi.Hash(msgFIL)

}
