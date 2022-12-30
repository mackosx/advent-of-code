package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/scanner"
)

func sumRange(start int, end int, arr []int) int {
	sum := 0
	for i := start; i < end && i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func ordered(left []any, right []any) (bool, bool) {
	for i := 0; i < len(left) && i < len(right); i++ {
		l := left[i]
		r := right[i]
		r_int, r_is_int := r.(int)
		l_int, l_is_int := l.(int)
		r_arr, r_is_arr := r.([]any)
		l_arr, l_is_arr := l.([]any)

		if r_is_arr && l_is_arr {
			is_ordered, cont := ordered(l_arr, r_arr)
			if cont {
				// No decision, keep going
				continue
			} else {
				// Decision made
				return is_ordered, false
			}
		} else if r_is_int && l_is_int {
			if l_int < r_int {
				// Make a decision; ordered
				return true, false
			} else if l_int > r_int {
				// Make a decision; not ordered
				return false, false
			} else {
				// Values are equal, keep going
				continue
			}
		} else if r_is_int {
			is_ordered, cont := ordered(l_arr, []any{r_int})
			if cont {
				// No decision, keep going
				continue
			} else {
				// Decision made
				return is_ordered, false
			}
		} else {
			is_ordered, cont := ordered([]any{l_int}, r_arr)
			if cont {
				// No decision, keep going
				continue
			} else {
				// Decision made
				return is_ordered, false
			}
		}
	}

	// We've run out of items in the list
	if len(left) == len(right) {
		// No decision, keep going
		return false, true
	} else {
		// Whichever list is smaller ran out first
		return len(left) < len(right), false
	}
}

func parseList(s *scanner.Scanner) []any {
	list := []any{}
	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		char := s.TokenText()
		switch char {
		case "[":
			list = append(list, parseList(s))
		case "]":
			return list
		case ",":
			continue
		default:
			num, _ := strconv.Atoi(char)
			list = append(list, num)
		}
	}
	return list
}

func isOrdered(pairs []string) (o bool) {
	var s scanner.Scanner
	s.Init(strings.NewReader(pairs[0]))
	s.Scan()
	left := parseList(&s)
	s.Init(strings.NewReader(pairs[1]))
	s.Scan()
	right := parseList(&s)
	o, _ = ordered(left, right)
	return
}

func findIndex(s []any, f func(any) bool) int {
	for i, x := range s {
		if f(x) {
			return i
		}
	}
	return -1
}

func puzzle1(input string) {
	pairs_text := strings.Split(input, "\n\n")
	ordered_indices := []int{}
	for i, pair_text := range pairs_text {
		pairs := strings.Split(pair_text, "\n")
		if isOrdered(pairs) {
			ordered_indices = append(ordered_indices, i+1)
		}
	}
	println(sumRange(0, len(ordered_indices), ordered_indices))
}

func puzzle2(input string) {
	packet_text := strings.Split(input+"\n[[2]]\n[[6]]", "\n")
	packets := []any{}
	var s scanner.Scanner
	for _, packet_text := range packet_text {
		if packet_text == "" {
			continue
		}
		s.Init(strings.NewReader(packet_text))
		s.Scan()
		l := parseList(&s)
		packets = append(packets, l)
	}

	sort.Slice(packets, func(i, j int) bool {
		ordered, _ := ordered(packets[i].([]any), packets[j].([]any))
		return ordered
	})
	div1_index := findIndex(packets, func(packet any) bool { return fmt.Sprint(packet) == "[[2]]" }) + 1
	div2_index := findIndex(packets, func(packet any) bool { return fmt.Sprint(packet) == "[[6]]" }) + 1

	println(div1_index * div2_index)
}

func main() {
	raw_input, _ := os.ReadFile("./day13_input.txt")
	input := string(raw_input)
	// puzzle1(input)
	puzzle2(input)
}
