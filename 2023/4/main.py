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
    return count


def do_card(vals, i, t):
    count = 1

    n = i+1
    for ii in range(n, n+vals[i]):
        if ii >= len(vals):
            break
        count += do_card(vals, ii, t + "\t")

    return count


if __name__ == "__main__":
    with open("input.txt") as f:
        vals = []
        for line in f.readlines():
            vals.append(eval_card(line))
        print(sum([int(math.pow(2, count-1)) for count in vals]))

    count = 0
    for i in range(0, len(vals)):
        count += do_card(vals, i, "")
    print(count)
