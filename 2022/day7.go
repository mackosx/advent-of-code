package main

import (
	"fmt"
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

var nodes = make(map[string]*Node)

func pwd(node *Node) {
	parent_ptr := node
	path := ""
	for parent_ptr != nil {
		path = parent_ptr.name + "/" + path
		if parent_ptr.parent == nil {
			break
		}
		parent_ptr = parent_ptr.parent
	}
	println(path)
}

func puzzle1(input string) {
	commands := strings.Split(input, "\n")
	nodes["/"] = &Node{"/", true, 0, nil}
	var current_node *Node
	for _, input_line := range commands {
		tokens := strings.Split(input_line, " ")
		switch tokens[0] {
		case "$":
			switch tokens[1] {
			case "cd":
				println(input_line)
				if next_dir := tokens[2]; next_dir == ".." {
					current_node = current_node.parent
					continue
				} else {
					current_node, _ = nodes[next_dir]
				}
				pwd(current_node)
			case "ls":
				continue
			}
		case "dir":
			node_ptr := &Node{tokens[1], true, 0, nil}
			nodes[tokens[1]] = node_ptr
			node_ptr.parent = current_node
		default:
			size, _ := strconv.Atoi(tokens[0])
			node_ptr := &Node{tokens[1], false, size, current_node}
			nodes[tokens[1]] = node_ptr

			// Update parent sizes
			parent_ptr := current_node
			for parent_ptr != nil {
				fmt.Printf("adding %d to %v from %s\n", size, *parent_ptr, node_ptr.name)
				parent_ptr.total_size += size
				fmt.Printf("Result %+v\n", *parent_ptr)
				parent_ptr = parent_ptr.parent
				if parent_ptr == nil {
					break
				}
			}

		}

	}
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
