package main

import (
	"fmt"
	"os"
	"strings"
)

var opponentMap = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

var beats = map[string]string{
	"R": "S",
	"P": "R",
	"S": "P",
}

var loss = map[string]string{
	"R": "P",
	"P": "S",
	"S": "R",
}

var myMap = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var points = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

func outcomePoints(opponent_shape string, my_shape string) int {
	if my_shape == opponent_shape {
		return 3
	}

	win_shape, _ := beats[my_shape]
	if win_shape == opponent_shape {
		return 6
	} else {
		return 0
	}
}

func puzzle1(input string) {
	total_score := 0
	rounds := strings.Split(input, "\n")
	for _, round := range rounds {
		letters := strings.Split(round, " ")
		opponent_shape, _ := opponentMap[letters[0]]
		my_shape, _ := myMap[letters[1]]
		shape_points, _ := points[my_shape]
		total_score += shape_points + outcomePoints(opponent_shape, my_shape)
	}
	fmt.Println(total_score)
}

func shapeForOutcome(required_outcome string, opponent_shape string) (shape string) {
	if required_outcome == "Y" {
		shape = opponent_shape
	} else if required_outcome == "Z" {
		shape, _ = loss[opponent_shape]
	} else {
		shape, _ = beats[opponent_shape]
	}
	return
}

func puzzle2(input string) {
	total_score := 0
	rounds := strings.Split(input, "\n")
	for _, round_data := range rounds {
		letters := strings.Split(round_data, " ")
		opponent_shape, _ := opponentMap[letters[0]]
		my_shape := shapeForOutcome(letters[1], opponent_shape)
		shape_points, _ := points[my_shape]
		total_score += shape_points + outcomePoints(opponent_shape, my_shape)
	}
	fmt.Println(total_score)
}

func main() {
	raw_input, _ := os.ReadFile("./day2_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
