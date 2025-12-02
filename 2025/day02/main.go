// day02/main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

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

	var expanded_ids []string
	for i := id_range_start; i <= id_range_end; i++ {
		expanded_ids = append(expanded_ids, fmt.Sprint(i))
	}
	return expanded_ids
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string) interface{} {
	total := 0
	line := lines[0]
	ranges := strings.Split(line, ",")
	for _, range_str := range ranges {
		range_ids := expandRange(range_str)
		for _, id := range range_ids {
			id_length := utf8.RuneCountInString(id)
			if id_length%2 == 0 {
				first_half := id[:id_length/2]
				last_half := id[id_length/2:]
				if first_half == last_half {
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
	ranges := strings.Split(line, ",")
	for _, range_str := range ranges {
		range_ids := expandRange(range_str)
		for _, id := range range_ids {
			id_length := utf8.RuneCountInString(id)
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
