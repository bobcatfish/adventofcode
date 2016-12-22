# http://adventofcode.com/2016/day/5
import hashlib


DOOR_ID = "uqwqemis"

PASSWORD_LEN = 8


def get_next_hash(door_id):
    i = 0
    while True:
        v = "%s%s" % (door_id, i)
        yield hashlib.md5(v).hexdigest()
        i += 1


def get_next_char(h):
    if h.startswith("00000"):
        try:
            pos = int(h[5])
        except ValueError:
            return None, None
        if pos < PASSWORD_LEN:
            return pos, h[6]
    return None, None


if __name__ == "__main__":
    password = ["_"] * PASSWORD_LEN
    for h in get_next_hash(DOOR_ID):
        next_pos, next_char = get_next_char(h)
        if (next_pos is not None and
                next_char is not None and
                password[next_pos] == "_"):
            password[next_pos] = next_char
            print "".join(password)
        if not any(c == "_" for c in password):
            break
    print "".join(password)
    
