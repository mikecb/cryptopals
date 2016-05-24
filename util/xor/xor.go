package xor

import (
	"bytes"
	"encoding/hex"
	"math"
)

//XorBytes xors some bytes and returns some bytes, usally around half as many as went in.
func XorBytes(plaintext, key []byte) ([]byte, error) {
	var out bytes.Buffer
	for i := range plaintext {
		keyindex := int(math.Mod(float64(i), float64(len(key))))
		b := int(plaintext[i]) ^ int(key[keyindex])
		out.WriteByte(byte(b))
	}

	return out.Bytes(), nil
}

//XorHex hexes bytes. I mean it bites hexes. It does something.
func XorHex(b1, b2 string) (string, error) {
	raw1, err := hex.DecodeString(b1)
	if err != nil {
		return "", err
	}
	raw2, _ := hex.DecodeString(b2)
	if err != nil {
		return "", err
	}

	out, err := XorBytes(raw1, raw2)

	return hex.EncodeToString(out), err
}
