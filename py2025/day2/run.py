from os import path
from posixpath import dirname


def get_input():
    filename = path.join(dirname(__file__), "input.txt")
    print(filename)
    with open(filename) as file:
        data = file.read().split(",")
    return list(range_item.split("-") for range_item in data)


def part1(id_ranges):
    invalid_ids = []
    for id_range in id_ranges:
        for num in range(int(id_range[0]), int(id_range[1])+1):
            str_num = str(num)
            if len(str_num) % 2 == 0:
                if str_num[0:len(str_num)//2] == str_num[len(str_num)//2:]:
                    invalid_ids.append(num)

    return sum(invalid_ids)

def part2(id_ranges):
    invalid_ids = []
    for id_range in id_ranges:
        for num in range(int(id_range[0]), int(id_range[1])+1):
            str_num = str(num)
            for sequence_len in range(1, (len(str_num)//2) + 1):
                if len(str_num) % sequence_len != 0:
                    continue
                parts = [str_num[i-sequence_len:i] for i in range(sequence_len, len(str_num) + 1, sequence_len)]
                if len(set(parts)) == 1:
                    invalid_ids.append(num)
                    break
    return sum(invalid_ids)


if __name__ == "__main__":
    input = get_input()
    answer = part2(input)
    print(answer)
