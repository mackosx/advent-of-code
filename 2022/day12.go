package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getKey(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func get(m map[string]int, k string) int {
	val, _ := m[k]
	return val
}

func printCurrentPath(path []string, width int, height int, neighbours []string) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			found := false
			found_n := false
			for _, node := range path {
				coords := strings.Split(node, ",")
				node_i, _ := strconv.Atoi(coords[0])
				node_j, _ := strconv.Atoi(coords[1])
				if node_i == i && node_j == j {
					found = true
				}
			}
			for _, node := range neighbours {
				coords := strings.Split(node, ",")
				node_i, _ := strconv.Atoi(coords[0])
				node_j, _ := strconv.Atoi(coords[1])
				if node_i == i && node_j == j {
					found_n = true
				}
			}
			if found {
				print("x")
			} else if found_n {
				print("n")
			} else {
				print("•")
			}
		}
		println()
	}
	println()
}

func bfs(start_key string, end_key string, elevations map[string]int, width int, height int, goal_elevation int) []string {
	getNeighbours := func(node string) []string {
		coords := strings.Split(node, ",")
		i, _ := strconv.Atoi(coords[0])
		j, _ := strconv.Atoi(coords[1])
		neighbours := []string{}
		current_elevation := get(elevations, node)
		if new_i := i - 1; new_i >= 0 {
			if current_elevation-get(elevations, getKey(new_i, j)) <= 1 {
				neighbours = append(neighbours, getKey(new_i, j))
			}
		}
		if new_i := i + 1; new_i < height {
			if current_elevation-get(elevations, getKey(new_i, j)) <= 1 {
				neighbours = append(neighbours, getKey(new_i, j))
			}
		}
		if new_j := j - 1; new_j >= 0 {
			if current_elevation-get(elevations, getKey(i, new_j)) <= 1 {
				neighbours = append(neighbours, getKey(i, new_j))
			}
		}
		if new_j := j + 1; new_j < width {
			if current_elevation-get(elevations, getKey(i, new_j)) <= 1 {
				neighbours = append(neighbours, getKey(i, new_j))
			}
		}
		return neighbours
	}
	visited := map[string]bool{start_key: true}
	queue := [][]string{}
	queue = append(queue, []string{start_key})
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]
		for _, neighbour := range getNeighbours(node) {
			var current_path []string
			current_path = append(current_path, path...)
			current_path = append(current_path, neighbour)
			if goal_elevation == 0 {
				if neighbour == end_key {
					return current_path
				}
			} else {
				if get(elevations, neighbour) == goal_elevation {
					return current_path
				}
			}
			_, has_node := visited[neighbour]
			if !has_node {
				visited[neighbour] = true
				queue = append(queue, current_path)
			}
		}
	}
	return []string{start_key}
}

func puzzle1(input string) {
	elevations := make(map[string]int)
	var start_key string
	var end_key string
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)
	for i, line := range lines {
		for j, char := range line {
			key := getKey(i, j)
			if string(char) == "S" {
				elevations[key] = int('a')
				start_key = key
			} else if string(char) == "E" {
				elevations[key] = int('z')
				end_key = key
			} else {
				elevations[key] = int(char)
			}
		}
	}
	path := bfs(start_key, end_key, elevations, width, height, 0)
	println(len(path) - 1)
}

func puzzle2(input string) {
	elevations := make(map[string]int)
	var start_key string
	var end_key string
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)
	for i, line := range lines {
		for j, char := range line {
			key := getKey(i, j)
			if string(char) == "S" {
				elevations[key] = int('a')
				// start_key = key
			} else if string(char) == "E" {
				elevations[key] = int('z')
				start_key = key
			} else {
				elevations[key] = int(char)
			}
		}
	}
	path := bfs(start_key, end_key, elevations, width, height, int('a'))
	println(len(path) - 1)
}

func main() {
	raw_input, _ := os.ReadFile("./day12_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
