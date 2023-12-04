#!/usr/bin/env python3

import collections
import math


dirs = [
 (-1, 0),
 (-1, 1),
 (0, 1),
 (1, 1),
 (1, 0),
 (1, -1),
 (0, -1),
 (-1, -1),
]


def ispart(schem, y, x):
    for d in dirs:
        dy, dx = y + d[0], x + d[1]
        if (dy >= 0 and dy < len(schem) and dx >= 0 and dx < len(schem[0])):
            val = schem[dy][dx]
            if not val.isdigit() and val != ".":
                return val, (dy, dx)
    return None, None


def get_num(nums):
    num = 0
    for i in range(0, len(nums)):
        num += nums[i] * int(math.pow(10, i))
    return num


def get_parts(schem):
    num = []
    part = None
    parts = []
    gears = collections.defaultdict(list)
    for y in range(0, len(schem)):
        for x in range(0, len(schem[y])):
            if schem[y][x].isdigit():
                num.insert(0, int(schem[y][x]))
                if not part:
                    part, loc = ispart(schem, y, x)
                    if part == "*":
                        part = loc
            else:
                if part:
                    parts.append(get_num(num))
                    if len(part) > 1:
                        gears[part].append(parts[-1])
                num = []
                part = False
        if part:
            parts.append(get_num(num))
            if len(part) > 1:
                gears[part].append(parts[-1])
            num = []
            part = False
    return parts, gears


if __name__ == "__main__":
    with open("input.txt") as f:
        schem = [line.rstrip() for line in f.readlines()]

        parts, gears = get_parts(schem)
        print(sum(parts))

        v = 0
        for loc, vals in gears.items():
            if len(vals) == 2:
                v += vals[0] * vals[1]
        print(v)

