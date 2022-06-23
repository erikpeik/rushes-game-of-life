package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func print_array(array [][]int, height int, width int) {
	var sb strings.Builder

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if array[i][j] == 1 {
				sb.WriteString("\u001b[41m  ")
			} else {
				sb.WriteString("\u001b[40;1m  ")
			}
		}
		if i != height-1 {
			sb.WriteString("\n")
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\033[H\033[2J")
	fmt.Println(sb.String())
	fmt.Print("\u001b[0m")
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
		par := []int{x, y, 1}
		changes = append(changes, par)
	}

	if num_of_live_neighbour < 2 && arr[x][y] == 1 {
		par := []int{x, y, 0}
		changes = append(changes, par)
	}

	if num_of_live_neighbour > 3 && arr[x][y] == 1 {
		par := []int{x, y, 0}
		changes = append(changes, par)
	}
}

func check_if_file_exist(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func full_check(width int, height int, arr [][]int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			check_rules(i, j, width, height, arr)
		}
	}
}

var changes [][]int

func main() {
	debug.SetGCPercent(-1)

	if len(os.Args) != 3 {
		println("Usage: ./life_animated initial_state iterations")
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

	for g := 0; g < iterations; g++ {
		changes = nil

		print_array(array, height, width)
		full_check(width, height, array)
		if len(changes) == 0 {
			break
		}
		for i := 0; i < len(changes); i++ {
			array[changes[i][0]][changes[i][1]] = changes[i][2]
		}
	}
	print_array(array, height, width)
}
