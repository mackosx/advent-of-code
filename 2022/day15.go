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

func abs(num int) int {
	if num < 0 {
		return -1 * num
	} else {
		return num
	}
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func manhattanDist(p1 Position, p2 Position) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func parseLine(input string) (signal Position, beacon Position) {
	input = strings.ReplaceAll(input, "=", ",")
	input = strings.ReplaceAll(input, ":", ",")
	s_and_b_data := strings.Split(input, ",")
	signal.x, _ = strconv.Atoi(string(s_and_b_data[1]))
	signal.y, _ = strconv.Atoi(string(s_and_b_data[3]))
	beacon.x, _ = strconv.Atoi(string(s_and_b_data[5]))
	beacon.y, _ = strconv.Atoi(string(s_and_b_data[7]))
	return
}

func isCovered(sensors map[Position]int, pos Position) bool {
	for sensor, dist := range sensors {
		if manhattanDist(sensor, pos) <= dist {
			return true
		}
	}
	return false
}

func getCoveredPositions(input string, row int) int {
	closest_beacon := make(map[Position]Position)
	lines := strings.Split(input, "\n")
	right_edge := math.MinInt
	left_edge := math.MaxInt
	for _, line := range lines {
		sensor, beacon := parseLine(line)
		dist := manhattanDist(sensor, beacon)
		left_edge = min(left_edge, sensor.x-dist)
		right_edge = max(right_edge, sensor.x+dist)
		closest_beacon[sensor] = beacon
	}

	occupied := 0
	for x := left_edge; x <= right_edge; x++ {
		var position *Position
		for sensor, beacon := range closest_beacon {
			dist := manhattanDist(sensor, beacon)
			position = &Position{x, row}
			// Check if we are looking at a Beacon
			if position.x == beacon.x && position.y == beacon.y {
				break
			}
			// Check if we are in range of the sensor
			if manhattanDist(*position, sensor) <= dist {
				occupied++
				break
			}
		}
	}
	return occupied
}

func puzzle1(input string) {
	y := 2000000
	occupied := getCoveredPositions(input, y)
	println(occupied)
}

func getPositionsAtRange(sensor Position, distance int) []*Position {
	max_range := 4000000
	positions := []*Position{}
	mx := sensor.x + distance
	x := max(sensor.x-distance, 0)
	dy := 0
	y := sensor.y + dy
	for x <= mx && x <= max_range && dy <= distance && y <= max_range && y >= 0 {
		positions = append(positions, &Position{x, y})
		if dy != -dy && -y <= max_range && -y >= 0 {
			positions = append(positions, &Position{x, -y})
		}
		x++
		dy++
		y = sensor.y + dy
	}
	return positions
}

func puzzle2(input string) {
	lines := strings.Split(input, "\n")
	valid_points_outside_count := make(map[Position]int)
	var highest_count_outside *Position
	max_count := 0
	sensor_dist := make(map[Position]int)
	for _, line := range lines {
		sensor, beacon := parseLine(line)
		dist := manhattanDist(sensor, beacon)
		sensor_dist[sensor] = dist
	}
	for sensor, dist := range sensor_dist {

		points := getPositionsAtRange(sensor, dist+1)
		for _, pos := range points {
			if !isCovered(sensor_dist, *pos) {
				count, _ := valid_points_outside_count[*pos]
				count++
				if count > max_count {
					max_count = count
					highest_count_outside = pos
				}
				valid_points_outside_count[*pos] = count
			}
		}
	}
	fmt.Printf("%+v Count: %d\n, Len: %d\n", *highest_count_outside, max_count, len(valid_points_outside_count))
	println((highest_count_outside.x * 4000000) + highest_count_outside.y)
}

func main() {
	raw_input, _ := os.ReadFile("./day15_input.txt")
	input := string(raw_input)
	// puzzle1(input)
	puzzle2(input)
}
