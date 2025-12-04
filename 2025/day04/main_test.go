package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// NOTE: Ensure your solving logic is in functions named solvePart1 and solvePart2 in main.go!

// Helper function to read the TEST input file
func readTestInput() ([][]string, error) {
	// Determines the directory of the currently running test file
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	inputPath := filepath.Join(dir, "test_input.txt")

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("could not read test input file at %s: %w", inputPath, err)
	}

	inputString := strings.TrimSpace(string(data))
	lines := strings.Split(inputString, "\n")

	var result [][]string
	for _, line := range lines {
		if line != "" {
			chars := strings.Split(line, "")
			result = append(result, chars)
		}
	}
	return result, nil
}

func TestPart1(t *testing.T) {
	const expected = 13

	grid, err := readTestInput()
	if err != nil {
		t.Fatalf("Failed to read test input: %v", err)
	}

	result := solvePart1(grid)

	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Part 1: Got %v, expected %v", result, expected)
	} else {
		t.Logf("Part 1: Passed (Result: %v)", result)
	}
}

func TestPart2(t *testing.T) {
	const expected = 43

	grid, err := readTestInput()
	if err != nil {
		t.Fatalf("Failed to read test input: %v", err)
	}

	result := solvePart2(grid)

	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Part 2: Got %v, expected %v", result, expected)
	} else {
		t.Logf("Part 2: Passed (Result: %v)", result)
	}
}
