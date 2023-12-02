#!/usr/bin/env python3

def get_num(line):
    digits = [int(str(c)) for c in line if str(c).isdigit()]
    return digits[0] * 10 + digits[-1]


if __name__ == "__main__":
    s = 0
    with open("input.txt") as f:
        for line in f.readlines():
            s += get_num(line)
    print(s)


