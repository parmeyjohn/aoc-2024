package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() ([]int, [][]int) {
	targets := []int{}
	values := [][]int{}

    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		split_line := strings.Split(line, ":")
		split_args := strings.Split(split_line[1][1:], " ")
		converted_vals := []int{}
		converted_target, _ := strconv.Atoi(split_line[0])
		for _, val := range split_args {
			new_val, _ := strconv.Atoi(val)
			converted_vals = append(converted_vals, new_val)
		}
		targets = append(targets, converted_target)
		values = append(values, converted_vals)
	
    }
    return targets, values
}

func can_apply_operator(index int, values []int, total int, target int) bool {
	if index == len(values) && total == target {
		return true
	}
	if index >= len(values) || total > target {
		return false
	}
	add := can_apply_operator(index + 1, values, total + values[index], target)
	mult := can_apply_operator(index + 1, values, total * values[index], target)

	return add || mult
}

func get_calibration_results(targets []int, values [][]int) int {
	total := 0
	for i, target := range targets {
		if can_apply_operator(1, values[i], values[i][0], target) {
			total += target
		}
	}

	return total
}

func main() {
	targets, values := get_input()
	// Part 1:	
	fmt.Println(get_calibration_results(targets, values))
}