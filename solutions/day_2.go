package main

import (
    "fmt"
    "math"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func solve_problem_1() int {
    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    safe_levels := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        report := scanner.Text()
        levels := strings.Fields(report)
        
        start_val, err1 := strconv.Atoi(levels[0])
        end_val, err2 := strconv.Atoi(levels[len(levels) - 1])
        if err1 != nil || err2 != nil {
            fmt.Println("Error converting strings to data")
        }

        prev_level := start_val
        increasing := true
        level_invalid := false

        if start_val > end_val { 
            increasing = false
        }
        
        for _, val := range levels[1:] {
            curr_level, err := strconv.Atoi(val)
            if err != nil {
                fmt.Println("Error converting string")
            }
            
            if increasing {
                if curr_level < prev_level {
                    level_invalid = true
                    break 
                }
            } else {
                if prev_level < curr_level {
                    level_invalid = true
                    break
                }
            }

            distance := find_distance(prev_level, curr_level)
            if distance > 3 || distance == 0 {
                level_invalid = true
                break
            }
            prev_level = curr_level

        }
        
        if !level_invalid {
            safe_levels += 1
        }
        
    }
    return safe_levels
}

func find_distance(a int, b int) int {
    if (a > b) {
        return int(math.Abs(float64(a) - float64(b)))
    } 
    return int(math.Abs(float64(b) - float64(a)))
}


func main() {
    fmt.Println(solve_problem_1())
}