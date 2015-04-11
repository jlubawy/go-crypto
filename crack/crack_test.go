package crack

import (
	"bufio"
	"encoding/hex"
	"os"
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

	t.Logf("key    = '%c (0x%X)'\n", k, k)
	t.Logf("output = '%s'\n", string(o))
}

func TestDetectSingleByteXOR(t *testing.T) {
	test_data, err := os.Open("test_data/challenge-4-test-data.txt")
	if err != nil {
		t.Fatal(err)
	}

	s := bufio.NewScanner(test_data)
	line, ct, key, pt, err := DetectSingleByteXOR(s)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("line   = %d\n", line)
	t.Logf("input  = '%s'\n", string(ct))
	t.Logf("key    = '%c (0x%X)'\n", key, key)
	t.Logf("output = '%s'\n", string(pt))
}

func TestGuessKeySize(t *testing.T) {
	key_str := "ICE"
	ct_str := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	k := []byte(key_str)

	ct := make([]byte, hex.DecodedLen(len(ct_str)))
	hex.Decode(ct, []byte(ct_str))

	t.Logf("expected key len = %d\n", len(k))
	t.Logf("  ciphertext len = %d\n", len(ct))

	keySize, err := GuessKeySize(ct)
	if err != nil {
		t.Fatal(err)
	}

	if keySize != len(k) {
		t.Fatalf("crack: keysize should be %d but was %d\n", len(k), keySize)
	}
}