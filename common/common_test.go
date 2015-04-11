package common

import (
    "testing"
)

func TestByteCountOnes(t *testing.T) {
    var full byte = 0xFF;
    var empty byte = 0;

    if ByteCountOnes(full) != 8 {
        t.Fatal("common: ByteCountOnes for 0xFF did not equal 8")
    }

    if ByteCountOnes(empty) != 0 {
        t.Fatal("common: ByteCountOnes for 0x00 did not equal 0")
    }
}

func TestHammingDistance(t *testing.T) {
    a_str := "this is a test"
    b_str := "wokka wokka!!!"
    h_dist_exp := 37

    h_dist, err := HammingDistance([]byte(a_str), []byte(b_str))
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("expect = %d\n", h_dist_exp)
    t.Logf("actual = %d\n", h_dist)

    if h_dist != h_dist_exp {
        t.Fatal("common: hamming distance output does not equal expected")
    }
}
