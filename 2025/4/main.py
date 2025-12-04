#!/usr/bin/env python3

import collections


Point = collections.namedtuple('Point', ['y', 'x'])

Left = Point(0, -1)
Right = Point(0, 1)
Up = Point(-1, 0)
Down = Point(1, 0)
UpperLeft = Point(-1, -1)
UpperRight = Point(-1, 1)
LowerLeft = Point(1, -1)
LowerRight = Point(1, 1)

Neighbors = [
    Left,
    Right,
    Up,
    Down,
    UpperLeft,
    UpperRight,
    LowerLeft,
    LowerRight
]

ROLL = '@'
MAX = 4


def is_roll(loc, grid):
    return grid[loc.y][loc.x] == ROLL


def in_grid(nn, grid):
    return (nn.y >= 0 and nn.y < len(grid) and
            nn.x >= 0 and nn.x < len(grid[0]))


def num_rolls(loc, grid):
    count = 0
    for n in Neighbors:
        nn = Point(loc.y + n.y, loc.x + n.x)
        if in_grid(nn, grid):
            if is_roll(nn, grid):
                count += 1
    return count


def get_accessible(grid):
    accessible = []
    for y in range(0, len(grid)):
        for x in range (0, len(grid[0])):
            p = Point(y, x)
            if is_roll(p, grid):
                c = num_rolls(p, grid)
                if c < 4:
                    accessible.append(p)

    return accessible


def remove(grid):
    rolls = get_accessible(grid)
    for loc in rolls:
        grid[loc.y][loc.x] = '.'
    return len(rolls)


if __name__ == "__main__":
    grid = []
    with open("input.txt") as f:
        for line in f.readlines():
            grid.append(list(line))


    count = 0
    while True:
        removed = remove(grid)
        print("Removed: ", removed)
        if removed == 0:
            break
        count += removed

    print("Total ", count)


