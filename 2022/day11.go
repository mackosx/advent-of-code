package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	throws_to          map[bool]int
	items              []int64
	test_divisible_num int
	operation          string
	throws             int
}

func parseMonkeys(monkey_setup []string) []*Monkey {
	// Brutal
	monkeys := make([]*Monkey, len(monkey_setup))
	for i, monkey_text := range monkey_setup {
		lines := strings.Split(monkey_text, "\n")
		items := []int64{}
		for _, worry_str := range strings.Split(strings.Split(lines[1], ": ")[1], ", ") {
			worry, _ := strconv.Atoi(worry_str)
			items = append(items, int64(worry))
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

func doOperation(old int64, operation string, test_num int) (result int64) {
	tokens := strings.Split(operation, " ")
	parseToken := func(token string) int64 {
		num := old
		if token != "old" {
			parsed, _ := strconv.Atoi(token)
			num = int64(parsed)
		}
		return num
	}
	first_num := parseToken(tokens[0])
	second_num := parseToken(tokens[2])
	operator := tokens[1]
	if operator == "*" {
		result = first_num * second_num
	} else {
		result = first_num + second_num
	}
	result = result % int64(test_num)
	return
}

func simulateRounds(monkeys []*Monkey, count int, divide bool) int64 {
	mod_num := 1
	for _, monkey := range monkeys {
		mod_num *= monkey.test_divisible_num
	}
	for i := 0; i < count; i++ {
		for _, monkey := range monkeys {
			for _, worry := range monkey.items {
				// Calculate new worry
				worry = doOperation(worry, monkey.operation, mod_num)
				if divide {
					worry /= 3
				}
				// Throw to another monkey
				test := worry%int64(monkey.test_divisible_num) == 0
				next_index := monkey.throws_to[test]
				next_monkey := monkeys[next_index]
				next_monkey.items = append(next_monkey.items, worry)
				monkey.throws++
			}
			monkey.items = []int64{}
		}
		if i+1 == 20 {
			println("State after round :", i+1)
			for i, monkey := range monkeys {
				fmt.Printf("Monkey: %d throws: %d\n", i, monkey.throws)
			}
		}
		println("round ", i+1, " done")
	}
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].throws > monkeys[j].throws })
	for _, monkey := range monkeys {
		fmt.Printf("%v\n", monkey)
	}
	return int64(monkeys[1].throws) * int64(monkeys[0].throws)
}

func puzzle1(input string) {
	monkey_setup := strings.Split(input, "\n\n")
	monkeys := parseMonkeys(monkey_setup)
	monkey_business := simulateRounds(monkeys, 20, true)
	println(monkey_business)
}

func puzzle2(input string) {
	monkey_setup := strings.Split(input, "\n\n")
	monkeys := parseMonkeys(monkey_setup)
	monkey_business := simulateRounds(monkeys, 10000, false)
	println(monkey_business)
}

func main() {
	raw_input, _ := os.ReadFile("./day11_input.txt")
	input := string(raw_input)
	// puzzle1(input)
	puzzle2(input)
}
