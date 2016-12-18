# http://adventofcode.com/2016/day/3


def get_triangles(filename):
    t = []
    with open(filename) as f:
        for l in f.readlines(): 
            t.append([int(v) for v in l.split()])
    return t


def get_triangles_vertical(filename):
    t, lines = [], []
    with open(filename) as f:
        for l in f.readlines():
            lines.append([int(v) for v in l.split()])

    for i in range(0, len(lines), 3):
        t.extend([
            (lines[i][0], lines[i+1][0], lines[i+2][0]),
            (lines[i][1], lines[i+1][1], lines[i+2][1]),
            (lines[i][2], lines[i+1][2], lines[i+2][2]),
        ])
    return t
        

def is_valid(t):
    all_pairs_larger = True
    for i, v in enumerate(t):
        if i == 0:
            o_sides = [1, 2]
        elif i == 1:
            o_sides = [0, 2]
        elif i == 2:
            o_sides = [0, 1]
        if t[o_sides[0]] + t[o_sides[1]] <= v:
            all_pairs_larger = False
    return all_pairs_larger


if __name__ == "__main__":
    triangles = get_triangles_vertical("input")
    valid_triangles = [t for t in triangles if is_valid(t)]
    print len(valid_triangles)
