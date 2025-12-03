#!/usr/bin/env python3


def invalid(n):
    if len(n) % 2 != 0:
        return False
    return n[:len(n)//2] == n[len(n)//2:]


if __name__ == "__main__":
    ids = []
    with open("input.txt") as f:
        for line in f.readlines():
            for group in line.split(','):
                start, end = group.split('-')
                ids.append((int(start), int(end)))

    invalids = []
    for (start, end) in ids:
        while start <= end:
            if invalid(str(start)):
                invalids.append(start)
            start += 1

    print(sum(invalids))


