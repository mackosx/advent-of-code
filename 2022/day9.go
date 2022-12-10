package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func knotNeedsToMove(head *Position, tail *Position) bool {
	return math.Abs(float64(head.x-tail.x)) > 1 || math.Abs(float64(head.y-tail.y)) > 1
}

func puzzle1(input string) {
	instructions := strings.Split(input, "\n")
	positions_occupied := make(map[string]bool)
	head := Position{0, 0}
	tail := Position{0, 0}
	positions_occupied["0,0"] = true
	for _, instruction_text := range instructions {
		tokens := strings.Split(instruction_text, " ")
		direction := tokens[0]
		amount, _ := strconv.Atoi(tokens[1])
		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}
			if knotNeedsToMove(&head, &tail) {
				tail = calculateCatchUp(&head, &tail)
				positions_occupied[fmt.Sprintf("%d,%d", tail.x, tail.y)] = true
			}
		}
	}
	println(len(positions_occupied))
}

func clamp(num, low, high int) int {
	if num < low {
		return low
	}
	if num > high {
		return high
	}
	return num
}

func calculateCatchUp(head *Position, tail *Position) Position {
	return Position{tail.x + clamp(head.x-tail.x, -1, 1), tail.y + clamp(head.y-tail.y, -1, 1)}
}

func puzzle2(input string) {
	instructions := strings.Split(input, "\n")
	positions_occupied := make(map[string]bool)
	knots := [10]*Position{}
	for i := 0; i < len(knots); i++ {
		knots[i] = &Position{}
	}
	positions_occupied["0,0"] = true
	for _, instruction_text := range instructions {
		tokens := strings.Split(instruction_text, " ")
		direction := tokens[0]
		amount, _ := strconv.Atoi(tokens[1])
		head := knots[0]
		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}
			for knot_index := 1; knot_index < len(knots); knot_index++ {
				current := knots[knot_index-1]
				next := knots[knot_index]
				if knotNeedsToMove(current, next) {
					pos := calculateCatchUp(current, next)
					knots[knot_index] = &pos
					if knot_index == len(knots)-1 {
						positions_occupied[fmt.Sprintf("%d,%d", next.x, next.y)] = true
					}
				}

			}
		}
	}
	println(len(positions_occupied))
}

func main() {
	raw_input, _ := os.ReadFile("./day9_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
