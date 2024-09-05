package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}
	filestr := string(file)
	arr := [2]int{}
	cordinates := make(map[[2]int]int)
	x, y := 0, 0
	arr[0] = x
	arr[1] = y
	cordinates[arr]++

	for _, v := range filestr {
		if v == '^' {
			y++
		}
		if v == 'v' {
			y--
		}
		if v == '>' {
			x++
		}
		if v == '<' {
			x--
		}
		arr[0] = x
		arr[1] = y
		cordinates[arr]++
	}
	fmt.Println(len(cordinates))
	fmt.Println(robosanta(filestr))
}

func robosanta(str string) int {
	arr := [2]int{}
	cordinates := make(map[[2]int]int)
	x, y, x1, y1 := 0, 0, 0, 0
	arr[0] = x
	arr[1] = y
	cordinates[arr]++

	for i, v := range str {
		if i%2 == 0 {

			if v == '^' {
				y++
			}
			if v == 'v' {
				y--
			}
			if v == '>' {
				x++
			}
			if v == '<' {
				x--
			}
			arr[0] = x
			arr[1] = y
			cordinates[arr]++
		} else {
			if v == '^' {
				y1++
			}
			if v == 'v' {
				y1--
			}
			if v == '>' {
				x1++
			}
			if v == '<' {
				x1--
			}
			arr[0] = x1
			arr[1] = y1
			cordinates[arr]++
		}
	}
	return len(cordinates)
}
