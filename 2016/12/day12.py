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
        if isinstance(self.src, Register):
            self.dst.val = self.src.val
        else:
            self.dst.val = self.src
        return i+1


class IncCommand(object):
    def __init__(self, dst):
        self.dst = dst

    def do_it(self, commands, i):
        self.dst.val += 1
        return i+1


class DecCommand(object):
    def __init__(self, dst):
        self.dst = dst

    def do_it(self, commands, i):
        self.dst.val -= 1
        return i+1


class JumpCommand(object):
    def __init__(self, flag, offset):
        self.flag = flag
        self.offset = offset

    def do_it(self, commands, i):
        if isinstance(self.flag, Register):
            flag = self.flag.val
        else:
            flag = self.flag
        
        if flag != 0:
            return i + self.offset
        else:
            return i+1
            

def get_commands(filename):
    registers = {key: Register(0) for key in REGISTERS}
    registers['c'].val = 1
    commands = []
    with open(filename) as f:
        for l in [l.strip() for l in f.readlines()]:
            t, arg1, arg2 = re.findall("(\w{3}) (\w*) ?([-\w]*)", l)[0]
            if t == "cpy":
                if arg1 in REGISTERS:
                    src = registers[arg1]
                else:
                    src = int(arg1)
                dst = registers[arg2]
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
                commands.append(JumpCommand(flag, int(arg2)))
            
    return commands, registers


if __name__ == "__main__":
    commands, registers = get_commands("input")
    i = 0
    while i < len(commands):
        i = commands[i].do_it(commands, i)
    print registers['a'].val
