package set_one

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func hex_to_bytes(hex_string string) []byte {
	data, err := hex.DecodeString(hex_string)
	check(err)
	return data
}

func hex_to_base64(hex_string string) string {
	return base64.StdEncoding.EncodeToString(hex_to_bytes(hex_string))
}

func bytes_to_hex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func get_file(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return data
}

func read_lines(filename string) []string {
	return strings.Split(string(get_file(filename)), "\n")
}

func b64_file(filename string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(string(get_file(filename)))
	check(err)
	return bytes
}

func cosine_similarity(map_one, map_two map[byte]float64) float64 {
	var tot float64 = 0
	for key, value := range map_one {
		tot += value * map_two[key]
	}
	return tot
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func chunk_data(data []byte, chunksize int) [][]byte {
	chunks := make([][]byte, 0)
	for j := 0; j < len(data); j += chunksize {
		chunks = append(chunks, data[j:min(len(data), j+chunksize)])
	}
	return chunks
}

func transpose(data [][]byte) [][]byte {
	trans := make([][]byte, len(data[0]))
	for j := 0; j < len(trans); j++ {
		trans[j] = make([]byte, len(data))
	}

	for row, row_data := range data {
		for col, entry := range row_data {
			trans[col][row] = entry
		}
	}
	return trans
}

func get_key_size(bytes []byte) int {
	best_distance, best_key := float64(8), 0
	for size := 2; size < 40; size++ {
		distance := data_avg_distance(chunk_data(bytes, size)) / float64(size)
		if distance < best_distance {
			best_distance = distance
			best_key = size
		}
	}
	return best_key
}

func byte_frequencies(data []byte) map[byte]float64 {
	m := make(map[byte]int)
	tot := 0
	for j := 0; j < len(data); j++ {
		m[data[j]] += 1
		tot += 1
	}
	percents := make(map[byte]float64)
	for key, value := range m {
		percents[key] = float64(value) / float64(tot)
	}
	return percents
}

func english_scorer() map[byte]float64 {
	data, err := ioutil.ReadFile("data/idleness.txt")
	check(err)
	return byte_frequencies(data)
}

func data_avg_distance(data [][]byte) float64 {
	tot, distance := 0, 0
	max_chunks := 4
	for j := 0; j < len(data) && j < max_chunks; j++ {
		for k := j + 1; k < len(data) && k < max_chunks; k++ {
			tot += 1
			distance += byte_distance(data[j], data[k])
		}
	}
	return float64(distance) / float64(tot)
}

func bit_sum(a byte) int {
	tot := 0
	for j := uint(0); j < 8; j++ {
		tot += int((a >> j) & 1)
	}
	return tot
}

func byte_distance(a, b []byte) int {
	tot := 0
	for idx, char := range a {
		tot += bit_sum(char ^ b[idx])
	}
	return tot
}

func FixedXOR(hex_one, hex_two string) []byte {
	byte_one := hex_to_bytes(hex_one)
	byte_two := hex_to_bytes(hex_two)

	if len(byte_one) != len(byte_two) {
		panic("Can only XOR slices of the same length")
	}

	array_len := len(byte_one)
	xor_bytes := make([]byte, array_len)
	for j := 0; j < array_len; j++ {
		xor_bytes[j] = byte_one[j] ^ byte_two[j]
	}
	return xor_bytes
}

func SingleByteXOR(encoded []byte, b byte) []byte {
	decoded := make([]byte, len(encoded))
	for j := 0; j < len(encoded); j++ {
		decoded[j] = encoded[j] ^ b
	}
	return decoded
}

func BreakSingleByteXOR(encoded []byte, base_freq map[byte]float64) ([]byte, float64) {
	best_similarity := float64(0)
	var message []byte
	for b := 0; b < 128; b++ {
		xord := SingleByteXOR(encoded, byte(b))
		similarity := cosine_similarity(byte_frequencies(xord), base_freq)
		if similarity > best_similarity {
			best_similarity = similarity
			message = xord
		}
	}
	return message, best_similarity
}

func RepeatingKeyXOR(text, key string) []byte {
	byte_text := []byte(text)
	byte_key := []byte(key)
	encrypted := make([]byte, len(byte_text))
	for j := 0; j < len(byte_text); j++ {
		encrypted[j] = byte_text[j] ^ byte_key[j%len(byte_key)]
	}
	return encrypted
}

func HammingDistance(first, second string) int {
	return byte_distance([]byte(first), []byte(second))
}

func BreakRepeatingKeyXOR(encoded []byte) []byte {
	blocks := transpose(chunk_data(encoded, get_key_size(encoded)))
	decrypted := make([][]byte, len(blocks))
	for idx, block := range blocks {
		decoded, _ := BreakSingleByteXOR(block, english_scorer())
		decrypted[idx] = []byte(decoded)
	}
	var buffer bytes.Buffer
	for _, x := range transpose(decrypted) {
		buffer.Write(x)
	}
	return buffer.Bytes()
}

func DecryptECB(ciphertext, key []byte) []byte {
	blocksize := len(key)
	plaintext := make([]byte, len(ciphertext))
	crypt, err := aes.NewCipher(key)
	check(err)
	for j := 0; j < len(ciphertext); j += blocksize {
		crypt.Decrypt(plaintext[j:j+blocksize], ciphertext[j:j+blocksize])
	}
	return plaintext
}
func FindXORLine(filename string) []byte {
	best_score := float64(0)
	scorer := english_scorer()
	var best_message []byte
	for _, hex_line := range read_lines(filename) {
		message, score := BreakSingleByteXOR(hex_to_bytes(hex_line), scorer)
		if score > best_score {
			best_score = score
			best_message = message
		}
	}
	return best_message
}

func DetectECB(filename string) []byte {
	var ciphertext, line []byte
	var byte_count, min_bytes int
	for _, hex_line := range read_lines(filename) {
		line = hex_to_bytes(hex_line)
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
	return ciphertext
}
