package main

import (
	"fmt"
	"strconv"
	"strings"
)

func get_input(input_str string) map[int]int {
	input_map := make(map[int]int)
	for _, val := range strings.Fields(input_str) {
		int_val, _ := strconv.Atoi(val)
		if _, ok := input_map[int_val]; ok {
			input_map[int_val] = input_map[int_val] + 1
		} else {
			input_map[int_val] = 1
		}
	}
	return input_map
}

func old_blink(rocks []int, seen map[int][]int) []int {
	new_rocks := []int{}
	for _, num := range rocks {
		seen_val, in_seen := seen[num]
		if in_seen {
			new_rocks = append(new_rocks, seen_val...)
		} else if num == 0 {
			new_rocks = append(new_rocks, 1)
		} else if len(strconv.Itoa(num))%2 == 0 {
			num_str := strconv.Itoa(num)
			left_val, _ := strconv.Atoi(num_str[:len(num_str)/2])
			right_val, _ := strconv.Atoi(num_str[len(num_str)/2:])
			new_rocks = append(new_rocks, left_val)
			new_rocks = append(new_rocks, right_val)
			seen[num] = []int{left_val, right_val}
		} else {
			new_rocks = append(new_rocks, num*2024)
			seen[num] = []int{num * 2024}
		}
	}
	return new_rocks
}

func blink(rocks map[int]int) map[int]int {
	new_rocks := make(map[int]int)
	for num, count := range rocks {
		if num == 0 {
			new_rocks[1] += count
		} else if str_num := strconv.Itoa(num); len(str_num)%2 == 0 {
			num_str := strconv.Itoa(num)
			left_val, _ := strconv.Atoi(num_str[:len(num_str)/2])
			right_val, _ := strconv.Atoi(num_str[len(num_str)/2:])
			new_rocks[left_val] += count
			new_rocks[right_val] += count

		} else {
			new_rocks[num*2024] += count

		}
	}
	return new_rocks
}

func blink_by_amount(input map[int]int, amount int) map[int]int {
	rocks := input
	for i := 0; i < amount; i++ {
		rocks = blink(rocks)
	}
	return rocks
}

func get_rock_count(rocks map[int]int) int {
	count := 0
	for _, v := range rocks {
		count += v
	}
	return count
}

func main() {
	// example_input := get_input("125 17")
	input := get_input("337 42493 1891760 351136 2 6932 73 0")
	// Part 1:
	fmt.Println(get_rock_count(blink_by_amount(input, 25)))
	// Part 2:
	fmt.Println(get_rock_count(blink_by_amount(input, 75)))

}
