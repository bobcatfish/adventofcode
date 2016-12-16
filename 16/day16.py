# http://adventofcode.com/2016/day/16


STARTING_DATA = "10111100110001111"
DISK_SIZE = 35651584


def dragon_curve(data):
    a = data
    b = a[::-1]
    mask = int("".join(["1" for _ in range(len(b))]), 2)
    b_masked = int(b, 2) ^ mask
    b_masked_str = bin(b_masked)[2:]
    return a + '0' + b_masked_str.zfill(len(b))


def checksum(data):
    checksum = data
    while len(checksum) % 2 == 0:
        next_checksum = ""
        for i in range(0, len(checksum), 2):
            pair = checksum[i:i+2]
            next_checksum += '0' if int(pair[0], 2) ^ int(pair[1], 2) else '1'
        checksum = next_checksum
    return checksum


if __name__ == "__main__":
    data = STARTING_DATA
    while len(data) < DISK_SIZE:
        data = dragon_curve(data)
    data = data[:DISK_SIZE]
    print checksum(data)
