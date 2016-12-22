# http://adventofcode.com/2016/day/9
import re


def get_marker(data, i):
    marker = ""
    while data[i] != ")":
        marker += data[i]
        i += 1
    return marker, i + 1


def extract_from_marker(marker):
    count, repeat = re.findall("(\d*)x(\d*)", marker)[0]
    count, repeat = [int(x) for x in (count, repeat)]
    return count, repeat


def decompress(l):
    output = ""
    i = 0
    while i < len(l):
        if l[i] == "(":
            marker, i = get_marker(l, i+1)
            count, repeat = extract_from_marker(marker)
            section_to_repeat = l[i:i+count]
            for _ in range(repeat):
                output += section_to_repeat
            i += count
        else:
            output += l[i]
            i += 1
    return output


def look_for_marker(l, i):
    length = 0
    if l[i] == "(":
        marker, i = get_marker(l, i+1)
        count, repeat = extract_from_marker(marker)

        section_to_repeat = l[i:i+count]

        section_i = 0
        total_increase = 0

        while section_i < len(section_to_repeat):
            increase, section_i = look_for_marker(section_to_repeat, section_i)
            total_increase += increase

        length += (repeat * total_increase)

        i += count
    else:
        length += 1
        i += 1
    return length, i


def count_decompressed_size(l):
    length = 0
    i = 0
    while i < len(l):
        increase, i = look_for_marker(l, i)
        length += increase
        print "count: %d" % i
    return length


if __name__ == "__main__":
    with open("input") as f:
        lines = [l.strip() for l in f.readlines()]

    #decompressed = []
    #for l in lines:
    #    decompressed.append(decompress(l))
    #print decompressed
    #print sum(len(l) for l in decompressed)

    print count_decompressed_size(lines[0])
