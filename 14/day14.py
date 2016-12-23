# http://adventofcode.com/2016/day/14
import collections
import hashlib


SALT = "qzyelonm"
NUM_KEYS = 64
LOOK_AHEAD_HASHES = 1000
STRETCH_COUNT = 2016


def contains_repeats(h, num):
    for i in range(len(h) - num + 1):
        chunk = h[i:i+num]
        if chunk == num * chunk[0]:
            return chunk[0]
    return None


def contains_repeating_char(h, num, char):
    for i in range(len(h) - num + 1):
        chunk = h[i:i+num]
        if chunk == num * char:
            return True
    return False



def is_key(hashes):
    char = contains_repeats(hashes[0], 3)
    if char:
        for i in range(1, LOOK_AHEAD_HASHES + 1):
            if contains_repeating_char(hashes[i], 5, char):
                return True

    return False


def get_next_hash(salt, i):
    salt = hashlib.md5("%s%d" % (salt, i)).hexdigest()
    for _ in range(STRETCH_COUNT):
        salt = hashlib.md5(salt).hexdigest()
    return salt


def get_initial_hashes():
    hashes = []
    index = 0
    for _ in range(LOOK_AHEAD_HASHES + 1):
        hashes.append(get_next_hash(SALT, index))
        index += 1
        
    return collections.deque(hashes), index


if __name__ == "__main__":
    keys = []
    hashes, i = get_initial_hashes()
    tries = 0
        
    while len(keys) < NUM_KEYS:
        if is_key(hashes):
            keys.append(hashes.popleft())
            print "%d/%d is key %d %s" % (len(keys), NUM_KEYS, tries, hashes[0])
        else:
            hashes.popleft()
        hashes.append(get_next_hash(SALT, i))
        i += 1
        tries += 1

    print tries - 1
