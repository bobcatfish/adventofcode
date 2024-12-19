#!/usr/bin/env python3

import collections


Point = collections.namedtuple("Point", ["y", "x"])
ReachedPoint = collections.namedtuple("ReachedPoint", ["p", "path"])


GRID_MAX = 6
STEP_MAX = 12
GRID_MAX = 70
STEP_MAX = 1024


def get_neighbors(p):
    n = [
        Point(p.y - 1, p.x),
        Point(p.y + 1, p.x),
        Point(p.y, p.x - 1),
        Point(p.y, p.x + 1),
    ]
    return [nn for nn in n if nn.y >= 0 and nn.y <= GRID_MAX and
                              nn.x >= 0 and nn.x <= GRID_MAX]


def bfs(fell):
    visited = set()
    start = Point(0, 0)
    to_visit = collections.deque([ReachedPoint(start, [start])])
    goal = Point(GRID_MAX, GRID_MAX)

    while(len(to_visit) > 0):
        loc = to_visit.popleft()
        if loc.p == goal:
            return loc
        if loc.p not in visited:
            for n in get_neighbors(loc.p):
                if n not in visited and n not in fell:
                    to_visit.append(ReachedPoint(n, loc.path + [n]))
            visited.update([loc.p])

    return None


def print_grid(path_points, fell):
    for y in range(GRID_MAX + 1):
        for x in range(GRID_MAX + 1):
            p = Point(y, x)
            if p in path_points:
                print('O', end="")
            elif p in fell:
                print('#', end="")
            else:
                print('.', end="")
        print()


if __name__ == "__main__":
    falls = []
    with open("input.txt") as f:
        for line in f.readlines():
            if line.strip():
                x, y = (int(s) for s in line.split(","))
                falls.append(Point(y, x))
    fell = set(falls[:STEP_MAX])

    loc = bfs(fell)
    path_points = set(loc.path)
    print_grid(path_points, fell)
    print("p1", len(loc.path) - 1)

    for i in range(len(falls) - STEP_MAX):
        fell = set(falls[:STEP_MAX + i])
        if i % 100 == 0:
            print(len(fell), "/", len(falls))
        loc = bfs(fell)
        if loc is None:
            block = falls[STEP_MAX + i - 1]
            print("{},{}".format(block.x, block.y))
            break
    else:
        print("none found")

