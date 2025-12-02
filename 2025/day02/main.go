// day02/main.go
package main

import (
	"fmt"
	"os"
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

func expandRange(range_string string) []string {
	product_ids := strings.Split(range_string, "-")
	id_range_start, err := strconv.Atoi(product_ids[0])
	if err != nil {
		panic(err)
	}
	id_range_end, err := strconv.Atoi(product_ids[1])
	if err != nil {
		panic(err)
	}

	// Pre-allocate the slice for better performance
	expanded_ids := make([]string, id_range_end-id_range_start+1)
	for i := id_range_start; i <= id_range_end; i++ {
		expanded_ids[i-id_range_start] = fmt.Sprint(i)
	}
	return expanded_ids
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string) interface{} {
	total := 0
	line := lines[0]
	for range_str := range strings.SplitSeq(line, ",") {
		range_ids := expandRange(range_str)
		for _, id := range range_ids {
			id_length := len(id) // Use len instead of utf8.RuneCountInString for ASCII strings
			if id_length%2 == 0 {
				mid := id_length / 2
				if id[:mid] == id[mid:] {
					id_int_value, err := strconv.Atoi(id)
					if err != nil {
						panic(err)
					}
					total += id_int_value
				}
			}
		}
	}
	return total
}

// This function holds your Part 2 logic.
func solvePart2(lines []string) interface{} {
	total := 0
	line := lines[0]
	for range_str := range strings.SplitSeq(line, ",") {
		range_ids := expandRange(range_str)
		for _, id := range range_ids {
			id_length := len(id) // Use len for ASCII strings
			for n_characters := 1; n_characters <= id_length/2; n_characters++ {
				if id_length%n_characters != 0 {
					continue
				}
				pattern := id[:n_characters]
				comparison := strings.Repeat(pattern, id_length/n_characters)
				if id == comparison {
					id_int_value, err := strconv.Atoi(id)
					if err != nil {
						panic(err)
					}
					total += id_int_value
					break
				}
			}
		}
	}
	return total
}
