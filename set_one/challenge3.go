/*Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

Achievement Unlocked
You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
*/
package set_one

import (
	"encoding/hex"
	"io/ioutil"
)

const Encoded string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func HexToBytes(hex_string string) []byte {
	encoded, err := hex.DecodeString(hex_string)
	check(err)
	return encoded
}

func SingleByteXOR(encoded []byte, b byte) []byte {
	decoded := make([]byte, len(encoded))
	for j := 0; j < len(encoded); j++ {
		decoded[j] = encoded[j] ^ b
	}
	return decoded
}

func ByteFrequencies(data []byte) map[byte]float64 {
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

func EnglishScorer() map[byte]float64 {
	data, err := ioutil.ReadFile("data/idleness.txt")
	check(err)
	return ByteFrequencies(data)
}

func CosineSimilarity(map_one, map_two map[byte]float64) float64 {
	tot := float64(0)
	for key, value := range map_one {
		tot += value * map_two[key]
	}
	return tot
}

func BreakXOR(base_freq map[byte]float64, encoded []byte) (string, float64) {
	best_similarity := float64(0)
	message := ""
	for b := 0; b < 128; b++ {
		xord := SingleByteXOR(encoded, byte(b))
		similarity := CosineSimilarity(ByteFrequencies(xord), base_freq)
		if similarity > best_similarity {
			best_similarity = similarity
			message = string(xord)
		}
	}
	return message, best_similarity
}

func SolveThree() string {
	doc_freq := EnglishScorer()
	message, _ := BreakXOR(doc_freq, HexToBytes(Encoded))
	return message
}
