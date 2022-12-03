def puzzle_2():
    with open("./2022/day3_input.txt", "r") as file:
        input = file.read().splitlines()
    total = sum(
        list(
            map(
                lambda x: x - 96 if x >= 97 and x <= 123 else x - 38,
                [
                    ord(
                        set(input[i])
                        .intersection(set(input[i + 1]), set(input[i + 2]))
                        .pop()
                    )
                    for i in range(0, len(input), 3)
                ],
            )
        )
    )
    print(total)


puzzle_2()
