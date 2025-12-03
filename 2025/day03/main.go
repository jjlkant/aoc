package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jjlkant/aoc/2025/utils"
)

func main() {
	lines, err := utils.ReadInputLines()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully read %d lines for Day 3.\n", len(lines))

	// --- PART 1 ---
	start := time.Now()
	resultPart1 := solvePart1(lines)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 Solution: %v (took %s)\n", resultPart1, elapsed)

	// --- PART 2 ---
	start = time.Now()
	resultPart2 := solvePart2(lines)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 Solution: %v (took %s)\n", resultPart2, elapsed)
}

func calculateJoltage(joltageValues []int) int {
	joltage := 0

	for _, joltageValue := range joltageValues {
		joltage = (joltage * 10) + joltageValue
	}
	return joltage
}

func getJoltageForBatteryBank(ratings string, nValues int) int {
	nRatings := len(ratings)

	// Use a fixed-size array for small nValues to avoid heap allocation
	var joltageValues [12]int // 12 is the max nValues used in this program
	// Only use the needed prefix
	joltageSlice := joltageValues[:nValues]

	for i := range nRatings {
		rating := ratings[i]
		// Fast digit conversion (assumes input is always '0'-'9')
		ratingValue := int(rating - '0')

		for idx, value := range joltageSlice {
			remainingValues := nRatings - i
			if nValues-idx > remainingValues {
				continue
			}
			if ratingValue > value {
				joltageSlice[idx] = ratingValue
				for j := idx + 1; j < nValues; j++ {
					joltageSlice[j] = 1
				}
				break
			}
		}
	}
	joltage := calculateJoltage(joltageSlice)
	return joltage
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string) interface{} {
	totalJoltage := 0
	for _, line := range lines {
		totalJoltage += getJoltageForBatteryBank(line, 2)
	}
	return totalJoltage
}

// This function holds your Part 2 logic.
func solvePart2(lines []string) interface{} {
	totalJoltage := 0
	for _, line := range lines {
		totalJoltage += getJoltageForBatteryBank(line, 12)
	}
	return totalJoltage
}
