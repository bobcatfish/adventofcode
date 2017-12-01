# http://adventofcode.com/2016/day/18

NUM_ROWS = 40
NUM_ROWS = 400000
TRAP = '^'
SAFE = '.'


def parse_file(filename):
    with open(filename) as f:
        return f.readlines()[0].strip()


def count_safe_tiles(rows):
    return len([t for r in rows for t in r if t == SAFE])


def next_tile(t):
    if t[0] == TRAP and t[1] == TRAP and t[2] == SAFE:
        return TRAP
    elif t[1] == TRAP and t[2] == TRAP and t[0] == SAFE:
        return TRAP
    elif t[1] == SAFE and t[2] == SAFE and t[0] == TRAP:
        return TRAP
    elif t[0] == SAFE and t[1] == SAFE and t[2] == TRAP:
        return TRAP
    return SAFE


def next_row(row):
    next_row = []
    for i in range(len(row)):
        prev_t_0 = row[i-1] if i > 0 else SAFE
        prev_t_1 = row[i]
        prev_t_2 = row[i+1] if i < len(row) - 1 else SAFE

        next_row.append(next_tile((prev_t_0, prev_t_1, prev_t_2)))
    return next_row


if __name__ == "__main__":
    r = parse_file("input")
    rows = [r]
    for i in range(NUM_ROWS - 1):
        rows.append(next_row(rows[i]))
    #for r in rows:
    #    print "".join(r)

    print count_safe_tiles(rows)
