#!/usr/bin/env python3

if __name__ == "__main__":
    l_nums, r_nums = [], []
    rotations = []
    with open("input.txt") as f:
        for line in f.readlines():
            d, rotato = line[0], int(line[1:])
            rotations.append((d, rotato))

    zero_count = 0
    curr = 50
    m = 100
    for d, rotato in rotations:
        if d == 'L':
            curr -= rotato
        else:
            curr += rotato
        curr %= m
        if curr == 0:
            zero_count += 1
        print(curr)

    print("count: ", zero_count)


