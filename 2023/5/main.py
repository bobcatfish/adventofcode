#!/usr/bin/env python3
import collections


Range = collections.namedtuple("Range", ["d_start", "s_start", "len"])


def get_seeds(line):
    return [int(seed.strip()) for seed in line.split(":")[1].split(" ") if seed.strip()]


def parse_group(group):
    nameline = group[0]
    vals = nameline.split(" ")[0].split("-")
    source, dest = vals[0], vals[2]

    ranges = []

    for line in group[1:]:
        vals = [int(v) for v in line.split(" ")]
        d_start, s_start, l = vals[0], vals[1], vals[2]
        ranges.append(Range(d_start, s_start, l))

    return source, dest, ranges


def get_pairs(seed, m):
    pairs = {"seed": seed}

    key = "seed"
    while key != "location":
        key, ranges = m[key]
        for r in ranges:
            if seed >= r.s_start and seed < r.s_start + r.len:
                offset = seed - r.s_start
                pairs[key] = r.d_start + offset
                seed = pairs[key]
                break
        else:
            pairs[key] = seed

    return pairs


if __name__ == "__main__":
    with open("input.txt") as f:
        groups = []
        group = []
        seeds = []
        for i, line in enumerate(f.readlines()):
            line = line.rstrip()
            if i == 0:
                seeds = get_seeds(line)
            elif i > 1:
                if line != "":
                    group.append(line)
                else:
                    groups.append(group)
                    group = []
        groups.append(group)

        print(seeds)

        m = {}

        for group in groups:
            source, dest, ranges = parse_group(group)
            m[source]= (dest, ranges)

        pairs = []
        for seed in seeds:
            pairs.append(get_pairs(seed, m))

        for p in pairs:
            print(p)

        loc = 1000000000000
        for p in pairs:
            if p["location"] < loc:
                loc = p["location"]

        print(loc)
