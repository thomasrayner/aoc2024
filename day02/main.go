package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    p1, p2 := 0, 0
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Cannot open file %v: %v\n", os.Args[1], err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }

        if safe(line) {
            p1++
            p2++
        } else if tryPerms(line, safe) {
            p2++
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning file: %v\n", err)
        return
    }

    fmt.Printf("P1: %d\n", p1)
    fmt.Printf("P2: %d\n", p2)
}

func safe(line string) bool {
    nums := strings.Fields(line)
    ints := make([]int, len(nums))

    for i, num := range nums {
        val, err := strconv.Atoi(num)
        if err != nil {
            return false
        }
        ints[i] = val
    }

    isIncreasing := ints[1] > ints[0]

    for i := 0; i < len(ints)-1; i++ {
        diff := ints[i+1] - ints[i]
        if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) || diff < -3 || diff > 3 || diff == 0 {
            return false
        }
    }

    return true
}

func tryPerms(input string, handler func(string) bool) bool {
    nums := strings.Fields(input)
    for i := 0; i < len(nums); i++ {
        modified := append([]string{}, nums[:i]...)
        modified = append(modified, nums[i+1:]...)
        if handler(strings.Join(modified, " ")) {
            return true
        }
    }
    return false
}
