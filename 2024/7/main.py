#!/usr/bin/env python3

import collections


OPS = {
    '+': lambda x, y: x + y,
    '*': lambda x, y: x * y,
}


class Line:
    def __init__(self, s, vals):
        self.sum = s
        self.vals = vals
        self.ops = []


def get_lines(l):
    s = [int(sss.strip()) for ss in l.split(":") for sss in ss.split(" ")
        if not sss.isspace() and len(sss) > 0]
    return Line(s[0], s[1:])


def find_good(total, l, i):
    if total > l.sum:
        return []

    good_ops = []
    for os, o in OPS.items():
        new_total = o(total, l.vals[i+1])
        if i == len(l.vals) - 2:
            if new_total == l.sum:
                good_ops.append([os])
        else:
            ops = find_good(new_total, l, i+1)
            for oo in ops:
                good_ops.append([os] + oo)

    if i == 0:
        l.ops = good_ops
    return good_ops


def do_it(lines):
    good = []
    for l in lines:
        find_good(l.vals[0], l, 0)
        if len(l.ops) > 0:
            good.append(l)

    total = 0
    for l in lines:
        if len(l.ops) > 0:
            total += l.sum

    print(total)


if __name__ == "__main__":
    lines = []
    with open("input.txt") as f:
        for l in f.readlines():
            lines.append(get_lines(l))

    do_it(lines)

    for l in lines:
        l.ops = []
    OPS['||'] = lambda x, y: int(str(x) + str(y))
    do_it(lines)
