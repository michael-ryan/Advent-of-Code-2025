package main

import (
	"fmt"

	"mike-pr.com/AdventOfCode2025/registry"
)

func main() {
	for day, solve := range registry.All() {
		fmt.Printf("Day %d:\n", day)
		ans := solve()
		fmt.Printf("  Part 1: %v\n", ans.Part1)
		fmt.Printf("  Part 2: %v\n", ans.Part2)
	}
}
