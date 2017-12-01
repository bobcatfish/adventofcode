#http://adventofcode.com/2016/day/20
import bisect


def get_ranges(filename):
    ranges = []
    with open(filename) as f:
        for l in [l.strip() for l in f.readlines()]:
            start, end = l.split("-")   
            bisect.insort(ranges, (int(start), int(end)))
    return ranges


if __name__ == "__main__":
    ranges = get_ranges("input")
    last_end = 0
    highest_blocked_ip = None
    allowed_ips = []
    for start, end in ranges:
        if highest_blocked_ip == None:
            highest_blocked_ip = end
        else:
            if start > (highest_blocked_ip + 1):
                allowed_ips.extend([ip for ip in range(highest_blocked_ip + 1, start)])
            if end > highest_blocked_ip:
                highest_blocked_ip = end
    print len(allowed_ips)
