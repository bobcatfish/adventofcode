import math

NUM_ELVES = 3001330


def get_victim(i, dist):
    victim = (i + 1) % len(dist)
    while victim != i:
        if not dist[victim]:
            victim = (victim + 1) % len(dist)
        else:
            return victim
    if victim == i:
        return None


def take_present(i, dist):
    if len(dist) % 1000 == 0:
        print len(dist)
    dist[i + 1]


def do_round(dist):
    for i in range(len(dist)):
        if dist[i]:
            victim = get_victim(i, dist)
            if victim is None:
                break
            else:
                dist[victim] = 0


if __name__ == "__main__":
    distribution = [1 for i in range(NUM_ELVES)]
    while len([i for i in distribution if i]) > 1:
        do_round(distribution)
    print "Part 1: %s" % ([i for i, e in enumerate(distribution) if e][0] + 1)
    
    print "Part 2: %s" % (NUM_ELVES - (3  ** int(math.floor(math.log(NUM_ELVES, 3)))))
