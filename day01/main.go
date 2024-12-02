package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    p1 := 0
    p2 := 0
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Cannot open file %v: %v\n", os.Args[1], err)
        return
    }
    defer file.Close()

    lhs := []int{}
    rhs := []int{}
    cnt := make(map[int]int)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        parts := strings.Fields(scanner.Text())
        l, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
        r, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
        lhs = append(lhs, l)
        rhs = append(rhs, r)
        cnt[r]++ 
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Scanner error: %v\n", err)
        return
    }

    sort.Ints(lhs)
    sort.Ints(rhs)

    for i := 0; i < len(lhs); i++ {
        p1 += absDiff(lhs[i], rhs[i])
        p2 += lhs[i] * cnt[lhs[i]]
    }

    fmt.Printf("P1: %v\n", p1)
    fmt.Printf("P2: %v\n", p2)
}

func absDiff(x, y int) int {
    if x < y {
        return y - x
    }
    return x - y
}
