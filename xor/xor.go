// XOR - all XOR related functions

package xor

import (
	"errors"
)

// Takes two equal-length buffers and produces their XOR combination
func Fixed(b, k []byte) ([]byte, error) {
	if len(b) != len(k) {
		return nil, errors.New("xor: buffer length and key length must be the same length")
	}

	o := make([]byte, len(b))

	for i := 0; i < len(b); i++ {
		o[i] = b[i] ^ k[i]
	}

	return o, nil
}

func SingleByte(b []byte, k byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, errors.New("xor: buffer length must be greater than zero")
	}

	o := make([]byte, len(b))

	for i, _ := range b {
		o[i] = b[i] ^ k
	}

	return o, nil
}
