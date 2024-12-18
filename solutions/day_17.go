package main

import (
	"fmt"
	"math"
)

func get_combo(val int, A int, B int, C int) int {
	switch val {
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		return val
	}
}

func execute_instructions(input []int, A int, B int, C int) []int {
	i := 0
	output := []int{}

	for i < len(input) {
		opcode := input[i]
		operand := input[i+1]
		switch opcode {
		case 0:
			A = int(float64(A) / math.Pow(2, float64(get_combo(operand, A, B, C))))
		case 1:
			B = B ^ operand
		case 2:
			B = get_combo(operand, A, B, C) % 8
		case 3:
			if A == 0 {
				i += 2
				continue
			}
			i = operand
			continue
		case 4:
			B = B ^ C
		case 5:
			output = append(output, get_combo(operand, A, B, C)%8)
		case 6:
			B = int(float64(A) / math.Pow(2, float64(get_combo(operand, A, B, C))))
		case 7:
			C = int(float64(A) / math.Pow(2, float64(get_combo(operand, A, B, C))))
		}

		// fmt.Println(opcode, operand)
		i += 2
	}
	return output
}

func main() {
	fmt.Println("hello")
	// Example Input:
	// input := []int{0, 1, 5, 4, 3, 0}
	// fmt.Println(execute_instructions(input, 729, 0, 0))
	// Part 1:
	input := []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 5, 0, 3, 5, 5, 3, 0}
	fmt.Println(execute_instructions(input, 30344604, 0, 0))
}
