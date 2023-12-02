#!/usr/bin/env python3
import collections
import re


Game = collections.namedtuple("Game", ["game_id", "rounds", "mins"])

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

    return Game(game_id=game_id, rounds=rounds, mins=get_mins(rounds))


def get_mins(rounds):
    min_bag = {"red": 0, "green": 0, "blue": 0}
    for r in rounds:
        for color, num in r.items():
            if num > min_bag[color]:
                min_bag[color] = num
    return min_bag


def valid_game(game):
    for r in game.rounds:
        for color, num in r.items():
            if num > bag[color]:
                return False
    return True


def valid_games(games):
    valid = []
    for game in games:
        if valid_game(game):
            valid.append(game)

    s = 0
    for game in valid:
        s += game.game_id

    print(s)


if __name__ == "__main__":
    games = []
    with open("input.txt") as f:
        for line in f.readlines():
            games.append(parse_game(line))

    valid_games(games)

    s = 0
    for game in games:
        s += game.mins["red"] * game.mins["green"] * game.mins["blue"]
    print(s)



