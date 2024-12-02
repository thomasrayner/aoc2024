package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Cannot open file %v: %v\n", os.Args[1], err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning: %v\n", err)
        return
    }
}
