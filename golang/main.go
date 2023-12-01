package main

import (
	"aoc23/pkg/day1"
	"log"
	"time"
)


func main() {
    start := time.Now()
    day1.Solve()
    elapsed := time.Since(start)

    log.Printf("Took: %s", elapsed)
}
