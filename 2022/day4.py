def puzzle_1():
    with open("./2022/day4_input.txt", "r") as file:
        input = file.read().splitlines()
    count = 0
    for line in input:
        pairs = [pair.split("-") for pair in line.split(",")]
        set1 = set(range(int(pairs[0][0]), int(pairs[0][1]) + 1))
        set2 = set(range(int(pairs[1][0]), int(pairs[1][1]) + 1))
        intersection = set1.intersection(set2)
        if intersection == set1 or intersection == set2:
            count += 1
    print(count)


def puzzle_2():
    with open("./2022/day4_input.txt", "r") as file:
        input = file.read().splitlines()
    count = 0
    for line in input:
        pairs = [pair.split("-") for pair in line.split(",")]
        set1 = set(range(int(pairs[0][0]), int(pairs[0][1]) + 1))
        set2 = set(range(int(pairs[1][0]), int(pairs[1][1]) + 1))
        intersection = set1.intersection(set2)
        if intersection:
            count += 1
    print(count)


puzzle_1()
puzzle_2()
