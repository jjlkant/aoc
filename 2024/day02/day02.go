package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Advent of Code: 2024 - Day 02 - Part 1")
	reports := readInput("02_1_ex.txt")
	numSafe := processReports(reports)
	fmt.Printf("  Total distance between IDs: %d\n", numSafe)
	reports = readInput("02_1_input.txt")
	numSafe = processReports(reports)
	fmt.Printf("  Total distance between IDs: %d\n", numSafe)
}

func partTwo() {
	fmt.Println("Advent of Code: 2024 - Day 02 - Part 2")
	reports := readInput("02_1_ex.txt")
	numSafe := processReportsDampened(reports)
	fmt.Printf("  Total distance between IDs: %d\n", numSafe)
	reports = readInput("02_1_input.txt")
	numSafe = processReportsDampened(reports)
	fmt.Printf("  Total distance between IDs: %d\n", numSafe)
}

func processReports(reports [][]int) int {
	var numSafe int
	for _, report := range reports {
		if isReportSafe(report) {
			numSafe += 1
		}
	}
	return numSafe
}

func removeIndex(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

func processReportsDampened(reports [][]int) int {
	var numSafe int
	for _, report := range reports {
		safe := isReportSafe(report)
		if safe {
			numSafe += 1
			continue
		}
		for idx, _ := range report {
			if isReportSafe(removeIndex(report, idx)) {
				numSafe += 1
				break
			}
		}
	}
	return numSafe
}

func isReportSafe(report []int) bool {
	increasing := report[1]-report[0] > 0

	for idx, level := range report[1:] {
		diff := level - report[idx]
		if (diff == 0) || (increasing && (diff <= 0)) || (!increasing && (diff >= 0)) {
			return false
		}
		if (diff < -3) || (diff > 3) {
			return false
		}
	}
	return true
}

func readInput(fname string) [][]int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportStr := strings.Fields(scanner.Text())
		var report []int
		for _, v := range reportStr {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, val)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reports
}

func diffInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
