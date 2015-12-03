/*Detect single-character XOR
One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)
*/
package set_one

import (
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	check(err)
	lines := strings.Split(string(content), "\n")
	return lines
}

func FindXORLine(filename string) string {
	base_freq := EnglishScorer()
	best_score := float64(0)
	best_message := ""
	for _, hex_line := range ReadLines(filename) {
		line := HexToBytes(hex_line)
		message, score := BreakXOR(base_freq, line)
		if score > best_score {
			best_score = score
			best_message = message
		}
	}
	return best_message
}
