// hex2bin - convert hexadecimal string to binary output

package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)

	if len(b) < 2 {
		os.Exit(0)
	}

	if (len(b)%2 != 0) || (err != nil) {
		log.Fatal(err)
	}

	o := make([]byte, hex.DecodedLen(len(b)))
	_, err = hex.Decode(o, b)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(o)
}
