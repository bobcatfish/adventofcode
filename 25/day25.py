import re


REGISTERS = 'abcd'
SIGNAL = ""


class Register(object):
    def __init__(self, val):
        self.val = val


class OutCommand(object):
    def __init__(self, x):
        self.x = x

    def do_it(self, commands, i):
        if isinstance(self.x, Register):
            val = self.x.val
        else:
            val = self.x
        global SIGNAL
        SIGNAL += str(val)
        return i+1


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
            offset = self.offset.val
        else:
            offset = self.offset

        if isinstance(self.flag, Register):
            flag = self.flag.val
        else:
            flag = self.flag
        
        if flag != 0:
            return i + offset
        else:
            return i+1


class ToggleCommand(object):
    def __init__(self, x):
        self.x = x

    def do_it(self, commands, i):
        if isinstance(self.x, Register):
            flag = self.x.val + i
        else:
            flag = self.x + i
        
        if flag >= len(commands) and flag >= 0:
            return i+1
        c = commands[flag]
        if isinstance(c, IncCommand):
            commands[flag] = DecCommand(c.dst)
        elif isinstance(c, DecCommand):
            commands[flag] = IncCommand(c.dst)
        elif isinstance(c, OutCommand):
            commands[flag] = IncCommand(c.x)
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
            t, arg1, arg2 = re.findall("(\w{3}) ([-\w]*) ?([-\w]*)", l)[0]
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
                if arg2 in REGISTERS:
                    offset = registers[arg2]
                else:
                    offset = int(arg2)
                commands.append(JumpCommand(flag, offset))
            elif t == "tgl":
                if arg1 in REGISTERS:
                    flag = registers[arg1]
                else:
                    flag = int(arg1)
                commands.append(ToggleCommand(flag))
            elif t == "out":
                if arg1 in REGISTERS:
                    flag = registers[arg1]
                else:
                    flag = int(arg1)
                commands.append(OutCommand(flag))
            else:
                print "no match: %s" % l
            
    return commands, registers

def run_with_init(init):
    commands, registers = get_commands("input")
    registers['a'].val = init
    i = 0
    while i < len(commands):
        i = commands[i].do_it(commands, i)
        if len(SIGNAL) >= 10:
            if SIGNAL.startswith("0101010101"):
                return True
            else:
                return False


if __name__ == "__main__":
    init = 1
    while True:
        if run_with_init(init):
            print "GOT IT %d" % init
            break
        else:
            if SIGNAL:
                print "Nope %d %s" % (init, SIGNAL)

        SIGNAL = ""
        init += 1
