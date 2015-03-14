package xor

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var bs string = "1c0111001f010100061a024b53535009181c"
var ks string = "686974207468652062756c6c277320657965"
var es string = "746865206b696420646f6e277420706c6179"

func TestFixed(t *testing.T) {
	b := make([]byte, hex.DecodedLen(len(bs)))
	_, err := hex.Decode(b, []byte(bs))
	if err != nil {
		t.Error(err)
	}

	k := make([]byte, hex.DecodedLen(len(ks)))
	_, err = hex.Decode(k, []byte(ks))
	if err != nil {
		t.Error(err)
	}

	e := make([]byte, hex.DecodedLen(len(es)))
	_, err = hex.Decode(e, []byte(es))
	if err != nil {
		t.Error(err)
	}

	r, err := Fixed(b, k)
	if err != nil {
		t.Error(err)
	}

	t.Logf("b: '%s'", b)
	t.Logf("k: '%s'", k)
	t.Logf("Expect: '%s'", e)
	t.Logf("Result: '%s'", string(r))

	if bytes.Compare(r, []byte(r)) != 0 {
		t.Error("output does not match expected")
	}
}
