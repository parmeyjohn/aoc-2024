package main

import (
    "fmt"
    "bufio"
    "os"
)


func get_input() []string {
    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }

    matrix := []string{}
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        matrix = append(matrix, line)
    }
    return matrix
}

func is_vertical_valid(matrix []string, row int, col int) bool {
    curr_string := ""
    for i := 0; i < 4; i++ {
        curr_string += string(matrix[row + i][col])
    }
    if curr_string == "XMAS" || curr_string == "SAMX" {
        return true
    }
    return false
}

func is_diag_valid(matrix []string, row int, col int) bool {
    curr_string := ""
    for i := 0; i < 4; i++ {
        curr_string += string(matrix[row + i][col + i])
    }
    if curr_string == "XMAS" || curr_string == "SAMX" {
        return true
    }
    return false

}

func is_diag_valid_backwards(matrix []string, row int, col int) bool {
    curr_string := ""
    for i := 0; i < 4; i++ {
        curr_string += string(matrix[row + i][col - i])
    }
    if curr_string == "XMAS" || curr_string == "SAMX" {
        return true
    }
    return false

}

func is_diag_valid_mas(matrix []string, row int, col int) bool {
    curr_string := ""
    for i := 0; i < 3; i++ {
        curr_string += string(matrix[row + i][col + i])
    }
    if curr_string != "MAS" && curr_string != "SAM" {
        return false
    }
    curr_string = ""
    for i := 0; i < 3; i++ {
        curr_string += string(matrix[row + i][col + 2 - i])
    }
    if curr_string == "MAS" || curr_string == "SAM" {
        return true
    }
    return false
    

    
    return false
}

func get_xmas_count(matrix []string) int {
    xmas_count := 0
    for i, row := range matrix {
        for j, val := range matrix[i] {
            if val == 'X' || val == 'S' {
                // left to right
                if j + 4 <= len(row) && (matrix[i][j:j+4] == "XMAS" || matrix[i][j:j+4] == "SAMX") {
                    xmas_count += 1
                }
                // top to bottom
                if i + 4 <= len(matrix) && is_vertical_valid(matrix, i, j) {
                     xmas_count += 1
                }
                // diag top left to bottom right
                if j + 4 <= len(row) && i + 4 <= len(matrix) && is_diag_valid(matrix, i, j) {
                    xmas_count += 1
                }
                // diag top right to bottom left
                if j - 3 >= 0 && i + 4 <= len(matrix) && is_diag_valid_backwards(matrix, i, j) {
                    xmas_count += 1
                }
            }
        }
    }
    return xmas_count
}

func get_mas_x_count(matrix []string) int {
    mas_x_count := 0
    for i, row := range matrix {
        for j, val := range matrix[i] { 
            // diag top left to bottom right
            if (val == 'M' || val == 'S') && 
                j + 3 <= len(row) && i + 3 <= len(matrix) && 
                is_diag_valid_mas(matrix, i, j) {
                mas_x_count += 1
            }
         }
    }
    return mas_x_count
}
func main() {
    input_matrix := get_input()
    fmt.Println("Part 1:", get_xmas_count(input_matrix))
    fmt.Println("Part 2:", get_mas_x_count(input_matrix))
}