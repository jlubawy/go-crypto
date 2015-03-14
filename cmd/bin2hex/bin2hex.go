// bin2hex - convert binary to hexadecimal string

package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)

	if len(b) == 0 {
		os.Exit(0)
	}

	if err != nil {
		log.Fatal(err)
	}

	o := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(o, b)

	os.Stdout.Write(o)
}
