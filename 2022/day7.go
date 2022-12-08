package main

import (
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name       string
	is_dir     bool
	total_size int
	parent     *Node
}

func parseCommands(commands []string) []*Node {
	var nodes = []*Node{}
	var current_node *Node
	for _, input_line := range commands {
		tokens := strings.Split(input_line, " ")
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if next_dir := tokens[2]; next_dir == ".." {
					current_node = current_node.parent
				} else {
					current_node = &Node{next_dir, true, 0, current_node}
					nodes = append(nodes, current_node)
				}
			}
		} else {
			size, _ := strconv.Atoi(tokens[0])
			node_ptr := &Node{tokens[1], false, size, current_node}
			nodes = append(nodes, node_ptr)
			// Update parent sizes
			parent_ptr := current_node
			for parent_ptr != nil {
				parent_ptr.total_size += size
				parent_ptr = parent_ptr.parent
			}
		}
	}
	return nodes
}

func puzzle1(input string) {
	commands := strings.Split(input, "\n")
	nodes := parseCommands(commands)
	sum := 0
	for _, node := range nodes {
		if node.total_size <= 100000 && node.is_dir {
			sum += node.total_size
		}
	}
	println(sum)
}

func puzzle2(input string) {

}

func main() {
	raw_input, _ := os.ReadFile("./day7_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
