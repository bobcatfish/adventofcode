# http://adventofcode.com/2016/day/15
import copy
import re


class Disk(object):
    def __init__(self, positions, starting_position):
        self.positions = positions
        self.curr_position = starting_position

    def rotate(self):
        self.curr_position = (self.curr_position + 1) % self.positions

    def will_fall_at_time(self, time):
        if (self.curr_position + time) % self.positions == 0:
            return True

    def __repr__(self):
        return "Total positions: %d Current position: %d" % (self.positions, self.curr_position)


def get_disks(filename):
    disks = []
    with open(filename) as f:
        for l in [l.strip() for l in f.readlines()]:
            positions, starting_position = re.findall(
                "Disc #\d* has (\d*) positions; at time=0, it is at position (\d*).", l)[0]
            disks.append(Disk(int(positions), int(starting_position)))
    return disks


def will_fall_through(disks):
    time = 0

    for d in disks:
        time += 1
        if d.will_fall_at_time(time):
            continue
        else:
            return False
    return True


def rotate_disks(disks):
    for d in disks:
        d.rotate()


if __name__ == "__main__":
    disks = get_disks("input")
    disks.append(Disk(11, 0))
    time = 0
    while not will_fall_through(disks):
        time += 1
        rotate_disks(disks)
        if time % 10000 == 0:
            print "Time: %d" % time

    print time
