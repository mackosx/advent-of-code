package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coverage struct {
	left_x   int
	right_x  int
	top_y    int
	bottom_y int
}

type Position struct {
	x int
	y int
}

func getCoverage(signal Position, closest_beacon Position) Coverage {
	dist := int(math.Abs(float64(signal.x-closest_beacon.x) + math.Abs(float64(signal.y-closest_beacon.y))))
	return Coverage{
		left_x:   signal.x - dist,
		right_x:  signal.x + dist,
		top_y:    signal.y + dist,
		bottom_y: signal.y - dist}
}

func isInside(position Position, coverage Coverage) bool {
	// TODO: Refactor to support "manhattan" inside; this generates a square box
	return (position.x >= coverage.left_x && position.x <= coverage.right_x &&
		position.y <= coverage.top_y && position.y >= coverage.bottom_y)
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

func puzzle1(input string) {
	// build up list of areas from signal/beacon input

	// check if an individual position intersects with a list of "areas"
	signal_to_coverage := make(map[Position]Coverage)
	lines := strings.Split(input, "\n")
	right_edge := math.MinInt
	left_edge := math.MaxInt
	for _, line := range lines {

		signal, beacon := parseLine(line)
		coverage := getCoverage(signal, beacon)

		left_edge = int(math.Min(float64(left_edge), float64(coverage.left_x)))
		right_edge = int(math.Max(float64(right_edge), float64(coverage.right_x)))

		signal_to_coverage[signal] = coverage
		fmt.Printf("%+v,  %d -> %d\n", signal, left_edge, right_edge)

	}
	y := 10
	occupied := 0
	for x := left_edge; x <= right_edge; x++ {
		inside := false
		for _, cov := range signal_to_coverage {
			if inside = isInside(Position{x, y}, cov); inside {
				occupied++
				print("#")
				break
			}
		}
		if !inside {
			print(".")
		}
	}
	println()
	println(left_edge, " -> ", right_edge)
	println(occupied)

	// need to handle beacons as well

}

func puzzle2(input string) {
}

func main() {
	raw_input, _ := os.ReadFile("./day15_test.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
