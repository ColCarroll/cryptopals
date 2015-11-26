package set_one

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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
		chunks = append(chunks, data[j:j+chunksize])
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

type KeySize struct {
	size     int
	distance float64
}

func (k KeySize) String() string { return fmt.Sprintf("%d: %.2f", k.size, k.distance) }

type ByDistance []KeySize

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }

func GetKeySize(filename string, n int) []KeySize {
	data, err := ioutil.ReadFile(filename)
	check(err)
	distances := make([]KeySize, 0)
	for size := 2; size < 40; size++ {
		distances = append(distances,
			KeySize{size, data_avg_distance(chunk_data(data, size)) / float64(size)})
	}
	sort.Sort(ByDistance(distances))

	return distances[:n]
}
