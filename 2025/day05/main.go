package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jjlkant/aoc/2025/utils"
)

func main() {
	lines, err := utils.ReadInputLinesWithEmpty()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully read %d lines for Day 5.\n", len(lines))

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

type Range struct {
	start int
	end   int
}

func simplifyFreshIngredients(lines []string) ([]Range, int) {
	blankLineIdx := 0
	var freshIngredients []Range
	for idx, line := range lines {
		if line == "" {
			blankLineIdx = idx
			break
		}
		ingredientRange := strings.Split(line, "-")
		ingredientRangeStart, err := strconv.Atoi(ingredientRange[0])
		if err != nil {
			panic(err)
		}
		ingredientRangeEnd, err := strconv.Atoi(ingredientRange[1])
		if err != nil {
			panic(err)
		}
		freshIngredients = append(freshIngredients, Range{ingredientRangeStart, ingredientRangeEnd})
	}

	sort.Slice(freshIngredients, func(i, j int) bool {
		return freshIngredients[i].start < freshIngredients[j].start
	})

	merged := []Range{freshIngredients[0]}
	for _, current := range freshIngredients[1:] {
		last := &merged[len(merged)-1]
		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged, blankLineIdx
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string) any {
	freshIngredients, blankLineIdx := simplifyFreshIngredients(lines)

	amountFreshIngredients := 0
	for _, line := range lines[blankLineIdx+1:] {
		ingredientId, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		for _, ingredientRange := range freshIngredients {
			// Given the sorted ranges, we can stop the loop if start > id
			if ingredientRange.start > ingredientId {
				break
			}
			if ingredientId >= ingredientRange.start && ingredientId <= ingredientRange.end {
				amountFreshIngredients += 1
				break
			}
		}
	}
	return amountFreshIngredients
}

// This function holds your Part 2 logic.
func solvePart2(lines []string) any {
	freshIngredients, _ := simplifyFreshIngredients(lines)

	numberFreshIngredients := 0
	for _, ingredients := range freshIngredients {
		numberFreshIngredients += (ingredients.end - ingredients.start + 1)
	}

	return numberFreshIngredients
}
