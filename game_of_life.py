import sys
import numpy as np

from os import system
clear = lambda: system('clear')
from random import randint

def print_array(arr_2d):
	for i in arr_2d:
		for j in i:
			if (j == 1):
				print("X", end ="")
			else:
				print(".", end ="")
		print()

def is_in_bound(x, y, width, height, arr):
	if (x >= 0 and y >= 0 and x < height and y < width):
		return True
	else:
		return False

def check_neighbour(x, y, width, height, arr):
	num_of_live_neighbour = 0

	if (is_in_bound(x, y+1, width, height, arr) and arr[x][y+1] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x, y-1, width, height, arr) and arr[x][y-1] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x-1, y, width, height, arr) and arr[x-1][y] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x+1, y, width, height, arr) and arr[x+1][y] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x+1, y+1, width, height, arr) and arr[x+1][y+1] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x-1, y-1, width, height, arr) and arr[x-1][y-1] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x-1, y+1, width, height, arr) and arr[x-1][y+1] == 1):
		num_of_live_neighbour += 1
	if (is_in_bound(x+1, y-1, width, height, arr) and arr[x+1][y-1] == 1):
		num_of_live_neighbour += 1
	return num_of_live_neighbour

def full_check(width, heigth, arr):
	changes = []
	for i in range(heigth):
		for j in range(width):
			round = check_rules(i, j, width, heigth, arr)
			if (len(round)):
				changes.append(round)
	return changes

# Any live cell with fewer than two live neighbours dies, as if by underpopulation.
# Any live cell with two or three live neighbours lives on to the next generation.
# Any live cell with more than three live neighbours dies, as if by overpopulation.
# Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

def check_rules(x: int, y: int, width, height, map):
	num_of_live_neighbours = check_neighbour(x, y, width, height, map)
	changes = []
	#born rule
	if (num_of_live_neighbours == 3 and map[x][y] == 0):
		par = [x, y, 1]
		changes.append(par)

	#underpopulation rule
	if (num_of_live_neighbours < 2 and map[x][y] == 1):
		par = [x, y, 0]
		changes.append(par)

	#overpopulate rule
	if (num_of_live_neighbours > 3 and map[x][y] == 1):
		par = [x, y, 0]
		changes.append(par)
	return changes

def main():
	if len(sys.argv) != 3:
		print("Usage: python3 game_of_life.py inital_state iterations")
		return
	array = []
	width = 0
	heigth = 0
	initial_state = sys.argv[1]
	fd = open(initial_state, "r")
	lines = fd.readlines()
	for line in lines:
		if len(line) - 1 > width:
			width = len(line) - 1
		int_line = []
		for c in line:
			if c == '.':
				int_line.append(0)
			elif c == 'X':
				int_line.append(1)
		array.append(int_line)
		heigth += 1

	iterations = int(sys.argv[2])
	for g in range(iterations):
		changes = []
		changes = full_check(width, heigth, array)
		for i in range(len(changes)):
			changes[i] = changes[i][0]
		if len(changes) == 0:
			break
		for i in range(len(changes)):
			array[changes[i][0]][changes[i][1]] = changes[i][2]
	print_array(array)

main()
