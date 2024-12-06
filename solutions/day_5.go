package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_rule_input() map[int][]int{
    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }

    rules := make(map[int][]int)

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        pages := strings.Split(line, "|")
        page_1, _ := strconv.Atoi(pages[0])
        page_2, _ := strconv.Atoi(pages[1])
        rules[page_1] = append(rules[page_1], page_2)
    }
    return rules
}

func get_update_input() [][]int {
    if len(os.Args) <= 2 {
        fmt.Println("Error: provide a valid txt file argument")
    }

    matrix := [][]int{}
    file, err := os.Open(os.Args[2])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        res_list := []int{}
        for _, val := range strings.Split(line, ",") {
            converted_val, _ := strconv.Atoi(val)
            res_list = append(res_list, converted_val)
        }
        
        matrix = append(matrix, res_list)
    }
    return matrix
}

func is_value_in_slice(target int, list []int) bool {
    for _, val := range list {
        if val == target {
            return true
        }
    }
    return false
}

func is_overlapping_slice(rules []int, seen map[int]bool) bool {
    for _, rule := range rules {
        _, ok := seen[rule]
        if ok {
            return true
        }
    }
    return false
}

func count_valid_updates(rules map[int][]int, updates [][]int) int {
    total_count := 0
    for _, update := range updates {
        seen := make(map[int]bool)
        invalid := false
        for _, page := range update {
            // for every new value we check if the vals we have already seen arent invalid
                _, ok := rules[page]
                if ok && is_overlapping_slice(rules[page], seen) {
                    invalid = true
                    break
                }
            
            seen[page] = true
        }
        
        // update is valid, sort and get middle value
        if !invalid {
            total_count += update[len(update) / 2]
        }
    }
    return total_count

}

func main() {
	updates := get_update_input()
	rules := get_rule_input()
    // fmt.Println(updates)
    // fmt.Println(rules)
    // Part 1:
    fmt.Println(count_valid_updates(rules, updates))

}