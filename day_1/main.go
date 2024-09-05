package main

import (
	"fmt"
	"os"

	part2 "day_1/part_2"
)

func main() {
	file, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("error reading file", err)
	}
	filestr := string(file)
	count := 0
	for _, v := range filestr {
		if string(v) == "(" {
			count++
		} else {
			count--
		}
	}
	fmt.Println(count)
	fmt.Println(part2.Firstchar(filestr))
}
