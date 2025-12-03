#!/usr/bin/env python3


def invalid(n):
    if len(n) % 2 != 0:
        return False
    return n[:len(n)//2] == n[len(n)//2:]


def invalid2(n):
    for i in range(1, len(n)//2+1):
        seq = n[:i]
        found = True
        for j in range(i, len(n), len(seq)):
            if seq != n[j:j+len(seq)]:
                found = False
                break
        if found:
            return True
    return False


if __name__ == "__main__":
    ids = []
    with open("input.txt") as f:
        for line in f.readlines():
            for group in line.split(','):
                start, end = group.split('-')
                ids.append((int(start), int(end)))

    invalids, invalids2 = [], []
    for (start, end) in ids:
        while start <= end:
            if invalid(str(start)):
                invalids.append(start)
            if invalid2(str(start)):
                invalids2.append(start)
            start += 1

    print(sum(invalids))
    print(sum(invalids2))


