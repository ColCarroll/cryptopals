package set_one

import (
	"bufio"
	"log"
	"os"
)

func SolveFour() string {
	base_freq := DocumentFrequencies("set_one/data/idleness.txt")
	best_score := float64(0)
	best_message := ""
	file, err := os.Open("set_one/data/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		message, score := BreakXOR(base_freq, line)
		if score > best_score {
			best_score = score
			best_message = message
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return best_message
}
