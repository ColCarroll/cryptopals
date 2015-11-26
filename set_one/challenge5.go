package set_one

import (
	"encoding/hex"
)

func RepeatingKeyXOR(text, key string) string {
	byte_text := []byte(text)
	byte_key := []byte(key)
	encrypted := make([]byte, len(byte_text))
	for j := 0; j < len(byte_text); j++ {
		encrypted[j] = byte_text[j] ^ byte_key[j%len(byte_key)]
	}
	return hex.EncodeToString(encrypted)
}
