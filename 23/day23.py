import math
import re


REGISTERS = 'abcd'


class Register(object):
    def __init__(self, val):
        self.val = val


class CopyCommand(object):
    def __init__(self, src, dst):
        self.src = src
        self.dst = dst

    def do_it(self, commands, i):
        if not isinstance(self.dst, Register):
            return i+1
        if isinstance(self.src, Register):
            self.dst.val = self.src.val
        else:
            self.dst.val = self.src
        return i+1


class MultCommand(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def do_it(self, commands, i):
        if isinstance(self.x, Register):
            x = self.x.val
        else:
            x = self.x

        if isinstance(self.y, Register):
            y = self.y.val
        else:
            y = self.y
        
        self.x.val = x * (x - 1)
        self.y.val = (x - 1)
        return i+1


class IncCommand(object):
    def __init__(self, dst):
        self.dst = dst

    def do_it(self, commands, i):
        if not isinstance(self.dst, Register):
            return i+1
        self.dst.val += 1
        return i+1


class DecCommand(object):
    def __init__(self, dst):
        self.dst = dst

    def do_it(self, commands, i):
        if not isinstance(self.dst, Register):
            return i+1
        self.dst.val -= 1
        return i+1


class JumpCommand(object):
    def __init__(self, flag, offset):
        self.flag = flag
        self.offset = offset

    def do_it(self, commands, i):
        if isinstance(self.offset, Register):
            #import pdb; pdb.set_trace()
            offset = self.offset.val
        else:
            offset = self.offset

        if isinstance(self.flag, Register):
            flag = self.flag.val
        else:
            flag = self.flag
        
        if flag != 0:
            new_i = i + offset
            if new_i >= 0 and new_i < len(commands) and new_i != i:
                return new_i
            else:
                return i+1
        else:
            return i+1


class ToggleCommand(object):
    def __init__(self, x, special_x=None):
        self.x = x
        self.special_x = x

    def do_it(self, commands, i):
        if not self.special_x:
            if isinstance(self.x, Register):
                flag = self.x.val + i
            else:
                flag = self.x + i
        else:
            flag = self.special_x
        
        #import pdb; pdb.set_trace()
        if flag >= len(commands) and flag >= 0:
            return i+1
        c = commands[flag]
        if isinstance(c, IncCommand):
            commands[flag] = DecCommand(c.dst)
        elif isinstance(c, DecCommand):
            commands[flag] = IncCommand(c.dst)
        elif isinstance(c, ToggleCommand):
            commands[flag] = IncCommand(c.x)
        elif isinstance(c, JumpCommand):
            commands[flag] = CopyCommand(c.flag, c.offset)
        elif isinstance(c, CopyCommand):
            commands[flag] = JumpCommand(c.src, c.dst)
        return i+1


def get_commands(filename):
    registers = {key: Register(0) for key in REGISTERS}
    commands = []
    with open(filename) as f:
        for l in [l.strip() for l in f.readlines()]:
            t, arg1, arg2 = re.findall("(\w{3,4}) ([-\w]*) ?([-\w]*)", l)[0]
            if t == "cpy":
                if arg1 in REGISTERS:
                    src = registers[arg1]
                else:
                    src = int(arg1)
                if arg2 in REGISTERS:
                    dst = registers[arg2]
                else:
                    dst = int(arg2)
                commands.append(CopyCommand(src, dst))
            elif t == "inc":
                commands.append(IncCommand(registers[arg1]))
            elif t == "dec":
                commands.append(DecCommand(registers[arg1]))
            elif t == "jnz":
                if arg1 in REGISTERS:
                    flag = registers[arg1]
                else:
                    flag = int(arg1)
                if arg2 in REGISTERS:
                    offset = registers[arg2]
                else:
                    offset = int(arg2)
                commands.append(JumpCommand(flag, offset))
            elif t == "mult":
                if arg1 in REGISTERS:
                    x = registers[arg1]
                else:
                    x = int(arg1)
                if arg2 in REGISTERS:
                    y = registers[arg2]
                else:
                    y = int(arg2)
                commands.append(MultCommand(x, y))
            elif t == "tgl":
                if arg1 in REGISTERS:
                    flag = registers[arg1]
                else:
                    flag = int(arg1)
                commands.append(ToggleCommand(flag))
            else:
                print "no match: %s" % l
            
    return commands, registers


if __name__ == "__main__":
    commands, registers = get_commands("input")
    registers['a'].val = 7
    i = 0

    # hacks
    registers['a'].val = math.factorial(12)
    registers['b'].val = 2
    registers['c'].val = 0
    registers['d'].val = 0
    i = 10

    ToggleCommand(24).do_it(commands, 0)
    ToggleCommand(22).do_it(commands, 0)
    ToggleCommand(20).do_it(commands, 0)
    ToggleCommand(18).do_it(commands, 0)

    while i < len(commands):
        i = commands[i].do_it(commands, i)
    print registers['a'].val
