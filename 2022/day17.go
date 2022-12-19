package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type Point struct {
	x int
	y int64
}

type Shape struct {
	points *[]Point
	id     int
}

type OccupiedMap = map[Point]bool

var shapes [5]Shape = [5]Shape{
	{&[]Point{{2, 0}, {3, 0}, {4, 0}, {5, 0}}, 0},         // —
	{&[]Point{{2, 1}, {3, 1}, {4, 1}, {3, 2}, {3, 0}}, 1}, // +
	{&[]Point{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}}, 2}, //backwards L
	{&[]Point{{2, 0}, {2, 1}, {2, 2}, {2, 3}}, 3},         // |
	{&[]Point{{2, 0}, {3, 0}, {2, 1}, {3, 1}}, 4},         // []
}

func maxY(rock *Shape) int64 {
	var maxY int64 = 0
	for _, point := range *rock.points {
		if point.y > maxY {
			maxY = point.y
		}
	}
	return maxY
}

func getNextShape(current_rock *Shape, o *OccupiedMap, height int64) *Shape {
	next_index := 0
	if current_rock != nil {
		next_index = current_rock.id + 1
	} else {
		height--
	}
	next_shape := shapes[next_index%5]
	points := make([]Point, len(*next_shape.points))
	for i := 0; i < len(*next_shape.points); i++ {
		point := &(*next_shape.points)[i]
		points[i] = Point{point.x, point.y + height + 3}
	}
	next_shape.points = &points
	return &next_shape
}

func simulateFall(occupied *OccupiedMap, rock_shape *Shape) error {
	new_points := make([]Point, len(*rock_shape.points))
	// Make a new shape 1 lower.
	for i := 0; i < len(*rock_shape.points); i++ {
		point := &(*rock_shape.points)[i]
		new_height := point.y - 1
		if new_height < 0 {
			return errors.New("Can't complete fall, we're at the bottom.")
		}
		new_points[i] = Point{point.x, new_height}
	}
	for i := 0; i < len(new_points); i++ {
		point := &(new_points)[i]
		is_occupied, _ := (*occupied)[*point]
		if is_occupied {
			return errors.New(fmt.Sprintf("Can't complete fall, we hit another rock at %v.", point))
		}
	}
	rock_shape.points = &new_points
	return nil
}

func simulateJet(direction rune, o *OccupiedMap, current_rock *Shape) {
	new_points := make([]Point, len(*current_rock.points))
	sign := -1
	if string(direction) == ">" {
		sign = 1
	}
	for i := 0; i < len(*current_rock.points); i++ {
		point := &(*current_rock.points)[i]
		new_x := point.x + sign
		if new_x > 6 || new_x < 0 {
			return
		}
		new_points[i] = Point{new_x, point.y}
	}
	for _, point := range new_points {
		is_occupied, _ := (*o)[point]
		if is_occupied {
			return
		}
	}
	current_rock.points = &new_points
}

func printOccupied(o *OccupiedMap, current_rock *Shape, height int64) {
	for h := height; h >= 0; h-- {
		for x := 0; x <= 6; x++ {
			current := Point{x, h}
			falling_pt := false
			for _, pt := range *current_rock.points {
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
	var rock_fall_count int64 = 0
	// var max_falls int64 = 1000000000000
	var max_falls int64 = 100
	occupied := make(OccupiedMap)
	i := 0
	var height int64 = 1
	current_rock := getNextShape(nil, &occupied, height)
	for rock_fall_count < max_falls {
		movement := input[i%(len(input))]
		simulateJet(rune(movement), &occupied, current_rock)
		err := simulateFall(&occupied, current_rock)
		if err != nil {
			rock_fall_count++
			for _, point := range *current_rock.points {
				occupied[point] = true
			}
			if top_of_rock := maxY(current_rock); top_of_rock+1 > height {
				height = top_of_rock + 1
			}
			current_rock = getNextShape(current_rock, &occupied, height)
		}
		if i%1000000 == 0 {
			fmt.Printf("%d\n", height)
		}
		i++
	}
	// Idea: we only need to store everything up from where the first #######
	// full line is. This will save memory. Is there any way to save CPU?
	printOccupied(&occupied, current_rock, height+5)
	fmt.Printf("%d\n", height)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	raw_input, _ := os.ReadFile("./day17_input.txt")
	input := string(raw_input)
	puzzle1(input)
	// puzzle2(input)
}
