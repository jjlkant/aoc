package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

// Helper: Extract a column from a 2D int slice
func getColumn(matrix [][]int, colIdx int) []int {
	col := make([]int, len(matrix))
	for i := range matrix {
		col[i] = matrix[i][colIdx]
	}
	return col
}

// Helper: Calculate problem total for a column of values and an operator
func calcProblemTotal(values []int, op string) int {
	total := values[0]
	for j := 1; j < len(values); j++ {
		switch op {
		case "*":
			total *= values[j]
		case "+":
			total += values[j]
		}
	}
	return total
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string) any {
	operatorLineIdx := len(lines) - 1
	var worksheet [][]int
	for _, line := range lines[:operatorLineIdx] {
		fields := strings.Fields(line)
		result := make([]int, len(fields))
		for i, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				panic(err)
			}
			result[i] = n
		}
		worksheet = append(worksheet, result)
	}

	grandTotal := 0
	operators := strings.Fields(lines[operatorLineIdx])
	for i, op := range operators {
		col := getColumn(worksheet, i)
		grandTotal += calcProblemTotal(col, op)
	}
	return grandTotal
}

// This function holds your Part 2 logic.
func solvePart2(lines []string) any {
	operatorLineIdx := len(lines) - 1
	operators := strings.Fields(lines[operatorLineIdx])
	re := regexp.MustCompile(`[+|*]`)
	matches := re.FindAllStringSubmatchIndex(lines[operatorLineIdx], -1)
	problemSizes := make([]int, len(operators))
	for i := 0; i < len(matches)-1; i++ {
		problemSizes[i] = matches[i+1][0] - matches[i][0] - 1
	}
	problemSizes[len(operators)-1] = len(lines[0]) - matches[len(matches)-1][0]

	grandTotal := 0
	for i, op := range operators {
		problemSize := problemSizes[i]
		startIdx := matches[i][0]
		endIdx := startIdx + problemSize
		values := make([]int, problemSize)
		for _, line := range lines[:operatorLineIdx] {
			slice := line[startIdx:endIdx]
			for j := range problemSize {
				val, err := strconv.Atoi(string(slice[j]))
				if err == nil {
					values[j] = values[j]*10 + val
				}
			}
		}
		grandTotal += calcProblemTotal(values, op)
	}
	return grandTotal
}
