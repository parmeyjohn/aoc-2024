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

func find_commands_2(arg_string string) []string {
    commands := []string{}
    r, _ := regexp.Compile(`do\(\)(.*?)don't\(\)`)
    initial_matches := r.FindAllString(arg_string, -1)

    r_2, _ := regexp.Compile(`(.*?)don't\(\)`)
    beginning_matches := r_2.FindAllString(arg_string, 1)
    initial_matches = append(initial_matches, beginning_matches...)

    r_3, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
    for _, match_str := range initial_matches {
        new_matches := r_3.FindAllString(match_str, -1)
        commands = append(commands, new_matches...)
    }
    return commands

}

func calculate_commands(matches []string) int {
    total := 0
    expression := `\d{1,3}`
    r, _ := regexp.Compile(expression)
    for _, val := range matches {
        nums := r.FindAllString(val, -1)
        num_1, _ := strconv.Atoi(nums[0])
        num_2, _ := strconv.Atoi(nums[1])
        total += num_1 * num_2
    }
    return total
    
}

func main() {
    // example_arg_string := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    // example_arg_string_2 := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    arg_string := get_input()
    // Part 1:
    // matches := find_commands(arg_string)
    // fmt.Println(calculate_commands(matches))
    // Part 2:
    matches_2 := find_commands_2(arg_string)
    fmt.Println(calculate_commands(matches_2))

}