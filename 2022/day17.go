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

func moveRock(o *OccupiedMap, old *Shape, new *Shape) {
	if old != nil {
		for _, point := range old.points {
			(*o)[point] = false
		}
	}
	for _, point := range new.points {
		(*o)[point] = true
	}
	old.points = new.points
}

func getNextShape(current_rock *Shape, o *OccupiedMap) *Shape {
	next_index := 0
	if current_rock != nil {
		next_index = current_rock.id + 1
	}
	max_y := 0
	if current_rock != nil {
		max_y = maxY(current_rock)
	}
	fmt.Printf("Current max Y %d\n", max_y)
	next_shape := shapes[next_index%4]
	points := []Point{}
	for _, point := range next_shape.points {
		points = append(points, Point{point.x, point.y + max_y + 3})
		println(point.y)
	}
	moveRock(o, &next_shape, &Shape{points: points})
	fmt.Printf("Next shape %v\n", next_shape)
	return &next_shape
}

func simulateFall(occupied *OccupiedMap, rock_shape *Shape) error {
	new_shape := Shape{[]Point{}, rock_shape.id}
	// Make a new shape 1 lower.
	for _, point := range rock_shape.points {
		new_height := point.y - 1
		if new_height < 0 {
			return errors.New("Can't complete fall, we're at the bottom.")
		}
		new_shape.points = append(
			new_shape.points,
			Point{point.x, new_height})
	}
	for _, point := range new_shape.points {
		is_occupied, _ := (*occupied)[point]
		if is_occupied {
			return errors.New("Can't complete fall, we hit another rock.")
		}
	}
	moveRock(occupied, rock_shape, &new_shape)
	return nil
}

func simulateJet(direction rune, o *OccupiedMap, current_rock *Shape) {
	new_shape := Shape{[]Point{}, current_rock.id}
	sign := -1
	if string(direction) == ">" {
		sign = 1
	}
	for _, point := range current_rock.points {
		new_x := point.x + sign
		if new_x > 6 {
			return
		}
		new_shape.points = append(
			new_shape.points,
			Point{point.x + 1, point.y})
	}
	for _, point := range new_shape.points {
		is_occupied, _ := (*o)[point]
		if is_occupied {
			return
		}
	}
	for _, point := range current_rock.points {
		(*o)[point] = false
	}
	for _, point := range new_shape.points {
		(*o)[point] = true
	}
	current_rock.points = new_shape.points
}

func puzzle1(input string) {
	rock_fall_count := 0
	max_falls := 2022
	occupied := make(OccupiedMap)
	i := 0
	current_rock := getNextShape(nil, &occupied)
	fmt.Printf("Starting rock: %v\n", current_rock)
	for rock_fall_count < max_falls {
		fmt.Printf("Simulating fall on id: %d\n", current_rock.id)
		err := simulateFall(&occupied, current_rock)
		if err != nil {
			// this means we can't move anymore
			println(err.Error())
			rock_fall_count++
			current_rock = getNextShape(current_rock, &occupied)
			fmt.Printf("Next shape: %v\n", current_rock)
		}
		println("Fall completed")
		movement := input[i%(len(input)-1)]
		fmt.Printf("Simulating jet (%s) on id: %d\n", string(movement), current_rock.id)
		simulateJet(rune(movement), &occupied, current_rock)
		i++
	}
	fmt.Printf("%d %+v %d\n", maxY(current_rock), current_rock, rock_fall_count)

}

func main() {
	raw_input, _ := os.ReadFile("./day17_input.txt")
	input := string(raw_input)
	puzzle1(input)
	// puzzle2(input)
}
