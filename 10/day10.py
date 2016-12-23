# http://adventofcode.com/2016/day/10
import collections
import re


class NoMoreMovesError(Exception):
    pass


class Bot(object):
    def __init__(self, id, low, high, chips):
        self.id = id
        self.low = low
        self.high = high
        self.chips = chips


Output = collections.namedtuple("Output", ["id", "chips"])


def process_bot_line(l):
    results = re.findall("bot (\d*) gives low to (bot|output) (\d*) and high to (bot|output) (\d*)", l)[0]
    bot_id, low_type, low, high_type, high = results
    return int(bot_id), low_type, int(low), high_type, int(high)


def make_bot(bots, output, curr_line):
    bot_id, low_type, low, high_type, high = process_bot_line(curr_line)
    if bot_id not in bots:
        bots[bot_id] = Bot(bot_id, None, None, [])

    low_recv, high_recv = None, None

    if low_type == "bot":
        if low not in bots:
            bots[low] = Bot(low, None, None, [])
        low_recv = bots[low]
    else:
        if low not in output:
            output[low] = Output(low, [])
        low_recv = output[low]
            

    if high_type == "bot":
        if high not in bots:
            bots[high] = Bot(high, None, None, [])
        high_recv = bots[high]
    else:
        if high not in output:
            output[high] = Output(high, [])
        high_recv = output[high]
            
    bots[bot_id].low = low_recv
    bots[bot_id].high = high_recv


def process_val_line(bots, l):
    val_id, bot_id = re.findall("value (\d*) goes to bot (\d*)", l)[0]
    return [int(x) for x in [val_id, bot_id]]


def get_value_to_bot(bots, l):
    val_id, bot_id = process_val_line(bots, l)

    if bot_id not in bots:
        bots[bot_id] = Bot(bot_id, None, None, [])

    bots[bot_id].chips.append(val_id)


def construct_bots(instr):
    bots, output = {}, {}
    unprocessed_lines = instr
    while len(unprocessed_lines) > 0:
        if unprocessed_lines[0].startswith("bot"):
            make_bot(bots, output, unprocessed_lines[0])
        if unprocessed_lines[0].startswith("value"):
            get_value_to_bot(bots, unprocessed_lines[0])
        del unprocessed_lines[0]

    return bots, output


def make_move(bots, output):
    for b in bots.values():
        if len(b.chips) == 2:
            high_chip = b.chips[0] if b.chips[0] > b.chips[1] else b.chips[1]
            low_chip = b.chips[0] if b.chips[0] < b.chips[1] else b.chips[1]
            
            print "%d to %s" % (high_chip, b.high)
            print "%d to %s" % (low_chip, b.low)

            b.high.chips.append(high_chip)
            b.low.chips.append(low_chip)
            b.chips = []
            return
    raise NoMoreMovesError()


def found_answer_part_1(bots):
    for b in bots.values():
        if 17 in b.chips and 61 in b.chips:
            return b
    return None


def find_output(outputs):
    val = 1
    for o in [outputs[0], outputs[1], outputs[2]]:
        for c in o.chips:
            val *= c

    return val

if __name__ == "__main__":
    with open("input") as f:
        instr = f.readlines()

    bots, output = construct_bots(instr)

    try:
        while True:
            make_move(bots, output)
    except NoMoreMovesError:
        val = find_output(output)
        print val

