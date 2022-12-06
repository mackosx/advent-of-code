def puzzle_1():
    with open("./2022/day6_input.txt", "r") as file:
        input = file.read()
        print(next(i + 4 for i in range(len(input)) if len(set(input[i : i + 4])) == 4))


def puzzle_2():
    with open("./2022/day6_input.txt", "r") as file:
        input = file.read()
        print(
            next(i + 14 for i in range(len(input)) if len(set(input[i : i + 14])) == 14)
        )


puzzle_1()
puzzle_2()
