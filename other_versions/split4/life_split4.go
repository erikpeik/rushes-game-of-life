package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
)

func print_array(array [][]int, height int, width int) {
	var sb strings.Builder

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if array[i][j] == 1 {
				sb.WriteString("X")
			} else {
				sb.WriteString(".")
			}
		}
		if i != height-1 {
			sb.WriteString("\n")
		}
	}
	fmt.Println(sb.String())
}

func is_in_bound(x int, y int, width int, height int, arr [][]int) bool {
	if x >= 0 && y >= 0 && x < height && y < width {
		return true
	} else {
		return false
	}
}

func check_neighbour(x int, y int, width int, height int, arr [][]int) int {
	var num_of_live_neighbour = 0
	if is_in_bound(x, y+1, width, height, arr) && arr[x][y+1] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x, y-1, width, height, arr) && arr[x][y-1] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x-1, y, width, height, arr) && arr[x-1][y] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x+1, y, width, height, arr) && arr[x+1][y] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x+1, y+1, width, height, arr) && arr[x+1][y+1] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x-1, y-1, width, height, arr) && arr[x-1][y-1] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x-1, y+1, width, height, arr) && arr[x-1][y+1] == 1 {
		num_of_live_neighbour += 1
	}
	if is_in_bound(x+1, y-1, width, height, arr) && arr[x+1][y-1] == 1 {
		num_of_live_neighbour += 1
	}
	return num_of_live_neighbour
}

func check_rules(x int, y int, width int, height int, arr [][]int) {
	var num_of_live_neighbour = check_neighbour(x, y, width, height, arr)

	if num_of_live_neighbour == 3 && arr[x][y] == 0 {
		has_changed = true
		changes[x][y] = 1
	}

	if num_of_live_neighbour < 2 && arr[x][y] == 1 {
		has_changed = true
		changes[x][y] = 0
	}

	if num_of_live_neighbour > 3 && arr[x][y] == 1 {
		has_changed = true
		changes[x][y] = 0
	}
}

func check_if_file_exist(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func part_check(x1 int, y1 int, x2 int, y2 int, width int, height int, arr [][]int) {
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			check_rules(i, j, width, height, arr)
		}
	}
	defer wg.Done()
}

var changes [][]int

var wg sync.WaitGroup
var has_changed bool

func main() {
	debug.SetGCPercent(-1)

	if len(os.Args) != 3 {
		println("Usage: ./life_split4 initial_state iterations")
		os.Exit(1)
	}
	initial_state := os.Args[1]
	var file, err = os.OpenFile(initial_state, os.O_RDWR, 0644)
	if check_if_file_exist(err) {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var array [][]int
	var height = 0
	var width = 0
	for scanner.Scan() {
		height += 1
		line := scanner.Text()
		if len(line) > width {
			width = len(line)
		}
		var int_line []int

		for c := 0; c < len(line); c++ {
			if line[c] == 46 {
				int_line = append(int_line, 0)
			} else if line[c] == 88 {
				int_line = append(int_line, 1)
			}
		}
		array = append(array, int_line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	iterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
		return
	}
	changes = make([][]int, height)
	for i := 0; i < len(changes); i++ {
		changes[i] = make([]int, width)
		copy(changes[i], array[i])
	}
	for g := 0; g < iterations; g++ {
		has_changed = false
		wg.Add(4)
		part_check(0, 0, height/2, width/2, width, height, array)
		part_check(0, width/2, height/2, width, width, height, array)
		part_check(height/2, 0, height, width/2, width, height, array)
		part_check(height/2, width/2, height, width, width, height, array)
		wg.Wait()
		if !has_changed {
			break
		}
		for i := 0; i < len(array); i++ {
			copy(array[i], changes[i])
		}
	}
	print_array(array, height, width)
}
