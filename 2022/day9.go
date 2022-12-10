package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Knot struct {
	position Position
	previous Position
}

func knotNeedsToMove(head *Knot, tail *Knot) bool {
	for _, i := range [3]int{-1, 0, 1} {
		for _, j := range [3]int{-1, 0, 1} {
			x := tail.position.x + i
			y := tail.position.y + j
			if head.position.x == x && head.position.y == y {
				return false
			}
		}
	}
	return true
}

func puzzle1(input string) {
	instructions := strings.Split(input, "\n")
	positions_occupied := make(map[string]bool)
	head := Knot{position: Position{0, 0}, previous: Position{}}
	tail := Knot{position: Position{0, 0}, previous: Position{}}
	positions_occupied["0,0"] = true
	for _, instruction_text := range instructions {
		tokens := strings.Split(instruction_text, " ")
		direction := tokens[0]
		amount, _ := strconv.Atoi(tokens[1])
		for i := 0; i < amount; i++ {
			head.previous = head.position
			switch direction {
			case "R":
				head.position.x++
			case "L":
				head.position.x--
			case "U":
				head.position.y++
			case "D":
				head.position.y--
			}
			if knotNeedsToMove(&head, &tail) {
				tail.position = head.previous
				positions_occupied[fmt.Sprintf("%d,%d", tail.position.x, tail.position.y)] = true
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

func calculateCatchUp(head *Knot, tail *Knot) Position {
	x_increment := clamp(head.position.x-tail.position.x, -1, 1)
	y_increment := clamp(head.position.y-tail.position.y, -1, 1)
	return Position{tail.position.x + x_increment, tail.position.y + y_increment}
}

func puzzle2(input string) {
	instructions := strings.Split(input, "\n")
	positions_occupied := make(map[string]bool)
	knots := [10]*Knot{}
	for i := 0; i < len(knots); i++ {
		knots[i] = &Knot{}
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
				head.position.x++
			case "L":
				head.position.x--
			case "U":
				head.position.y++
			case "D":
				head.position.y--
			}
			for knot_index := 1; knot_index < len(knots); knot_index++ {
				current := knots[knot_index-1]
				next := knots[knot_index]
				if knotNeedsToMove(current, next) {
					next.position = calculateCatchUp(current, next)
					if knot_index == len(knots)-1 {
						positions_occupied[fmt.Sprintf("%d,%d", next.position.x, next.position.y)] = true
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
