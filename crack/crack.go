// Crack single-byte XOR cipher

package crack

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"errors"

	"github.com/jlubawy/go-crypto/common"
	"github.com/jlubawy/go-crypto/score"
	"github.com/jlubawy/go-crypto/xor"
)

const (
	MAX_KEY_SIZE = 40
)

func SingleByteXOR(b []byte) (byte, []byte, error) {
	highScore := 0
	highIndex := 0
	var highOutput []byte

	if len(b) == 0 {
		return 0, []byte{}, errors.New("crack: xor: buffer length must be greater than zero")
	}

	for i := 0; i < 255; i++ {
		o, err := xor.SingleByte(b, byte(i))
		if err != nil {
			return 0, []byte{}, err
		}

		s := score.LetterFrequency(string(o))

		if s > highScore {
			highScore = s
			highIndex = i
			highOutput = o
		}
	}

	return byte(highIndex), highOutput, nil
}

func DetectSingleByteXOR(s *bufio.Scanner) (int, []byte, byte, []byte, error) {
	var highScore int
	var highIndex byte
	var highLine int
	var highInput string
	var highOutput []byte

	line := 1

	for s.Scan() {
		input := s.Text()

		b := make([]byte, hex.DecodedLen(len(input)))
		_, err := hex.Decode(b, []byte(input))
		if err != nil {
			return 0, []byte{}, 0, []byte{}, err
		}

		key, plaintext, err := SingleByteXOR(b)
		if err != nil {
			return 0, []byte{}, 0, []byte{}, err
		}


		if err = s.Err(); err != nil {
			return 0, []byte{}, 0, []byte{}, err
		}

		s := score.LetterFrequency(string(plaintext))
		if s > highScore {
			highScore = s
			highIndex = key
			highLine = line
			highInput = input
			highOutput = plaintext
		}

		line += 1
	}

	return highLine, []byte(highInput), highIndex, highOutput, nil
}

// RepeatingKeyXOR takes a repeating-key-XOR buffer and returns the cracked key


// GuessKeySize uses the Hamming distance to guess the possible key size.
// Assumes that the buffer is at least greater than 2x the KEYSIZE
func GuessKeySize(b []byte) (int, error) {
	var maxKeyLen int
	var b []byte

	bestHamming := 0
	bestKeySize := 0

	maxKeyLen = len(b) / 2
	for i := 0; i < maxKeyLen; i++ {
		hamming, err := common.HammingDistance(b[i*maxKeyLen:i], b[maxKeyLen:2*maxKeyLen-1])
	}

	maxKeySize := len(b) / 2
	for i := 1; i < maxKeySize

	fmt.Println(maxKeyLen)

	for i := 0; i < maxKeyLen; i++ {

		if err != nil {
			return 0, err
		}

		if hamming > bestHamming {
			bestHamming = hamming
			bestKeySize = i
		}
	}

	return bestKeySize, nil
}