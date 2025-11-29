package main

import (
	"fmt"
	"os"
	"strconv"

	"mike-pr.com/AdventOfCode2025/registry"
)

func main() {
	args := os.Args[1:]

	daysToSolve := make(map[int]bool)
	if len(args) == 0 {
		// do every day
		for day := range registry.All() {
			daysToSolve[day] = true
		}
	} else {
		// do some days as defined by args
		for _, day := range args {
			d, err := strconv.Atoi(day)
			if err != nil {
				fmt.Printf("could not interpret argument \"%v\" as an integer", day)
				return
			}
			daysToSolve[d] = true
		}
	}

	for day, solve := range registry.All() {
		if !daysToSolve[day] {
			continue
		}
		fmt.Printf("Day %d:\n", day)
		ans := solve()
		fmt.Printf("  Part 1: %v\n", ans.Part1)
		fmt.Printf("  Part 2: %v\n", ans.Part2)
	}
}
