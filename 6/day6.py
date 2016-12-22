# http://adventofcode.com/2016/day/6
import operator


if __name__ == "__main__":
    with open("input") as f:
        lines = f.readlines()
    solution = ""

    for i in range(len(lines[0])):
        letter_map = {}
        for l in lines:
            letter_map.setdefault(l[i], 0)
            letter_map[l[i]] += 1

        solution += str(sorted(letter_map.items(), key=operator.itemgetter(1))[0][0])

    print solution
