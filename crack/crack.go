// Crack single-byte XOR cipher

package crack

import (
	"errors"

	"github.com/jlubawy/go-crypto/score"
	"github.com/jlubawy/go-crypto/xor"
)

func SingleByteXOR(b []byte) (byte, []byte, error) {
	highScore := 0
	highIndex := 0

	if len(b) == 0 {
		return 0, []byte{}, errors.New("crack: xor: buffer length must be greater than zero")
	}

	highOutput := make([]byte, len(b))

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
