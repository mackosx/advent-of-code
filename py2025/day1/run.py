from os import path
from posixpath import dirname


def get_input():
    filename = path.join(dirname(__file__), "input.txt")
    print(filename)
    with open(filename) as file:
        data = file.read().split("\n")
    return list(filter(None, data))


def part1(input):
    count = 0
    current = 50

    for num in input:

        if num[0] == "R":
            current = (current + int(num[1:])) % 100
        else:
            current = (current - int(num[1:])) % 100

        if current == 0:
            count += 1
    return count


def part2(input):
    count = 0
    current = 50

    for num in input:
        direction = num[0]
        clicks = int(num[1:])
        hits = 0
        if direction == "R":
            hits = (current + clicks) // 100
            current = (current + clicks) % 100
        else:
            hits = (((100 - current) % 100) + clicks) // 100
            current = (current - clicks) % 100

        count += hits
    return count


if __name__ == "__main__":
    input = get_input()
    answer = part2(input)
    print(answer)
