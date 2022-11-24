import os
import utils.utils as utils


def puzzle_1(input):
    count = 0
    for i in range(1, len(input)):
        if input[i] > input[i - 1]:
            count += 1
    return count


def puzzle_2(input):
    count = 0
    for i in range(3, len(input)):
        if sum([input[i], input[i - 1], input[i - 2]]) > sum(
            [input[i - 1], input[i - 2], input[i - 3]]
        ):
            count += 1
    return count


parsed_input = utils.parse_file(os.path.join(os.path.dirname(__file__), "input.txt"))

result = puzzle_1(parsed_input)
print(result)

result = puzzle_2(parsed_input)
print(result)
