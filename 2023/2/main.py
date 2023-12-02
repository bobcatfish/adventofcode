#!/usr/bin/env python3
import collections
import re


Game = collections.namedtuple("Game", ["game_id", "rounds"])

bag = {"red": 12, "green": 13, "blue": 14}



def parse_game(line):
    pattern = re.compile("Game (\d*): (.*)$")
    m = pattern.match(line)
    if not m:
        raise Exception(f"no game id {line}")

    game_id = int(m.group(1))
    round_strs = m.group(2)
    rounds = []

    for s in round_strs.split(";"):
        showings = {}
        for ss in s.split(","):
            num, color = ss.strip().split(" ")
            showings[color] = int(num)
        rounds.append(showings)

    return Game(game_id=game_id, rounds=rounds)


def valid_game(game):
    for r in game.rounds:
        for color, num in r.items():
            if num > bag[color]:
                return False
    return True


if __name__ == "__main__":
    games = []
    with open("input.txt") as f:
        for line in f.readlines():
            games.append(parse_game(line))

    valid = []
    for game in games:
        if valid_game(game):
            valid.append(game)

    s = 0
    for game in valid:
        s += game.game_id

    print(s)
