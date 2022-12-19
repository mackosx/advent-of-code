package main

import (
	"errors"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

type Shape struct {
	points []Point
	id     int
}

type OccupiedMap = map[Point]bool

var shapes [5]Shape = [5]Shape{
	{[]Point{{2, 0}, {3, 0}, {4, 0}, {5, 0}}, 0},         // —
	{[]Point{{2, 1}, {3, 1}, {4, 1}, {3, 2}, {3, 0}}, 1}, // +
	{[]Point{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}}, 2}, //backwards L
	{[]Point{{2, 0}, {2, 1}, {2, 2}, {2, 3}}, 3},         // |
	{[]Point{{2, 0}, {3, 0}, {2, 1}, {3, 1}}, 4},         // []
}

func maxY(rock *Shape) int {
	maxY := 0
	for _, point := range rock.points {
		if point.y > maxY {
			maxY = point.y
		}
	}
	return maxY
}

func getNextShape(current_rock *Shape, o *OccupiedMap, height int) *Shape {
	next_index := 0
	if current_rock != nil {
		next_index = current_rock.id + 1
	} else {
		height--
	}
	next_shape := shapes[next_index%5]
	points := []Point{}
	for _, point := range next_shape.points {
		points = append(points, Point{point.x, point.y + height + 3})
	}
	next_shape.points = points
	return &next_shape
}

func simulateFall(occupied *OccupiedMap, rock_shape *Shape) error {
	new_points := []Point{}
	// Make a new shape 1 lower.
	for _, point := range rock_shape.points {
		new_height := point.y - 1
		if new_height < 0 {
			return errors.New("Can't complete fall, we're at the bottom.")
		}
		new_points = append(
			new_points,
			Point{point.x, new_height})
	}
	for _, point := range new_points {
		is_occupied, _ := (*occupied)[point]
		if is_occupied {
			return errors.New(fmt.Sprintf("Can't complete fall, we hit another rock at %v.", point))
		}
	}
	rock_shape.points = new_points
	return nil
}

func simulateJet(direction rune, o *OccupiedMap, current_rock *Shape) {
	new_points := []Point{}
	sign := -1
	if string(direction) == ">" {
		sign = 1
	}
	for _, point := range current_rock.points {
		new_x := point.x + sign
		if new_x > 6 || new_x < 0 {
			return
		}
		new_points = append(
			new_points,
			Point{new_x, point.y})
	}
	for _, point := range new_points {
		is_occupied, _ := (*o)[point]
		if is_occupied {
			return
		}
	}
	current_rock.points = new_points
}

func printOccupied(o *OccupiedMap, current_rock *Shape, height int) {
	for h := height; h >= 0; h-- {
		for x := 0; x <= 6; x++ {
			current := Point{x, h}
			falling_pt := false
			for _, pt := range current_rock.points {
				if pt == current {
					falling_pt = true
					print("@")
					break
				}
			}
			is_occupied, _ := (*o)[current]
			if !falling_pt {
				if is_occupied {
					print("#")
				} else {
					print(".")
				}
			}
		}
		println()
	}
	println()
}

func puzzle1(input string) {
	rock_fall_count := 0
	max_falls := 2022
	occupied := make(OccupiedMap)
	i := 0
	height := 1
	current_rock := getNextShape(nil, &occupied, height)
	for rock_fall_count < max_falls {
		movement := input[i%(len(input))]
		simulateJet(rune(movement), &occupied, current_rock)
		err := simulateFall(&occupied, current_rock)
		if err != nil {
			rock_fall_count++
			for _, point := range current_rock.points {
				occupied[point] = true
			}
			if top_of_rock := maxY(current_rock); top_of_rock+1 > height {
				height = top_of_rock + 1
			}
			current_rock = getNextShape(current_rock, &occupied, height)
		}
		i++
	}
	fmt.Printf("%d\n", height)
}

func main() {
	raw_input, _ := os.ReadFile("./day17_input.txt")
	input := string(raw_input)
	puzzle1(input)
	// puzzle2(input)
}
