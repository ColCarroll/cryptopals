package set_one

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hex_string string) string {
	if data, err := hex.DecodeString(hex_string); err != nil {
		panic(err)
	} else {
		return base64.StdEncoding.EncodeToString(data)
	}
}
