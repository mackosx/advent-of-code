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

func puzzle1(input string) {
	instructions := strings.Split(input, "\n")
	vals := []int{}
	for _, instr := range instructions {
		vals = append(vals, 0)
		if instr != "noop" {
			val, _ := strconv.Atoi(strings.Split(instr, " ")[1])
			vals = append(vals, val)
		}
	}
	sum := 0
	for _, num := range []int{20, 60, 100, 140, 180, 220} {
		sum += sumRange(0, num-1, vals) * num
	}
	println(sum)
}

func puzzle2(input string) {

}

func main() {
	raw_input, _ := os.ReadFile("./day10_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
