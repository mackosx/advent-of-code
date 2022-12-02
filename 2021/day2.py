import os
from python_utils import utils


def puzzle_1(input):
    position = 0
    depth = 0
    for instruction in input:
        direction, amount = instruction.split(" ")
        amount = int(amount)
        if direction == "forward":
            position += amount
        elif direction == "up":
            depth -= amount
        elif direction == "down":
            depth += amount
    return position * depth


parsed_input = utils.parse_file(
    os.path.join(os.path.dirname(__file__), "day2_input.txt"), as_int=False
)


def puzzle_2(input):
    position = 0
    depth = 0
    aim = 0
    for instruction in input:
        direction, amount = instruction.split(" ")
        amount = int(amount)
        if direction == "forward":
            position += amount
            depth += aim * amount
        elif direction == "up":
            aim -= amount
        elif direction == "down":
            aim += amount
    return position * depth


result = puzzle_1(parsed_input)
print(result)

result = puzzle_2(parsed_input)
print(result)
