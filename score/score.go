// score - based on input return an integer score, the higher the better

package score

import (
	"strings"
)

var freq = map[string]int{
	"z": 0,
	"q": 1,
	"x": 2,
	"j": 3,
	"k": 4,
	"v": 5,
	"b": 6,
	"p": 7,
	"y": 8,
	"g": 9,
	"f": 10,
	"w": 11,
	"m": 12,
	"u": 13,
	"c": 14,
	"l": 15,
	"d": 16,
	"r": 17,
	"h": 18,
	"s": 19,
	"n": 20,
	"i": 21,
	"o": 22,
	"a": 23,
	"t": 24,
	"e": 25,
}

// Scores based on English letter-frequency
func LetterFrequency(i string) int {
	l := strings.ToLower(i)
	s := 0

	for _, v := range l {
		s += freq[string(v)]
	}

	return s
}
