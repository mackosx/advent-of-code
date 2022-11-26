import os
from utils import utils
from math import ceil


def common_bit(input, position, f):
    bit_count = 0
    threshold = int(ceil(len(input) / 2))
    for line in input:
        bit_count += int(line[position])
    if f == "most":
        return "1" if bit_count >= threshold else "0"
    else:
        return "1" if bit_count < threshold else "0"


def get_rating(input, f):
    positions = len(input[0])
    filtered_list = input.copy()
    for position in range(positions):
        bit = common_bit(filtered_list, position, f)
        # Only keep items with common bit at current position
        filtered_list = [item for item in filtered_list if item[position] == bit]
        if len(filtered_list) == 1:
            return filtered_list[0]


def puzzle_1(input):
    # Keep track of occurences
    counts = [0 for _ in range(len(input[0].strip()))]
    for line in input:
        for i in range(len(line.strip())):
            counts[i] += int(line[i])
    gamma_str = "".join("1" if count > len(input) / 2 else "0" for count in counts)
    # Flip the bits for epsilon
    epsilon_str = "".join("1" if x == "0" else "0" for x in gamma_str)
    return int(epsilon_str, 2) * int(gamma_str, 2)


def puzzle_2(input):
    input = [line.strip() for line in input]
    co2_rating = get_rating(input, "least")
    oxygen_rating = get_rating(input, "most")
    # print(co2_rating, oxygen_rating)

    return int(oxygen_rating, 2) * int(co2_rating, 2)


parsed_input = utils.parse_file(
    os.path.join(os.path.dirname(__file__), "day3_input.txt"), as_int=False
)

# Testing
assert common_bit(["1011", "1111", "0000", "0000"], 0, f="most") == "1"
assert common_bit(["1011", "1111", "0000", "0000"], 1, f="most") == "0"
assert common_bit(["1111", "1111", "1000", "1000"], 0, f="most") == "1"
assert common_bit(["1111", "1111", "0000", "0000"], 0, f="least") == "0"
assert common_bit(["1111", "1111", "1000", "1000"], 0, f="least") == "0"
assert common_bit(["1010", "1011"], 1, f="most") == "0"
assert get_rating(["1010", "1011", "0110"], "most") == "1011"

result = puzzle_1(parsed_input)
print(result)

result = puzzle_2(parsed_input)
print(result)
