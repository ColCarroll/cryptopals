/*Detect AES in ECB mode
In this file are a bunch of hex-encoded ciphertexts.

One of them has been encrypted with ECB.

Detect it.

Remember that the problem with ECB is that it is stateless and deterministic;
the same 16 byte plaintext block will always produce the same 16 byte ciphertext.
*/
package set_one

import (
	"encoding/hex"
)

func DetectECB(filename string) string {
	var ciphertext, line []byte
	var byte_count, min_bytes int
	lines := ReadLines(filename)
	for _, hex_line := range lines {
		line = HexToBytes(hex_line)
		byte_counts := make(map[byte]bool)
		for _, b := range line {
			byte_counts[b] = true
		}
		byte_count = len(byte_counts)
		if byte_count == 0 {
			continue
		}
		if byte_count < min_bytes || min_bytes == 0 {
			min_bytes = byte_count
			ciphertext = line
		}
	}
	return hex.EncodeToString(ciphertext)
}
