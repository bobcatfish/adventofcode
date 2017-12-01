# http://adventofcode.com/2016/day/2


#STARTING_LOC = (1, 1)
#KEYPAD = [
#    [1, 2, 3],
#    [4, 5, 6],
#    [7, 8, 9]
#]
STARTING_LOC = (0, 2)
KEYPAD = [
  [0, 0, 1, 0, 0],
  [0, 2, 3, 4, 0],
  [5, 6, 7, 8, 9],
  [0, 'A', 'B', 'C', 0],
  [0, 0, 'D', 0, 0]
]
MOVEMENTS = {
    'U': (0, -1),
    'L': (-1, 0),
    'R': (1, 0),
    'D': (0, 1)
}


def move(loc, b):
    movement = MOVEMENTS[b]
    new_loc = (loc[0] + movement[0],
               loc[1] + movement[1])
    
    if (new_loc[0] >= 0 and new_loc[0] < len(KEYPAD[0]) and
        new_loc[1] >= 0 and new_loc[1] < len(KEYPAD) and 
        KEYPAD[new_loc[1]][new_loc[0]] != 0):
        return new_loc 
    return loc


if __name__ == "__main__":
    with open("input") as f:
        buttons = [b.strip() for b in f.readlines() if b.strip()]

    loc = STARTING_LOC

    buttons_pressed = []
    for b in buttons:
        for i in b:
            loc = move(loc, i)
        buttons_pressed.append(KEYPAD[loc[1]][loc[0]])

    print "".join([str(b) for b in buttons_pressed])
