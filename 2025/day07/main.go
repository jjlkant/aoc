package main

import (
	"fmt"
	"os"
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

func solve(lines []string) (splits, timelines int) {
	beamIndex := strings.IndexRune(lines[0], 'S')
	beams := make([]int, len(lines[0]))
	beams[beamIndex] = 1
	for _, line := range lines[1:] {
		for i := range line {
			if beams[i] > 0 && line[i] == '^' {
				beams[i+1] += beams[i]
				beams[i-1] += beams[i]
				beams[i] = 0
				splits++
			}
		}
	}
	for _, beam := range beams {
		timelines += beam
	}
	return
}

func solvePart1(lines []string) any { splits, _ := solve(lines); return splits }
func solvePart2(lines []string) any { _, timelines := solve(lines); return timelines }
