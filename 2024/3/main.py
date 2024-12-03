#!/usr/bin/env python3
import re


R_OPS = r'(mul\(\d*,\d*\)|do\(\)|don\'t\(\))'
R_MULT = r'mul\((\d*),(\d*)\)'


def get_ops(lines):
    raw_ops, ops = [], []
    for l in lines:
        raw_ops.extend(re.findall(R_OPS, l))
    return raw_ops


def add_ops(ops):
    total = 0
    mults = []
    for o in ops:
        raw_mults = re.findall(R_MULT, o)
        for o in raw_mults:
            mults.append((int(o[0]), int(o[1])))
    for o in mults:
        total += o[0] * o[1]
    return total


def filter_ops(ops):
    filtered = []
    do = True
    for o in ops:
        if o.startswith("don"):
            do = False
        elif o.startswith("do"):
            do = True
        else:
            if do:
                filtered.append(o)
    return filtered



if __name__ == "__main__":
    lines = []
    with open("input.txt") as f:
        lines = f.readlines()

    raw_ops = get_ops(lines)
    print(add_ops(raw_ops))

    ops = filter_ops(raw_ops)
    print(add_ops(ops))
