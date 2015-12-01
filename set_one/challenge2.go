/*Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179
*/
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
