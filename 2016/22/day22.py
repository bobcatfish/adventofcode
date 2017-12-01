# http://adventofcode.com/2016/day/22
import copy
import collections
import re
from time import gmtime, strftime


USED_I = 0
AVAIL_I = 1

SWAP_TO_I = 0
SWAP_FROM_I = 1


class Disk(object):
    def __init__(self, x, y, used, avail):
        self.x = x
        self.y = y
        self.used = used
        self.avail = avail

    def __repr__(self):
        return "(%d, %d)" % (self.x, self.y)


DIR = [
    (0, 1),
    (0, -1),
    (-1, 0),
    (1, 0)
]


def load_disks(filename):
    with open(filename) as f:
        disk_lines = [l.strip() for l in f.readlines()]

    # skip first two irrelvant lines
    disk_lines = disk_lines[2:]

    disks = []
    for l in disk_lines:
        x, y, used, avail = re.findall("/dev/grid/node-x(\d*)-y(\d*)\s*\d*T\s*(\d*)T\s*(\d*)T\s*\d*%", l)[0]
        disks.append(Disk(int(x), int(y), int(used), int(avail)))

    sorted(disks, key=lambda x: x.x)
    x_len = disks[-1].x + 1
    sorted(disks, key=lambda x: x.y)
    y_len = disks[-1].y + 1

    disk_grid = []
    for _ in range(y_len):
        disk_grid.append([None] * x_len)
        
    for y in range(y_len):
        for x in range(x_len):
            disk_grid[y][x] = [[d.used, d.avail] for d in disks if d.y == y and d.x == x][0]
    
    goal = (len(disk_grid[0]) - 1, 0)
    empty = None
    full = []
        
    for y in range(y_len):
        for x in range(x_len):
            v = disk_grid[y][x]
            goal_v = disk_grid[goal[1]][goal[0]]
            if v[USED_I]  > ((goal_v[USED_I] + goal_v[AVAIL_I])):
                full.append((x, y))
            if v[USED_I] == 0:
                empty = (x, y)

    max_y = len(disk_grid)
    max_x = len(disk_grid[0])

    return Grid(max_x, max_y, full, empty), goal


def find_locs_directly_connected(disk_grid, x, y):
    max_y = disk_grid.max_y
    max_x = disk_grid.max_x
    
    locs = []
    for d in DIR:
        check_x, check_y = x + d[0], y + d[1]

        if (check_x >= 0 and check_x < max_x and
            check_y >=0 and check_y < max_y):

            check_loc = (check_x, check_y)
            if check_loc != (x, y) and check_loc not in disk_grid.full:
                locs.append(check_loc)
    return locs


Grid = collections.namedtuple("Grid", ["max_x", "max_y", "full", "empty"])

Game = collections.namedtuple("Game", ["empty", "goal"])
Game2 = collections.namedtuple("Game2", ["empty", "goal"])


def print_grid(grid, games):
    for y in range(grid.max_y):
        for x in range(grid.max_x):
            cell_format = "(%s)" if any(g.empty == (x,y) for g in games) else " %s "
            if (x,y) in grid.full:
                print cell_format % "#",
            else:
                print cell_format % ".",
        print


def get_next(grid, g, game1_seen):
    games = []

    locs = find_locs_directly_connected(grid, g.empty[0], g.empty[1])

    for new_loc in locs:
            if isinstance(g, Game) and new_loc == g.goal:
                if new_loc not in game1_seen:
                    print "Going to game 2!"
                    games.append(Game2(new_loc, g.empty))
                    game1_seen.add(new_loc)
            elif isinstance(g, Game):
                if new_loc not in game1_seen:
                    games.append(Game(new_loc, g.goal))
                    game1_seen.add(new_loc)

            elif isinstance(g, Game2):
                if new_loc == g.goal:
                    games.append(Game2(new_loc, g.empty))
                else:
                    games.append(Game2(new_loc, g.goal))

            else:
                print "don't know what to do"   
                import pdb; pdb.set_trace()
            
    return games


if __name__ == "__main__":
    grid, goal = load_disks("input")
    
    games = set([Game(grid.empty, goal)])
    game1_seen = set([grid.empty])

    steps = 0
    while not any(isinstance(g, Game2) and g.goal == (0, 0) for g in games):
        next_games = []
        for g in games:
            next_games.extend(get_next(grid, g, game1_seen))
        games = set(next_games)

        # Uncomment to watch algorithm in action
        #print_grid(grid, games)
        #import pdb; pdb.set_trace()

        steps += 1
        print "%s Step: %d Games: %d" % (strftime("%Y-%m-%d %H:%M:%S", gmtime()), steps, len(games))
    print steps
                
