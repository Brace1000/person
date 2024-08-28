package main

import (
	"fmt"
	"os"
)

func LoadPuzzle(puzzle string) int {
	count := 0
	for _, char := range puzzle {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
	}
	return count
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error")
		return
	}
	puzzle := os.Args[1]
	byte, err := os.ReadFile(puzzle)
	if err != nil {
		fmt.Println(err)
	}
	str := string(byte)
	Destination := LoadPuzzle((str))
	fmt.Println("the destination is :", Destination)
}
