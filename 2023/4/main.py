#!/usr/bin/env python3
import math


def nums(vv):
    return [int(v) for v in vv.split()]


def eval_card(line):
    vv = [v.strip() for v in  line.split(":")[1].split("|")]
    wins, vals  = nums(vv[0]), nums(vv[1])

    count = 0
    for v in vals:
        if v in wins:
            count += 1

    return int(math.pow(2, count-1))


if __name__ == "__main__":
    with open("input.txt") as f:
        vals = []
        for line in f.readlines():
            vals.append(eval_card(line))
        print(vals)
        print(sum(vals))
