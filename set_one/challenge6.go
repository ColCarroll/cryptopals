/* Break repeating-key XOR
It is officially on, now.
This challenge isn't conceptually hard, but it involves actual error-prone coding. The other challenges in this set are there to bring you up to speed. This one is there to qualify you. If you can do this one, you're probably just fine up to Set 6.

There's a file here. It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

Here's how:

Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the number of differing bits. The distance between:
this is a test
and
wokka wokka!!!
is 37. Make sure your code agrees before you proceed.
For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4 KEYSIZE blocks instead of 2 and average the distances.
Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
Solve each block as if it was single-character XOR. You already have code to do this.
For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.
This code is going to turn out to be surprisingly useful later on. Breaking repeating-key XOR ("Vigenere") statistically is obviously an academic exercise, a "Crypto 101" thing. But more people "know how" to break it than can actually break it, and a similar technique breaks something much more important.

No, that's not a mistake.
We get more tech support questions for this challenge than any of the other ones. We promise, there aren't any blatant errors in this text. In particular: the "wokka wokka!!!" edit distance really is 37.
*/
package set_one

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func bit_sum(a byte) int {
	tot := 0
	for j := uint(0); j < 8; j++ {
		tot += int((a >> j) & 1)
	}
	return tot
}

func ByteDistance(a, b []byte) int {
	tot := 0
	for idx, char := range a {
		tot += bit_sum(char ^ b[idx])
	}
	return tot
}

func HammingDistance(first, second string) int {
	return ByteDistance([]byte(first), []byte(second))
}

func data_avg_distance(data [][]byte) float64 {
	tot, distance := 0, 0
	max_chunks := 4
	for j := 0; j < len(data) && j < max_chunks; j++ {
		for k := j + 1; k < len(data) && k < max_chunks; k++ {
			tot += 1
			distance += ByteDistance(data[j], data[k])
		}
	}
	return float64(distance) / float64(tot)
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

func B64File(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check(err)
	bytes, err := base64.StdEncoding.DecodeString(string(data))
	check(err)
	return bytes
}

func GetKeySize(bytes []byte) int {
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

func BreakRepeatingKeyXor(encoded []byte) string {
	scorer := EnglishScorer()
	key_size := GetKeySize(encoded)
	blocks := transpose(chunk_data(encoded, key_size))
	decrypted := make([][]byte, len(blocks))
	for idx, block := range blocks {
		decoded, _ := BreakXOR(scorer, block)
		decrypted[idx] = []byte(decoded)
	}
	var buffer bytes.Buffer
	for _, x := range transpose(decrypted) {
		buffer.Write(x)
	}
	return buffer.String()
}
