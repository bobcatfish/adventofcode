# http://adventofcode.com/2016/day/22
import hashlib
from time import gmtime, strftime


DIR = [
    (0, 1),
    (0, -1),
    (-1, 0),
    (1, 0)
]

WALL = '#'
STARTING_ITEM = '0'


class Game(object):
    def __init__(self, loc, found_goals):
        self.loc = loc
        self.found_goals = found_goals

    def add_found(self, i):
        self.found_goals |= (1 << i)

    def already_found(self, i):
        return self.found_goals & (1 << i)

    def reached_goal(self, goal):
        return self.found_goals == int('1' * (goal + 1), 2)

    def __hash__(self):
        return int(hashlib.md5("(%d,%d)-%s" % (self.loc[0], self.loc[1], self.found_goals)).hexdigest(), 16)

    def __eq__(self, other):
        return (self.loc == other.loc and
                self.found_goals == other.found_goals)


def get_maze(filename):
    maze = []
    goal_num = -1
    with open(filename) as f:
        for l in [l.strip() for l in f.readlines()]:
            row = list(l)
            maze.append(row)
            for x in row:
                if x.isdigit():
                    if int(x) > goal_num:
                        goal_num = int(x)
    return maze, goal_num


def get_next_steps(maze, loc_obj):
    loc = loc_obj.loc
    max_x, max_y = len(maze[0]), len(maze), 
    next_steps = []
    for d in DIR:
        new_loc = (loc[0] + d[0], loc[1] + d[1])
        if (new_loc[0] < max_x and new_loc[0] > 0 and
            new_loc[1] < max_y and new_loc[1] > 0):

            maze_item = maze[new_loc[1]][new_loc[0]]
            if maze_item != WALL:
                new_point = maze_item.isdigit() and not loc_obj.already_found(int(maze_item))
                game = Game(new_loc, loc_obj.found_goals)
                if new_point:
                    game.add_found(int(maze_item))
                next_steps.append(game)
    return next_steps


def get_next_steps_ignore_targets(maze, loc_obj):
    loc = loc_obj.loc
    max_x, max_y = len(maze[0]), len(maze), 
    next_steps = []
    for d in DIR:
        new_loc = (loc[0] + d[0], loc[1] + d[1])
        if (new_loc[0] < max_x and new_loc[0] > 0 and
            new_loc[1] < max_y and new_loc[1] > 0):
            maze_item = maze[new_loc[1]][new_loc[0]]
            if maze_item != WALL:
                game = Game(new_loc, loc_obj.found_goals)
                next_steps.append(game)
    return next_steps


def without_duplicates(seen_games, next_locs):
    cleaned_locs = [l for l in next_locs if l not in seen_games]
    for l in cleaned_locs:
        seen_games.add(l)
    return cleaned_locs


def find_starting_loc(maze):
    for y, row in enumerate(maze):
        for x, i in enumerate(row):
            if i == STARTING_ITEM:
                return (x, y)


if __name__ == "__main__":
    maze, goal_num = get_maze("input")
    starting_loc = find_starting_loc(maze)
    locs = [Game(starting_loc, 0)]
    if maze[starting_loc[1]][starting_loc[0]].isdigit():
        locs[0].add_found(int(maze[starting_loc[1]][starting_loc[0]]))

    seen_games = set([locs[0]])
    num_steps = 0
    wins = {}

    while len(wins.keys()) < goal_num:
        next_locs = []
        for loc in locs:
            next_locs.extend(without_duplicates(seen_games, get_next_steps(maze, loc)))

        num_steps += 1
        win_steps = [l for l in next_locs if l.reached_goal(goal_num)]
        if win_steps:
            wins[num_steps] = win_steps
        locs = [l for l in next_locs if not l.reached_goal(goal_num)]
        print "%s step %d, len %d" % (strftime("%Y-%m-%d %H:%M:%S", gmtime()),
                                      num_steps, len(locs))

    print num_steps

    win_seen_games = {steps: set() for steps in wins.keys()}

    next_steps = 0
    keep_going = True
    paths_back_found = []

    while len(paths_back_found) < len(wins):
        next_steps += 1
        for steps in wins.keys():
            next_locs = []
            for loc in wins[steps]:
                next_locs.extend(without_duplicates(win_seen_games[steps], get_next_steps_ignore_targets(maze, loc)))
            wins[steps] = next_locs
            if any(l.loc == starting_loc for l in wins[steps]):
                print steps + next_steps
                paths_back_found.append(steps + next_steps)
            
        print "%s step %d, total %d" % (strftime("%Y-%m-%d %H:%M:%S", gmtime()), next_steps, sum(len(l) for l in wins.values()))

    print paths_back_found
    print min(paths_back_found)
