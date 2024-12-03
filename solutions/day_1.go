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
// sort input numbers
// loop through and summarize

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

func problem_1_soln(l1 []int, l2 []int) int {
    gather_input()
    distance_sum := 0
    sort.Ints(l1)
    sort.Ints(l2)

    for i := range l1 {
        distance_sum += find_distance(l1[i], l2[i])
    }
    
    return distance_sum
}

func main() {
    list1, list2 := gather_input()
    fmt.Println("Hello, World!")
    fmt.Println(problem_1_soln(list1, list2))
}