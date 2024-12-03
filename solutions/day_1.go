package main

import (
    "fmt"
    "sort"
    "math"
    "os"
    "bufio"
    "strings"
    "strconv"
)



// first gather input from text files
// part 1:
// sort input numbers
// loop through and summarize
// part 2:
// add left vals to dictionary
// loop through right vals to count
// calculate similarity score after

func gather_input() ([]int, []int) {
    if len(os.Args) <= 1 {
        fmt.Println("Error: provide a valid txt file argument")
    }
    list1 := []int{}
    list2 := []int{}

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error")
    }
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        cols := strings.Fields(line)
        val1, err1 := strconv.Atoi(cols[0])
        val2, err2 := strconv.Atoi(cols[1])
        if err1 != nil || err2 != nil {
            fmt.Println("Error converting str to int")
        }
        list1 = append(list1, val1)
        list2 = append(list2, val2)
    }

    return list1, list2
}

func find_distance(a int, b int) int {
    if (a > b) {
        return int(math.Abs(float64(a) - float64(b)))
    } 
    return int(math.Abs(float64(b) - float64(a)))
}

func solve_problem_1(l1 []int, l2 []int) int {
    gather_input()
    distance_sum := 0
    sort.Ints(l1)
    sort.Ints(l2)

    for i := range l1 {
        distance_sum += find_distance(l1[i], l2[i])
    }
    
    return distance_sum
}

func solve_problem_2(l1 []int, l2 []int) int {
    similarity_sum := 0
    seen_left := make(map[int]int)
    for _, val := range l1 {
        seen_left[val] = 0
    }

    for _, val := range l2 {
        _, ok := seen_left[val]
        if ok {
            seen_left[val] += 1
        }
    }

    for _, val := range l1 {
        similarity_sum += val * seen_left[val]
    }
    
    return similarity_sum
}

func main() {
    //example_list1 := []int{3, 4, 2, 1, 3, 3}
    //example_list2 := []int{4, 3, 5, 3, 9, 3}
    list1, list2 := gather_input()
    fmt.Println("Hello, World!")
    fmt.Println(solve_problem_1(list1, list2))
    fmt.Println(solve_problem_2(list1, list2))
}