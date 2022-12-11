package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	throws_to          map[bool]int
	items              []int
	test_divisible_num int
	operation          string
	throws             int
}

func doOperation(old int, operation string) int {
	tokens := strings.Split(operation, " ")
	parseToken := func(token string) int {
		num := old
		if token != "old" {
			num, _ = strconv.Atoi(token)
		}
		return num
	}
	if tokens[1] == "*" {
		return parseToken(tokens[0]) * parseToken(tokens[2])
	} else {
		return parseToken(tokens[0]) + parseToken(tokens[2])
	}
}

func parseMonkeys(monkey_setup []string) []*Monkey {
	// Brutal
	monkeys := make([]*Monkey, len(monkey_setup))
	for i, monkey_text := range monkey_setup {
		lines := strings.Split(monkey_text, "\n")
		items := []int{}
		for _, worry_str := range strings.Split(strings.Split(lines[1], ": ")[1], ", ") {
			worry, _ := strconv.Atoi(worry_str)
			items = append(items, worry)
		}
		true_throw, _ := strconv.Atoi(string(lines[4][29]))
		false_throw, _ := strconv.Atoi(string(lines[5][30]))
		test_num, _ := strconv.Atoi(strings.Split(lines[3], "by ")[1])
		operation := strings.Split(lines[2], "= ")[1]
		monkeys[i] = &Monkey{
			throws_to:          map[bool]int{false: false_throw, true: true_throw},
			test_divisible_num: test_num,
			operation:          operation,
			items:              items,
		}
	}

	return monkeys
}

func puzzle1(input string) {
	monkey_setup := strings.Split(input, "\n\n")
	monkeys := parseMonkeys(monkey_setup)
	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for _, worry := range monkey.items {
				// Calculate new worry
				worry = doOperation(worry, monkey.operation) / 3
				// Throw to another monkey
				next_index := monkey.throws_to[worry%monkey.test_divisible_num == 0]
				monkeys[next_index].items = append(monkeys[next_index].items, worry)
				monkey.throws++
			}
			monkey.items = []int{}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].throws > monkeys[j].throws })
	println(monkeys[1].throws * monkeys[0].throws)

}

func puzzle2(input string) {

}

func main() {
	raw_input, _ := os.ReadFile("./day11_input.txt")
	input := string(raw_input)
	puzzle1(input)
	puzzle2(input)
}
