#!/usr/bin/env python3

import collections


SEARCH = "XMAS"

Point = collections.namedtuple('Point', ['y', 'x'])


neighbor_funcs = [
    lambda p: Point(p.y+1, p.x),
    lambda p: Point(p.y+1, p.x+1),
    lambda p: Point(p.y+1, p.x-1),
    lambda p: Point(p.y-1, p.x),
    lambda p: Point(p.y-1, p.x+1),
    lambda p: Point(p.y-1, p.x-1),
    lambda p: Point(p.y, p.x+1),
    lambda p: Point(p.y, p.x-1),
]


def found(puzz, p, n):
    fpoints = []
    for i in range(len(SEARCH)):
        if p.y >= len(puzz) or p.y < 0 or p.x >= len(puzz[0]) or p.x < 0:
            return False, []
        if puzz[p.y][p.x] != SEARCH[i]:
            return False, []
        fpoints.append(p)
        p = n(p)
    return True, fpoints


def find(puzz):
    count = 0
    fpoints = set()
    for y in range(len(puzz)):
        for x in range(len(puzz[y])):
            for n in neighbor_funcs:
                is_found, nfpoints = found(puzz, Point(y, x), n)
                if is_found:
                    count += 1
                    fpoints.update(nfpoints)

    for y in range(len(puzz)):
        for x in range(len(puzz[y])):
            p = Point(y, x)
            if p in fpoints:
                print(puzz[p.y][p.x], end='')
            else:
                print(".", end='')
        print("\n")
    return count


if __name__ == "__main__":
    puzz = []
    with open("input.txt") as f:
        puzz = f.readlines()

    count = find(puzz)
    print(count)

