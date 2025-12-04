#!/usr/bin/env python3


def get_biggest(s):
    j1, j2 = s[0], 0
    for i in range(1, len(s)-1):
        if s[i] > j1:
            j1 = s[i]
            j2 = s[i+1]
        elif s[i] > j2:
            j2 = s[i]
    # last digit is not a candidate for first in joltage
    if s[len(s)-1] > j2:
        j2 = s[len(s)-1]
    return j1, j2


if __name__ == "__main__":
    banks = []
    with open("input.txt") as f:
        for line in [l.strip() for l in f.readlines()]:
            bank = []
            for c in line:
                bank.append(int(c))
            if bank:
                banks.append(bank)

    joltages = []
    for bank in banks:
        biggest = get_biggest(bank)
        joltages.append(int('{}{}'.format(*biggest)))
    print(sum(joltages))

