# http://adventofcode.com/2016/day/13
import collections
import math


DESIGNER_NUM = 1352

X_MAX, Y_MAX = 45, 45
DESTINATION = (31, 39)
DIRECTIONS = [[-1, 0],
              [0, -1],
              [0, 1],
              [1, 0]]


def is_wall(x, y):
    dat_num = x*x + 3*x + 2*x*y + y + y*y + DESIGNER_NUM
    sum = 0
    for mask in [2**i for i in range(int(math.ceil(math.log(dat_num, 2))) + 1)]:
        if mask & dat_num:
            sum += 1

    return sum % 2 != 0


Location = collections.namedtuple("Location", ["x", "y", "moves"])


def look_for_arrival(next_locs):
    for l in next_locs:
        if l.x == DESTINATION[0] and l.y == DESTINATION[1]:
            return l
    return None


def get_next_locs(next_locs, visited_locs):
    new_next_locs = []
    for l in next_locs:
        for d in DIRECTIONS:
            new_next_locs.append(Location(l.x + d[0], l.y + d[1], l.moves + 1))

    new_next_locs = [l for l in new_next_locs if l.x >= 0 and l.y >= 0]
    new_next_locs = [l for l in new_next_locs if (l.x, l.y) not in visited_locs]
    new_next_locs = [l for l in new_next_locs if not is_wall(l.x, l.y)]

    for l in new_next_locs:
        visited_locs.add((l.x, l.y))

    return new_next_locs


def draw_graph():
    for y in range(Y_MAX):
        for x in range(X_MAX):
            if x == 31 and y == 39:
                print "B",
            else:
                if is_wall(x, y):
                    print "#",
                else:
                    print ".",
        print


if __name__ == "__main__":
    curr_loc = Location(1, 1, 0)
    main_visited_locs = set([(1, 1)])
    main_next_locs = [curr_loc]
    arrival_loc = None
    while not arrival_loc:
        main_next_locs = get_next_locs(main_next_locs, main_visited_locs)
        if main_next_locs[0].moves == 50:
            print len(main_visited_locs)
        arrival_loc = look_for_arrival(main_next_locs)
    
    print arrival_loc.moves
        
