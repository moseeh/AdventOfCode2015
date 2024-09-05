package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}
	filestr := string(file)
	filearr := strings.Split(filestr, "\n")
	tsa := 0
	tslack := 0
	tribbow := 0
	for _, line := range filearr {
		linearr := strings.Split(line, "x")
		l, err := strconv.Atoi(linearr[0])
		w, err1 := strconv.Atoi(linearr[1])
		h, err2 := strconv.Atoi(linearr[2])
		if err1 == nil && err == nil && err2 == nil {
			sa := (2 * l * h) + (2 * l * w) + (2 * w * h)
			tsa += sa
			tslack += slack(l, w, h)
			rib, bow := ribbon(l, w, h)
			tribbow += rib + bow

		}
	}
	fmt.Println(tsa + tslack)
	fmt.Println(tribbow)
}

func slack(l, w, h int) int {
	a := l * w
	b := l * h
	c := w * h
	if a <= b && a <= c {
		return a
	}
	if b <= c && b <= a {
		return b
	}
	return c
}

func ribbon(l, w, h int) (int, int) {
	a := l + w
	b := l + h
	c := w + h

	bow := l * w * h
	rib := 0
	if a <= b && a <= c {
		return 2 * a, bow
	}
	if b <= a && b <= c {
		return 2 * b, bow
	}
	rib = 2 * c
	return rib, bow
}
