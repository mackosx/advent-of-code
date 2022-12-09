package main

import (
	"os"
	"strconv"
	"strings"
)

func isVisible(i int, j int, trees [][]int) bool {
	tree := trees[i][j]
	visible_left := true
	visible_right := true
	visible_top := true
	visible_bottom := true

	if i == 0 || j == 0 {
		return true
	}
	for a := i + 1; a < len(trees[i]); a++ {
		if trees[a][j] >= tree {
			visible_right = false
			break
		}
	}
	for b := i - 1; b >= 0; b-- {
		if trees[b][j] >= tree {
			visible_left = false
			break
		}
	}
	for c := j + 1; c < len(trees); c++ {
		if trees[i][c] >= tree {
			visible_bottom = false
			break
		}
	}

	for d := j - 1; d >= 0; d-- {
		if trees[i][d] >= tree {
			visible_top = false
			break
		}
	}

	return visible_bottom || visible_left || visible_top || visible_right
}

func puzzle1(input string) {
	lines := strings.Split(input, "\n")
	trees := make([][]int, len(lines[0]))
	for i, line := range lines {
		trees[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			height, _ := strconv.Atoi(string(line[j]))
			trees[i][j] = height
		}
	}

	visible_count := 0
	for i := 0; i < len(trees); i++ {
		tree_line := trees[i]
		for j := 0; j < len(tree_line); j++ {
			if isVisible(i, j, trees) {
				visible_count++
			}
		}
	}
	println(visible_count)
}

func puzzle2(input string) {

}

func main() {
	raw_input, _ := os.ReadFile("./day8_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
