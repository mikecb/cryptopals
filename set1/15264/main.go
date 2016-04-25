package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	input := ("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	rawbytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Printf("Error decoding hex: %s", err)
	}
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(rawbytes)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}
