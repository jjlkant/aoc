package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ReadInputLines reads the content of a file named "input.txt"
// located in the caller's directory and returns it as a slice of strings (lines).
func ReadInputLines() ([]string, error) {
	// The key: Caller(1) gets the path of the file that *called* this function (e.g., day01/main.go)
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("failed to get caller information")
	}

	dir := filepath.Dir(filename)
	inputPath := filepath.Join(dir, "input.txt")

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("could not read input file at %s: %w", inputPath, err)
	}

	inputString := string(data)
	inputString = strings.TrimSpace(inputString)

	// Handles splitting and cleaning up empty lines
	lines := strings.Split(inputString, "\n")
	var result []string
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}

	return result, nil
}

func ReadInputGrid() ([][]string, error) {
	// The key: Caller(1) gets the path of the file that *called* this function (e.g., day01/main.go)
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("failed to get caller information")
	}

	dir := filepath.Dir(filename)
	inputPath := filepath.Join(dir, "input.txt")

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("could not read input file at %s: %w", inputPath, err)
	}

	inputString := string(data)
	inputString = strings.TrimSpace(inputString)

	// Handles splitting into grid and cleaning up empty lines
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
