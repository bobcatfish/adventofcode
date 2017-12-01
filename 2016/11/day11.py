# http://adventofcode.com/2016/day/11
import copy
import itertools
import re


POSSIBLE_ITEMS_IN_ELEVATOR = [1, 2]
ALL_SEEN_STATES = set()


class Generator(object):
    def __init__(self, t):
        self.type = t

    def __repr__(self):
        return "%s generator" % self.type

    def __eq__(self, other):
        return isinstance(other, Generator) and other.type == self.type


class Microchip(object):
    def __init__(self, t):
        self.type = t

    def __repr__(self):
        return "%s microchip" % self.type

    def __eq__(self, other):
        return isinstance(other, Microchip) and other.type == self.type


class State(object):
    def __init__(self, floors, current_floor_index, steps_taken):
        self.floors = floors
        self.steps_taken = steps_taken
        self.current_floor_index = current_floor_index

    def new_state(self):
        return State(
            [Floor([i for i in f.items], f.floor_i, f.has_elevator) for f in self.floors],
            self.current_floor_index,
            self.steps_taken
        )

    def __eq__(self, other):
        return self.floors == other.floors

    def finished(self):
        return all([f.is_empty() for f in self.floors[:-1]])


class Floor(object):
    def __init__(self, items, floor_i, has_elevator=False):
        self.items = items
        self.floor_i = floor_i
        self.has_elevator = has_elevator

    def is_empty(self):
        return len(self.items) == 0

    @property
    def matched_types(self):
        return sorted([g.type for m in self.generators for g in self.microchips if g.type == m.type])

    @property
    def matched_pairs(self):
        pairs = []
        for t in self.matched_types:
            for m in self.generators:
                for g in self.microchips:
                    if m.type == t and g.type == t:
                        pairs.append([g, m])
        return sorted(pairs)

    @property
    def unmatched_generators(self):
        return sorted([i for i in self.generators if i.type not in self.matched_types])

    @property
    def unmatched_microchips(self):
        return sorted([i for i in self.microchips if i.type not in self.matched_types])

    @property
    def generators(self):
        return sorted([i for i in self.items if isinstance(i, Generator)])

    @property
    def microchips(self):
        return sorted([i for i in self.items if isinstance(i, Microchip)])

    def is_irradiation(self):
        unmatched_types = list(
            set([i.type for i in self.generators]) ^
            set([i.type for i in self.microchips])
        )
        return any(u == i.type for i in self.microchips for u in unmatched_types) and len(self.generators) > 0

    def __eq__(self, other):
        if self.floor_i != other.floor_i:
            return False
        if self.has_elevator != other.has_elevator:
            return False
        if len(self.matched_types) != len(other.matched_types):
            return False
        if len(self.unmatched_generators) != len(other.unmatched_generators):
            return False
        if len(self.unmatched_microchips) != len(other.unmatched_microchips):
            return False
        return True


def remove_irradiation_states(states):
    return [s for s in states if not any(f.is_irradiation() for f in s.floors)]


def possible_combos(floor):
    possible_items = (floor.unmatched_generators +
                      floor.unmatched_microchips +
                      (floor.matched_pairs[0] if len(floor.matched_types) > 0 else []))
    for count in POSSIBLE_ITEMS_IN_ELEVATOR:
        for combo in itertools.combinations(possible_items, count):
            yield combo


def get_next_states_up(state):
    current_floor_i = state.current_floor_index
    current_floor = state.floors[current_floor_i]

    if current_floor_i == len(state.floors) - 1:
        return []

    items = copy.copy(current_floor.items)
    
    states = []
    
    for combo in possible_combos(current_floor):
        new_state = state.new_state()
        new_state.current_floor_index = current_floor_i + 1

        # No point moving up to an empty floor with one item
        if len(combo) == 1 and len(new_state.floors[current_floor_i +1].items) == 0:
            continue

        next_floor_items = copy.copy(new_state.floors[current_floor_i + 1].items)
        next_floor_items += combo
        new_state.floors[current_floor_i + 1] = Floor(next_floor_items, current_floor_i + 1, True)

        current_floor_items = list(set(items) - set(combo))
        new_state.floors[current_floor_i] = Floor(current_floor_items, current_floor_i, False)

        states.append(new_state)

    return states


def all_floors_below_empty(state):
    floors_below = [f for f in state.floors[:state.current_floor_index]]
    return all(f.is_empty() for f in floors_below)


def get_next_states_down(state):
    current_floor_i = state.current_floor_index
    current_floor = state.floors[current_floor_i]

    if current_floor_i == 0 or all_floors_below_empty(state):
        return []

    states = []

    for combo in possible_combos(current_floor):
        new_state = state.new_state()
        new_state.current_floor_index = current_floor_i - 1

        next_floor_items = [i for i in new_state.floors[current_floor_i - 1].items]
        next_floor_items += combo
        new_state.floors[current_floor_i - 1] = Floor(next_floor_items, current_floor_i - 1, True)

        current_floor_items = list(set(current_floor.items) - set(combo))
        new_state.floors[current_floor_i] = Floor(current_floor_items, current_floor_i, False)

        states.append(new_state)

    return states


def remove_waste_of_time_states(states):
    global ALL_SEEN_STATES
    return [s for s in states if s not in ALL_SEEN_STATES]


def get_next_states(state):
    global ALL_SEEN_STATES

    next_states = []

    states = get_next_states_up(state)
    states = remove_irradiation_states(states)
    states = remove_waste_of_time_states(states)

    next_states += states

    states = get_next_states_down(state)
    states = remove_irradiation_states(states)
    states = remove_waste_of_time_states(states)

    next_states += states
    
    for s in next_states:
        ALL_SEEN_STATES.add(s)

    return next_states


def parse_input(filename):
    floors = []
    with open(filename) as f:
        i = 0
        for l in f.readlines():

            generators = [Generator(m) for m in re.findall('([^ ]*) generator', l)]
            microchips = [Microchip(m) for m in  re.findall('([^ ]*)-compatible microchip', l)]

            f = Floor(generators + microchips, i)
            floors.append(f)
            i += 1
        
    floors[0].has_elevator = True
    return floors


def branch_to_next_step(game):
    print game.steps_taken

    states = get_next_states(game)
    games = [s for s in states]
    for s in states:
        s.steps_taken += 1

    return games


if __name__ == "__main__":
    main_floors = parse_input("input")
    main_state = State(main_floors, 0, 0)

    main_finished_games = []
    main_games = [main_state]

    main_seen_steps = [main_state]

    while not main_finished_games and len(main_games) > 0:
        main_next_games = []

        for g in main_games:
            main_branches = branch_to_next_step(g)
            for main_branch in main_branches:
                if main_branch.finished():
                    main_finished_games.append(main_branch)
                if main_branch not in main_seen_steps:
                    main_next_games.append(main_branch)
                    main_seen_steps.append(main_branch)
        main_games = main_next_games

    if main_finished_games:
        print "Steps in game: %d" % main_finished_games[0].steps_taken
        print "Num games: %d" % len(main_finished_games)
    else:
        print "no finished games!"
