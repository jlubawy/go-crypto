// XOR - all XOR related functions

package xor

import (
	"errors"
)

// Fixed takes two equal-length buffers and produces their XOR combination
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

// SingleByte takes a byte key and XORs it with byte slice b
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

// RepeatingKey XORs pt with key, wrapping every len(key)
func RepeatingKey(pt []byte, key []byte) []byte {
	if len(pt) == 0 {
		return []byte{}
	}

	ct := make([]byte, len(pt))
	kl := len(key)

	for i, _ := range pt {
		ct[i] = key[i % kl] ^ pt[i]
	}

	return ct
}
