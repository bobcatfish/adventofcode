# http://adventofcode.com/2016/day/1
import re

MOVEMENT = {
    'N': (0, 1),
    'E': (1, 0),
    'S': (0, -1),
    'W': (-1, 0),
}
DIRS = ['N', 'E', 'S', 'W']


def get_moves(filename):
    with open(filename) as f:
        return[(m[0], int(m[1]))
               for l in f.readlines()
               for m in  re.findall(r"(\w)(\d*)", l)]


def turn(curr_dir, turn):
    if turn == 'L':
        return DIRS[(DIRS.index(curr_dir) - 1) % len(DIRS)]
    if turn == 'R':
        return DIRS[(DIRS.index(curr_dir) + 1) % len(DIRS)]


def follow_moves(moves):
    curr_dir = 'N'
    loc = (0, 0)
    visited_twice = False
    visited_locations = set()
    for m in moves:
        curr_dir = turn(curr_dir, m[0])
        for _ in range(m[1]):
            loc = (loc[0] + MOVEMENT[curr_dir][0],
                   loc[1] + MOVEMENT[curr_dir][1])
            if loc in visited_locations and not visited_twice:
                print "Visted twice: %s" % str(loc)
                print blocks_away(loc)
                visited_twice = True
            visited_locations.add(loc)
    return loc


def blocks_away(loc):
    return sum(abs(l) for l in loc)


if __name__ == "__main__":
    moves = get_moves("input")
    final_loc = follow_moves(moves)
    print final_loc
    print blocks_away(final_loc)
