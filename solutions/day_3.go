package main

import (
    "fmt"
    "regexp"
    "strconv"
    "os"
    "bufio"
)

func get_input() string {
    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }

    input_str := ""

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        input_str += line
    }
    return input_str
}

func find_commands(arg_string string) []string {
    expression := `mul\(\d{1,3},\d{1,3}\)`
    r, _ := regexp.Compile(expression)
    matches := r.FindAllString(arg_string, -1)
    fmt.Println(matches)
    return matches
}

func calculate_commands(matches []string) int {
    total := 0
    expression := `\d{1,3}`
    r, _ := regexp.Compile(expression)
    for _, val := range matches {
        nums := r.FindAllString(val, -1)
        fmt.Println(nums[0], nums[1])
        num_1, _ := strconv.Atoi(nums[0])
        num_2, _ := strconv.Atoi(nums[1])
        total += num_1 * num_2
    }
    return total
    
}

func main() {
    //example_arg_string := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    arg_string := get_input()
    matches := find_commands(arg_string)
    fmt.Println(calculate_commands(matches))

}