package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Advent of Code: 2024 - Day 01 - Part 1")
	fmt.Println("Processing example input...")
	leftList, rightList := readInput("01_1_ex.txt")
	distance := processDistance(leftList, rightList)
	fmt.Printf("Total distance between IDs: %d\n", distance)
	fmt.Println("Processing puzzle input...")
	leftList, rightList = readInput("01_1_input.txt")
	distance = processDistance(leftList, rightList)
	fmt.Printf("Total distance between IDs: %d\n", distance)
}

func partTwo() {
	fmt.Println("Advent of Code: 2024 - Day 01 - Part 2")
	fmt.Println("Processing example input...")
	leftList, rightList := readInput("01_1_ex.txt")
	similarity := processSimilarity(leftList, rightList)
	fmt.Printf("Total similarity between IDs: %d\n", similarity)
	fmt.Println("Processing puzzle input...")
	leftList, rightList = readInput("01_1_input.txt")
	similarity = processSimilarity(leftList, rightList)
	fmt.Printf("Total similarity between IDs: %d\n", similarity)
}

func processDistance(leftList []int, rightList []int) int {
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	var totalDistance int

	for idx := range leftList {
		dist := diffInt(leftList[idx], rightList[idx])
		totalDistance += dist
	}
	return totalDistance
}

func processSimilarity(leftList []int, rightList []int) int {
	var similarity int
	// Create a dictionary of values for each element
	rightCounts := make(map[int]int)

	for _, rightVal := range rightList {
		rightCounts[rightVal] += 1
	}

	for _, leftVal := range leftList {
		similarity += leftVal * rightCounts[leftVal]
	}

	return similarity
}

func readInput(fname string) ([]int, []int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		numbers := strings.Fields(ln)
		leftNum, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		rightNum, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return leftList, rightList
}

func diffInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
