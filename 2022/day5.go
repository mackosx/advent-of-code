package main

import (
	"os"
	"strconv"
	"strings"
)

func puzzle1(input string) {
	parts := strings.Split(input, "\n\n")
	state := strings.Split(parts[0], "\n")
	// Populate initial state
	stacks := make([][]string, 9)
	for height := 0; height < len(state)-1; height++ {
		line := state[height]
		for i := 1; i < len(line); i += 4 {
			if crate := string(line[i]); crate != " " {
				index := (i / 4)
				stacks[index] = append(stacks[index], crate)
			}
		}
	}
	instructions := strings.Split(parts[1], "\n")
	for _, line := range instructions {
		parts = strings.Split(line, " ")
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		for i := 0; i < count; i++ {
			from_stack := stacks[from-1]
			elem, from_stack := from_stack[0], from_stack[1:]
			stacks[from-1] = from_stack
			// Prepend item
			stacks[to-1] = append([]string{elem}, stacks[to-1]...)
		}
	}
	results := ""
	for _, stack := range stacks {
		results += stack[0]
	}
	println(results)
}

func main() {
	raw_input, _ := os.ReadFile("./day5_input.txt")
	input := string(raw_input)
	puzzle1(input)
}
