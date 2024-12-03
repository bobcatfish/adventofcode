#!/usr/bin/env python3


MAX = 3


def get_nums(line):
    return [int(str(n)) for n in line.split()]


def all_meet(line, comp):
    for i in range(1, len(line)):
        d = line[i] - line[i-1]
        if not comp(d) or abs(d) > MAX:
            return False
    return True


def is_safe(line):
    return (all_meet(line, lambda d: d > 0) or
            all_meet(line, lambda d: d < 0))


def is_safe2(line):
    for i in range(0, len(line)):
        eline = line[:i] + line[i+1:]
        if is_safe(eline):
            return True
    return False


if __name__ == "__main__":
    nums = []
    with open("input.txt") as f:
        for line in f.readlines():
            nums.append(get_nums(line))

    p1, p2 = 0, 0
    for l in nums:
        if is_safe(l):
            p1 += 1
        elif is_safe2(l):
            p2 += 1

    print(p1 , p1 + p2)
