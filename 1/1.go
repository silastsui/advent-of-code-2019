package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strconv"
)

func getFuel(mass int) (fuel int) {
    fuel = mass / 3 - 2

    if (fuel <= 0) {
        return 0
    } 
    return fuel + getFuel(fuel) 
}

func main() {
    file, err := os.Open("1.in")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    fuel := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        mass, _ := strconv.Atoi(scanner.Text())
        fuel += getFuel(mass)
    }

    fmt.Println(fuel)
}
