package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jjlkant/aoc/2025/utils"
)

func main() {
	lines, err := utils.ReadInputLines()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully read %d lines for Day 1.\n", len(lines))

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

func getJoltageForBatteryBank(bank string, nValues int) int {
	ratings := []rune(bank)
	nRatings := len(ratings)

	joltageValues := make([]int, nValues)

	for i, rating := range ratings {
		ratingValue, err := strconv.Atoi(string(rating))
		if err != nil {
			panic(err)
		}

		for idx, value := range joltageValues {
			remainingValues := nRatings - i
			if nValues-idx > remainingValues {
				continue
			}
			if ratingValue > value {
				joltageValues[idx] = ratingValue
				for i := idx + 1; i < nValues; i++ {
					joltageValues[i] = 1
				}
				break
			}
		}
	}
	joltage := calculateJoltage(joltageValues)
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
