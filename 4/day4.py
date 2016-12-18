# http://adventofcode.com/2016/day/4
import re


def get_rooms(filename):
    rooms = []
    with open(filename) as f:
        r_lines = [l.strip() for l in f.readlines()]
        for r in r_lines:
            letters, sector_id, checksum = re.findall(r"(.*)-(\d*)\[(.*)]", r)[0]
            letters = "".join(letters)
            rooms.append((letters, int(sector_id), checksum))
    return rooms


def expected_checksum(letters):
    counted_letters = {i: [] for i in range(len(letters))}
    for l in set(letters):
        if l != '-':
            counted_letters[letters.count(l)].append(l)
    counted_letters = {k: v for k, v in counted_letters.iteritems() if len(v) > 0}
    for v in counted_letters.values():
        v.sort()

    checksum = ""
    for i in range(len(letters) - 1, 0, -1):
        if i in counted_letters:
            for v in counted_letters[i]:
                for letter in v:
                    checksum += letter
                    if len(checksum) == 5:
                        return checksum

def shift(l, count):
    ascii_val = ord(l) - 97
    shifted_val = (ascii_val + count) % 26
    return chr(shifted_val + 97)


def decrypt_name(room):
    letters = []
    for l in room[0]:
        if l == "-":
            letters.append(" ")
        else:
            letters.append(shift(l, room[1]))
    return "".join(letters)


if __name__ == "__main__":
    rooms = get_rooms("input")
    valid_rooms = [r for r in rooms if r[2] == expected_checksum(r[0])]

    print sum(room[1] for room in valid_rooms)

    for room in valid_rooms:
        name = decrypt_name(room)
        # print name
        if name == "northpole object storage":
            print room
        
        
