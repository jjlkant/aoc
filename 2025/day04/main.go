package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jjlkant/aoc/2025/utils"
)

func main() {
	grid, err := utils.ReadInputGrid()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully read %d lines for Day 4.\n", len(grid))

	// --- PART 1 ---
	start := time.Now()
	resultPart1 := solvePart1(grid)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 Solution: %v (took %s)\n", resultPart1, elapsed)

	// --- PART 2 ---
	start = time.Now()
	resultPart2 := solvePart2(grid)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 Solution: %v (took %s)\n", resultPart2, elapsed)
}

type Offset struct {
	x int
	y int
}

// For better performance, use [][]byte instead of [][]string for grid.
func getAccessibleLocations(grid [][]byte, nIterations int, remove bool) int {
	neighborOffsets := []Offset{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	lenX := len(grid)
	lenY := len(grid[0])
	accessableLocations := 0
	for range nIterations {
		updated := false
		for i := range lenX {
			for j := range lenY {
				if grid[i][j] != '@' {
					continue
				}
				hits := 0
				for _, neighborOffset := range neighborOffsets {
					checkX := i + neighborOffset.x
					checkY := j + neighborOffset.y
					if checkX < 0 || checkX >= lenX || checkY < 0 || checkY >= lenY {
						continue
					}
					if grid[checkX][checkY] == '@' {
						hits++
					}
					if hits >= 4 {
						break
					}
				}
				if hits >= 4 {
					continue
				}
				if remove {
					grid[i][j] = 'X'
				}
				updated = true
				accessableLocations++
			}
		}
		if !updated {
			break
		}
	}
	return accessableLocations
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(chars [][]byte) interface{} {
	return getAccessibleLocations(chars, 1, false)
}

// This function holds your Part 2 logic.
func solvePart2(chars [][]byte) interface{} {
	return getAccessibleLocations(chars, 1000, true)
}
