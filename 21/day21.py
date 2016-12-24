# http://adventofcode.com/2016/day/21
import re


UNSCRAMBLED_PASS = "abcdefgh"
SCRAMBLED_PASS = "fbgdceah"


class SwapPos(object):
    def __init__(self, x, y):
        self.x = int(x)
        self.y = int(y)

    def do_it(self, s):
        s_list = list(s)
        x_val = s_list[self.x]
        y_val = s_list[self.y]
        s_list[self.x] = y_val
        s_list[self.y] = x_val
        return "".join(s_list)

    def undo_it(self, s):
        return self.do_it(s)
    
    def __str__(self):
        return "Swap %d with %d" % (self.x, self.y)


class SwapLetter(object): 
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def do_it(self, s):
        s_list = list(s)
        for i, l in enumerate(s_list):
            if l == self.x:
                s_list[i] = self.y
            if l == self.y:
                s_list[i] = self.x
        return "".join(s_list)

    def undo_it(self, s):
        tempCommand = SwapLetter(self.y, self.x)
        return tempCommand.do_it(s)
    
    def __str__(self):
        return "Swap letter %s with %s" % (self.x, self.y)


class RotateSteps(object):
    def __init__(self, d, x):
        self.d = d
        self.x = int(x)

    def do_it(self, s):
        s_list = list(s)
        for _ in range(self.x):
            if self.d == "right":
                v = s_list.pop()
                s_list.insert(0, v)
            else:
                v = s_list[0]
                del s_list[0]
                s_list.append(v)
        return "".join(s_list)

    def undo_it(self, s):
        d = "left" if self.d == "right" else "right"
        tempCommand = RotateSteps(d, self.x)
        return tempCommand.do_it(s)
    
    def __str__(self):
        return "Rotate %s %d times" % (self.d, self.x)


class RotatePos(object):
    def __init__(self, x):
        self.x = x

    def _shifts_from_index(self, shifts):
        if shifts >= 4:
            shifts += 2
        else:
            shifts += 1
        return shifts

    def do_it(self, s):
        s_list = list(s)
        shifts = self._shifts_from_index(s_list.index(self.x))
        for _ in range(shifts):
            v = s_list.pop()
            s_list.insert(0, v)

        return "".join(s_list)

    def undo_it(self, s):
        original_pos = {
            0:7,
            1:0,
            2:4,
            3:1,
            4:5,
            5:2,
            6:6,
            7:3
        }
        s_list = list(s)
        pos = s_list.index(self.x)
        o = original_pos[pos]

        while s_list.index(self.x) != o:
            v = s_list[0]
            del s_list[0]
            s_list.append(v)
        return "".join(s_list)
    
    def __str__(self):
        return "Rotate based on %s" % (self.x)


class ReversePos(object):
    def __init__(self, x, y):
        self.x = int(x)
        self.y = int(y)

    def do_it(self, s):
        s_list = list(s)

        for i, j in zip(range(self.x, self.y+1),
                        range(self.y, self.x - 1, -1)):
            s_list[i] = s[j]

        return "".join(s_list)

    def undo_it(self, s):
        return self.do_it(s)

    def __str__(self):
        return "Reverse from %d to %d" % (self.x, self.y)


class MovePos(object):
    def __init__(self, x, y):
        self.x = int(x)
        self.y = int(y)

    def do_it(self, s):
        s_list = list(s)
        v = s_list.pop(self.x)
        s_list.insert(self.y, v)
        return "".join(s_list)

    def undo_it(self, s):
        tempCommand = MovePos(self.y, self.x)
        return tempCommand.do_it(s)
    
    def __str__(self):
        return "Move %d to %d" % (self.x, self.y)


def get_commands(filename):
    commands = []
    with open(filename) as f:
        for l in [l.strip() for l in f.readlines()]:
            if l.startswith("swap position"):
                x, y = re.findall("swap position (\d*) with position (\d*)", l)[0]
                commands.append(SwapPos(x, y))
            elif l.startswith("swap letter"):
                x, y = re.findall("swap letter (\w) with letter (\w)", l)[0]
                commands.append(SwapLetter(x, y))
            elif l.startswith("rotate based on"):
                x = re.findall("rotate based on position of letter (\w)", l)[0]
                commands.append(RotatePos(x))
            elif l.startswith("rotate"):
                d, x = re.findall("rotate (\S*) (\d*) step", l)[0]
                commands.append(RotateSteps(d, x))
            elif l.startswith("reverse positions"):
                x, y = re.findall("reverse positions (\d*) through (\d*)", l)[0]
                commands.append(ReversePos(x, y))
            elif l.startswith("move position"):
                x, y = re.findall("move position (\d*) to position (\d*)", l)[0]
                commands.append(MovePos(x, y))
            else:
                print "no match found for %s" % l
    return commands


if __name__ == "__main__":
    commands = get_commands("input")
    s = SCRAMBLED_PASS
    for c in commands[::-1]:
        print c
        s = c.undo_it(s)
        print s
    print s
