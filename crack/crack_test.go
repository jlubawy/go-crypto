package crack

import (
	"encoding/hex"
	"testing"
)

var bs string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func TestSingleByteXOR(t *testing.T) {
	b := make([]byte, hex.DecodedLen(len(bs)))
	_, err := hex.Decode(b, []byte(bs))
	if err != nil {
		t.Error(err)
	}

	k, o, err := SingleByteXOR(b)
	if err != nil {
		t.Error(err)
	}

	t.Logf("key    = '%c'\n", k)
	t.Logf("output = '%s'\n", string(o))
}
