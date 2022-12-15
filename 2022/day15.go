package main

import (
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

func puzzle2(input string) {
}

func main() {
	raw_input, _ := os.ReadFile("./day15_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
