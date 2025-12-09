from math import copysign
import os

from python_utils.utils import parse_file


def main():
    parsed_input = parse_file(
        os.path.join(os.path.dirname(__file__), "day14_test.txt"), as_int=False
    )
    coords = []
    for line_text in parsed_input:
        line_coords = []
        for coords_text in line_text.split(" -> "):
            line_coords.append((int(coords_text[0]), int(coords_text[1])))
    # coords = [
    #     [
    #         (int(coord[0]), int(coord[1]))
    #         for coords_text in line_text.split(" -> ")
    #         for coord in coords_text.split(",")
    #     ]
    #     for line_text in parsed_input
    # ]
    occupied = set()
    # Build map
    for line_coords in coords:
        for i in range(1, len(line_coords)):
            start = line_coords[i - 1]
            end = line_coords[i]
            d_x = end[0] - start[0]
            d_y = end[1] - start[1]
            current_idx = 0 if d_x == 0 else 1
            magnitude = d_y if d_x == 0 else d_x
            current = [start[0], start[1]]
            sign = int(copysign(1, magnitude))
            while current[i] != end[i]:
                occupied.add(tuple(current))
                current[current_idx] += sign

    for coord in occupied:
        print(coord)

    units = 0
    # assert units == 24  # For test data


main()
