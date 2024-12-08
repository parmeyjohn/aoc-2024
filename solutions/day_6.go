package main

import (
	"bufio"
	"fmt"
	"os"
)

type Key struct {
    row int
    col int
}

func get_input() [][]rune {
    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }

	input_grid := [][]rune{}

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		input_grid = append(input_grid, []rune(line))
    }
    return input_grid
}

func move_direction(row int, col int, direction int) (int, int) {
	switch direction {
	case 0:
		row -= 1
	case 1:
		col += 1
	case 2:
		row += 1
	case 3:
		col -= 1
	}
	return row, col
}

func get_distinct_guard_positions(grid [][]rune, row int, col int) int {
	positions := 0
	seen := make(map[Key]bool)
	//directions := []string{"up", "right", "down", "left"}
	direction_index := 0
	for row < len(grid) && row >= 0 && col < len(grid[0]) && col >= 0 {
		for string(grid[row][col]) == "#" {
			//fmt.Println(string(grid[row][col]), directions[direction_index], row, col, positions)
			row, col = move_direction(row, col, (direction_index + 2) % 4)
			direction_index = (direction_index + 1) % 4
		}  
		_, in_seen := seen[Key{row, col}]
		if !in_seen {
			positions += 1
			seen[Key{row,col}] = true
		}
		
		row, col = move_direction(row, col, direction_index)
		
	}
	// subtract one to account for starting space
	return positions - 1
}

func main() {
	grid := get_input()
	// Example:
	// fmt.Println(get_distinct_guard_positions(grid, 6, 4))
	// Part 1:
	fmt.Println("Part 1:", get_distinct_guard_positions(grid, 62, 60))
}