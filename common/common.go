package common

import (
    "github.com/jlubawy/go-crypto/xor"
)

func HammingDistance(a, b []byte) (int, error) {
    // Compute fixed-XOR, every 1 is a different byte
    c, err := xor.Fixed(a, b)
    if err != nil {
        return 0, err
    }

    distance := 0
    for _, d := range c {
        distance += ByteCountOnes(d)
    }

    return distance, nil
}

func ByteCountOnes(b byte) int {
    var i uint

    distance := 0
    for i = 0; i < 8; i++ {
        if byte(1 << i) & b != 0 {
            distance += 1
        }
    }

    return distance
}