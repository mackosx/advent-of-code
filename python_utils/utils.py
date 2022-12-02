def parse_file(filepath, as_int=True):
    with open(filepath, "r") as file:
        lines = file.readlines()
        if as_int:
            return [int(line) for line in lines]
        return lines
