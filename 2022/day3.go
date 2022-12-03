package main

import (
	"fmt"
	"os"
	"strings"
)

func puzzle1(input string) {
	rucksacks := strings.Split(input, "\n")
	var total_priority int
	for _, rucksack := range rucksacks {
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]
		var shared_item int
		for _, item := range compartment1 {
			if strings.Contains(compartment2, string(item)) {
				shared_item = int(item)
				if shared_item >= 97 && shared_item <= 123 {
					total_priority += shared_item - 96
				} else {
					total_priority += shared_item - 38
				}
				break
			}
		}
	}
	fmt.Println(total_priority)
}

func main() {
	raw_input, _ := os.ReadFile("./day3_input.txt")
	input := string(raw_input)
	puzzle1(input)
}
