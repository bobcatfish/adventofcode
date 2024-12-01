#!/usr/bin/env python3
import collections


def get_nums(line):
    return (int(str(n)) for n in line.split())


def get_diff(l_nums, r_nums):
    d = 0
    for i in range(len(l_nums)):
        d += abs(l_nums[i]-r_nums[i])
    return d


def get_sim(l_nums, r_nums):
    counts = collections.defaultdict(int)
    for r in r_nums:
        counts[r] += 1

    return sum([l * counts[l] for l in l_nums])



if __name__ == "__main__":
    l_nums, r_nums = [], []
    with open("input.txt") as f:
        for line in f.readlines():
            l, r = get_nums(line)
            l_nums.append(l)
            r_nums.append(r)
    l_nums.sort()
    r_nums.sort()

    #print(l_nums, r_nums)
    d = get_diff(l_nums, r_nums)
    print(d)

    s = get_sim(l_nums, r_nums)
    print(s)

