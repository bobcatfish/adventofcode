#!/usr/bin/env python3

import collections


Point = collections.namedtuple("Point", ["y", "x"])


def is_neighbor(p1, p2):
    return ((p1.y == p2.y and abs(p1.x - p2.x) == 1) or
            (p1.x == p2.x and abs(p1.y - p2.y) == 1))


if __name__ == "__main__":
    lines = []
    with open("input.txt") as f:
        lines = [l.strip() for l in f.readlines()]

    locs = collections.defaultdict(list)
    for y in range(len(lines)):
        for x in range(len(lines[0])):
            p = Point(y, x)
            if lines[y][x] != '.':
                locs[int(lines[y][x])].append(p)

    paths = [[p] for p in locs[0]]
    for i in range(1, 10):
        next_paths = []
        for path in paths:
            for cand in locs[i]:
                if is_neighbor(cand, path[-1]):
                    next_paths.append(path + [cand])
        paths = next_paths

    trailheads = collections.defaultdict(set)
    for p in paths:
        trailheads[p[0]].update([p[-1]])

    score = 0
    for p, t in trailheads.items():
        score += len(t)
    print("p1: ", score)
    print("p2: ", len(paths))
