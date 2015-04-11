package xor

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestFixed(t *testing.T) {
	var bs string = "1c0111001f010100061a024b53535009181c"
	var ks string = "686974207468652062756c6c277320657965"
	var es string = "746865206b696420646f6e277420706c6179"

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

func TestRepeatingKey(t *testing.T) {
	pt := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	ct_exp := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	ct := RepeatingKey([]byte(pt), []byte(key))

	ct_string := make([]byte, hex.EncodedLen(len(ct)))
	hex.Encode(ct_string, ct)

	if bytes.Compare([]byte(ct_exp), ct_string) != 0 {
		t.Logf("expect = '%s'\n", ct_exp)
		t.Logf("actual = '%s'\n", ct_string)
		t.Fatal("xor: repeating key output does not match expected")
	}
}
