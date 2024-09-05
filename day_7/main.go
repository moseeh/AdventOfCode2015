package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	instructions = make(map[string]string)
	signals      = make(map[string]uint16)
)

func main() {
	file, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	filestr := string(file)
	filearr := strings.Split(filestr, "\n")
	
	for _, v := range filearr {
		if v == "" {
			continue
		}
		parts := strings.Split(v, "->")
		wire := strings.TrimSpace(parts[1])
		instructions[wire] = strings.TrimSpace(parts[0])
	}

	fmt.Println("Signal on wire 'a':", getSignal("a"))
}

func getSignal(wire string) uint16 {
	if signal, ok := signals[wire]; ok {
		return signal
	}

	instruction := instructions[wire]
	parts := strings.Fields(instruction)

	var result uint16

	switch len(parts) {
	case 1:
		result = parseSignal(parts[0])
	case 2:
		operand := getSignal(parts[1])
		result = ^operand
	case 3:
		left := parseSignal(parts[0])
		right := parseSignal(parts[2])
		switch parts[1] {
		case "AND":
			result = left & right
		case "OR":
			result = left | right
		case "LSHIFT":
			result = left << right
		case "RSHIFT":
			result = left >> right
		}
	}

	signals[wire] = result
	return result
}

func parseSignal(s string) uint16 {
	if num, err := strconv.ParseUint(s, 10, 16); err == nil {
		return uint16(num)
	}
	return getSignal(s)
}