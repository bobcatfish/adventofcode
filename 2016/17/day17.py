# http://adventofcode.com/2016/day/17
import collections
import copy
import hashlib


PASS_CODE = "vwbaicqe"
VAULT_LOCATION = (3, -3)
UNLOCKED_DOOR_CODES = "bcdef"
MOVEMENTS = {
    'U': (0, 1),
    'D': (0, -1),
    'L': (-1, 0),
    'R': (1, 0), 
}

Step = collections.namedtuple("Step", ["location", "path"])


def get_unlocked_locations(step):
    h = hashlib.md5(step.path).hexdigest()
    doors = h[:4]
    doors_map = {
        'U': doors[0],
        'D': doors[1],
        'L': doors[2],
        'R': doors[3]
    }
    locations = []
    for d, movement in MOVEMENTS.iteritems():
        if doors_map[d] in UNLOCKED_DOOR_CODES:
            new_location = (step.location[0] + movement[0], step.location[1] + movement[1])
            if ((new_location[0] >= 0 and new_location[0] <= VAULT_LOCATION[0]) and
                (new_location[1] <= 0 and new_location[1] >= VAULT_LOCATION[1])):
                locations.append(Step(
                    (step.location[0] + movement[0], step.location[1] + movement[1]),
                     step.path + d))
                
    return locations


if __name__ == "__main__":
    next_steps = [Step((0, 0), PASS_CODE)]
    longest_path = None
    #while not any(s.location == VAULT_LOCATION for s in next_steps):
    while True:
        new_next_steps = []
        for step in next_steps:
            new_next_steps.extend(get_unlocked_locations(step))
        next_steps = new_next_steps
        #for step in next_steps:
        #    print step.path
        if any(s.location == VAULT_LOCATION for s in next_steps):
            longest_path = [s for s in next_steps if s.location == VAULT_LOCATION][0]
            print str(len(longest_path.path) - len(PASS_CODE))
            # paths end once they hit the end room
            next_steps = [s for s in next_steps if s.location != VAULT_LOCATION]
        if len(next_steps) == 0:
            print "No more steps!"
            break

    #for s in next_steps:
    #    if s.location == VAULT_LOCATION:
    #        print "SOLUTION: " + s.path
    print str(len(longest_path.path) - len(PASS_CODE))

