package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
    p1 := 0
    p2 := 0
    mulEnabled := true
    mulPat := regexp.MustCompile(`mul\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)`)

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Cannot open file %v: %v\n", os.Args[1], err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`).FindAllString(line, -1)

        for _, part := range parts {
            if part == "do()" {
                mulEnabled = true
            } else if part == "don't()" {
                mulEnabled = false
            } else if mulMatch := mulPat.FindStringSubmatch(part); mulMatch != nil {
                x, _ := strconv.Atoi(mulMatch[1])
                y, _ := strconv.Atoi(mulMatch[2])
                p1 += x * y

                if mulEnabled {
                    p2 += x * y
                }
            }
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning: %v\n", err)
        return
    }

    fmt.Printf("P1: %v\n", p1)
    fmt.Printf("P2: %v\n", p2)
}
