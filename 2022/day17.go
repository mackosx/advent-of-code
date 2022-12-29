package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
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

func getHashKey(row string, input_index int, rock_id int) string {
	return fmt.Sprintf("%s %d %d", row, input_index, rock_id)
}

type CycleResult struct {
	count  int64
	height int64
}

func puzzle(input string) {
	var rock_fall_count int64 = 0
	const max_falls int64 = 1_000_000_000_000
	occupied := make(OccupiedMap)
	var height int64 = 1
	var height_diff int64
	cycle_rocks := make(map[string]CycleResult)
	current_rock := getNextShape(nil, &occupied, height)
	cycled := false
	for i := 0; rock_fall_count <= max_falls; i++ {
		input_index := i % (len(input))
		movement := input[input_index]
		simulateJet(rune(movement), &occupied, current_rock)
		err := simulateFall(&occupied, current_rock)
		if err != nil {
			rock_fall_count++
			for _, point := range *current_rock.points {
				occupied[point] = true
			}
			top_of_rock := maxY(current_rock)
			difference := top_of_rock + 1 - height
			if difference > 0 {
				height += difference
			}
			// Start at 10,000 to avoid initial settling
			if rock_fall_count > 10_000 && !cycled {
				row_str := strings.Builder{}
				for col := 0; col < 7; col++ {
					_, ok := occupied[Point{col, height - 1}]
					if ok {
						row_str.WriteString("#")
					} else {
						row_str.WriteString(".")
					}
				}
				key := getHashKey(row_str.String(), input_index, current_rock.id)
				cycle_result, ok := cycle_rocks[key]
				if ok {
					// Increment rock falls and height by cycle increments
					cycle_count := rock_fall_count - cycle_result.count
					cycle_height := height - cycle_result.height
					for rock_fall_count+cycle_count < max_falls {
						rock_fall_count += cycle_count
						height_diff += cycle_height
					}
					cycled = true
				} else {
					cycle_rocks[key] = CycleResult{rock_fall_count, height}
				}
			}
			current_rock = getNextShape(current_rock, &occupied, height)
		}
		if i%1000 == 0 {
			fmt.Printf("Height: %d Rock Count: %d\n", height, rock_fall_count)
		}
	}
	fmt.Printf("%d\n", height-1+height_diff)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	raw_input, _ := os.ReadFile("./day17_input.txt")
	input := string(raw_input)
	puzzle(input)
}
