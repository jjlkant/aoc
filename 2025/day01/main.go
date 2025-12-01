// day01/main.go
package main

import (
	"fmt"
	"os"
	"strconv"

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
	resultPart1 := solvePart1(lines)
	fmt.Printf("Part 1 Solution: %v\n", resultPart1)

	// --- PART 2 ---
	resultPart2 := solvePart2(lines)
	fmt.Printf("Part 2 Solution: %v\n", resultPart2)
}

// Positive modulo, returns non negative solution to x % d
func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string) interface{} {
	counter := 0
	dial_value := 50
	for _, line := range lines {
		direction := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "L":
			dial_value -= amount
		case "R":
			dial_value += amount
		default:
			fmt.Println("Received unexpected input", line)
		}

		// Take the positive modulo function to wrap to [0,100)
		dial_value = pmod(dial_value, 100)

		if dial_value == 0 {
			counter += 1
		}
	}
	return counter
}

// This function holds your Part 2 logic.
func solvePart2(lines []string) interface{} {
	counter := 0
	dial_value := 50
	dial_zero := false
	for _, line := range lines {
		direction := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		n_rotations := amount / 100
		amount -= n_rotations * 100

		switch direction {
		case "L":
			dial_value -= amount
		case "R":
			dial_value += amount
		default:
			fmt.Println("Received unexpected input", line)
		}

		// Add to counter for every full rotation
		counter += n_rotations
		// Add to counter for passing zero within part of rotation (by knowing we didn't start there)
		if !dial_zero && (dial_value < 0 || dial_value > 100) {
			counter += 1
		}

		// Take the positive modulo function to wrap to [0,100)
		dial_value = pmod(dial_value, 100)

		// Add to counter if we end up at zero
		dial_zero = dial_value == 0
		if dial_zero {
			counter += 1
		}
	}
	return counter
}
