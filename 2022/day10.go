package main

import (
	"os"
	"strconv"
	"strings"
)

func sumRange(start int, end int, arr []int) int {
	sum := 1
	for i := start; i < end && i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func parseCycles(instructions []string) []int {
	vals := []int{}
	for _, instr := range instructions {
		vals = append(vals, 0)
		if instr != "noop" {
			val, _ := strconv.Atoi(strings.Split(instr, " ")[1])
			vals = append(vals, val)
		}
	}
	return vals
}

func puzzle1(input string) {
	instructions := strings.Split(input, "\n")
	vals := parseCycles(instructions)
	sum := 0
	for _, num := range []int{20, 60, 100, 140, 180, 220} {
		sum += sumRange(0, num-1, vals) * num
	}
	println(sum)
}

func puzzle2(input string) {
	instructions := strings.Split(input, "\n")
	register := 1
	width := 40
	cycles := parseCycles(instructions)
	for cycle_idx, increment := range cycles {
		pixel := cycle_idx % width
		if pixel == register-1 || pixel == register || pixel == register+1 {
			print("#")
		} else {
			print(".")
		}
		register += increment
		if pixel == width-1 {
			println()
		}
	}
}

func main() {
	raw_input, _ := os.ReadFile("./day10_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
