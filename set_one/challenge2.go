package set_one

import (
	"encoding/hex"
)

func FixedXOR(hex_one, hex_two string) string {
	byte_one, _ := hex.DecodeString(hex_one)
	byte_two, _ := hex.DecodeString(hex_two)
	array_len := len(byte_one)
	xor_bytes := make([]byte, array_len)
	for j := 0; j < array_len; j++ {
		xor_bytes[j] = byte_one[j] ^ byte_two[j]
	}
	return hex.EncodeToString(xor_bytes)
}
