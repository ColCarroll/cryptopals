package main

import (
	"cryptopals/set_one"
	"fmt"
)

func main() {
	fmt.Println(set_one.BreakRepeatingKeyXor(set_one.B64File("set_one/data/6.txt")))
}
