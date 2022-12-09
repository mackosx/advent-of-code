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

func tailNeedsToMove(head Knot, tail Knot) bool {
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
			if tailNeedsToMove(head, tail) {
				tail.position = head.previous
				positions_occupied[fmt.Sprintf("%d,%d", tail.position.x, tail.position.y)] = true
			}
		}
	}
	println(len(positions_occupied))
}

func puzzle2(input string) {

}

func main() {
	raw_input, _ := os.ReadFile("./day9_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
