package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func puzzle1(input string) {
	groups := strings.Split(input, "\n\n")
	biggest := 0
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		group_slice := strings.Split(group, "\n")
		sum := 0
		for j := 0; j < len(group_slice); j++ {
			value, _ := strconv.Atoi(group_slice[j])
			sum += value
		}
		if sum > biggest {
			biggest = sum
		}
	}
	fmt.Println(biggest)

}

func puzzle2(input string) {
	groups := strings.Split(input, "\n\n")
	biggest_list := []int{0, 0, 0}
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		group_slice := strings.Split(group, "\n")
		sum := 0
		for j := 0; j < len(group_slice); j++ {
			value, _ := strconv.Atoi(group_slice[j])
			sum += value
		}
		sort.Ints(biggest_list)
		for x := 0; x < len(biggest_list); x++ {
			if sum > biggest_list[x] {
				biggest_list[x] = sum
				break
			}
		}
	}
	fmt.Println(sum(biggest_list))
}

func main() {
	raw_input, _ := os.ReadFile("./day1_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
