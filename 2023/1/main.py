#!/usr/bin/env python3


nums = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

def get_nums_p1(line):
    return [int(str(c)) for c in line if str(c).isdigit()]


def get_nums_p2(line):
    i = 0
    digits = []
    while i < len(line):
        inc = 0
        sc = str(line[i])
        if sc.isdigit():
            digits.append(int(sc))
            inc = 1
        else:
            for num, val in nums.items():
                if line[i:].startswith(num):
                    digits.append(val)
                    inc = len(num) - 1
        if inc > 0:
            i += inc
        else:
            i += 1

    return digits



def get_num(digits):
    return digits[0] * 10 + digits[-1]


if __name__ == "__main__":
    s = 0
    with open("input.txt") as f:
        for line in f.readlines():
            digits = get_nums_p2(line)
            n = get_num(digits)
            s += n
    print(s)


