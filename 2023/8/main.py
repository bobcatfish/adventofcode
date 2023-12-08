#!/usr/bin/env python3
import re


class Node:
    def __init__(self, name, l, r):
        self.name = name
        self.l = l
        self.r = r


def parse_map(lines):
    nodes = {}
    pattern = re.compile("(.{3}) = \((.{3}), (.{3})\)")

    for line in lines:
        m = pattern.match(line)
        if not m:
            raise Exception(f"unexpected format {line}")
        name, l, r = m.group(1), m.group(2), m.group(3)
        nodes[name] = Node(name, l, r)

    for _, node in nodes.items():
        node.l = nodes[node.l]
        node.r = nodes[node.r]

    return nodes


if __name__ == "__main__":
    with open("input.txt") as f:
        instr = f.readline().strip()
        f.readline()

        lines = f.readlines()
        nodes = parse_map(lines)

        curr = nodes["AAA"]

        steps = 0

        while curr.name != "ZZZ":
            d = instr[steps%len(instr)]
            if d == 'L':
                curr = curr.l
            elif d == 'R':
                curr = curr.r
            else:
                raise Exception("off the map {} {}".format(steps, d))
            steps += 1

        print(steps)



