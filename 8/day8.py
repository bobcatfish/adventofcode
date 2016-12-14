# http://adventofcode.com/2016/day/8
NUM_COLUMNS = 50
NUM_ROWS = 6

def shift(l, count):
	for i in range(count):
		end_val = l.pop()
		l.insert(0, end_val)


class RotateCommand(object):
	def __init__(self, args):
		self.axis, index, _, count = args.split(" ")
		self.count = int(count)
		_, i = index.split("=")
		self.i = int(i)

	def do_it(self, screen):
		if self.axis == "column":
			column = [screen[i][self.i] for i in range(NUM_ROWS)]
			shift(column, self.count)
			for i, c in enumerate(column):
				screen[i][self.i] = c
		elif self.axis == "row":
			shift(screen[self.i], self.count)
		

class RectCommand(object):
	def __init__(self, args):
		self.width, self.height = [int(a) for a in args.split("x")]

	def do_it(self, screen):
		for x in range(self.width):
			for y in range(self.height):
				screen[y][x] = "#"


command_map = {
	"rotate": RotateCommand,
	"rect": RectCommand
}


def get_instr(filename):
	with open(filename) as f:
		lines = f.readlines()

	commands = []
	for l in lines:
		command, args = l.split(" ", 1)
		commands.append(command_map[command](args))
		
	return commands


def print_screen(screen):
	for row in screen:
		for column in row:
			print column,
		print


if __name__ == "__main__":
	screen = [["."] * NUM_COLUMNS  for _ in range(NUM_ROWS)]

	commands = get_instr("input")
	for c in commands:
		c.do_it(screen)

	print_screen(screen)

	print len([p for column in screen for p in column if p == "#"])
