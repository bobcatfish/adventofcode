#!/usr/bin/env python3


races = [
#        (7, 9),
#        (15, 40),
#        (30, 200),
        (41, 244),
        (66, 1047),
        (72, 1228),
        (66, 1040),
]


def num_wins(time, distance):
    count = 0
    for charge_time in range(1, time -1):
        travel_time = time - charge_time
        if travel_time * charge_time > distance:
            count += 1
    return count


if __name__ == "__main__":
    nums = []
    for race in races:
        nums.append(num_wins(race[0], race[1]))

    total = 1
    for n in nums:
        total *= n
    print(nums, total)
