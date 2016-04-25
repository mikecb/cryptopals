package hex2b64

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func hex2b64(s string) int {
	b64input, err := hex.DecodeString(s)
	if err != nil {
		fmt.Printf("Error decoding hex: %s", err)
	}
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	b, err := encoder.Write(b64input)
	if err != nil {
		fmt.Printf("Error encoding b64: %s", err)
	}
	encoder.Close()

	return b
}
