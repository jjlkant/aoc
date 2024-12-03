package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Advent of Code: 2024 - Day 03 - Part 1")
	memory := readInput("example.txt")
	mulTotal := processMemory(memory, false)
	fmt.Printf("  Total of multiplications: %d\n", mulTotal)
	memory = readInput("input.txt")
	mulTotal = processMemory(memory, false)
	fmt.Printf("  Total of multiplications: %d\n", mulTotal)
}

func partTwo() {
	fmt.Println("Advent of Code: 2024 - Day 03 - Part 2")
	memory := readInput("example_2.txt")
	mulTotal := processMemory(memory, true)
	fmt.Printf("  Total of multiplications: %d\n", mulTotal)
	memory = readInput("input.txt")
	mulTotal = processMemory(memory, true)
	fmt.Printf("  Total of multiplications: %d\n", mulTotal)
}

func processMemory(memory []string, conditional bool) int {
	var sum int
	var enabled bool = true

	for _, part := range memory {
		r := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
		matches := r.FindAllStringSubmatch(part, -1)
		for _, v := range matches {
			switch v[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if !conditional || (conditional && enabled) {
					sum += parseMul(v[0])
				}
			}
		}
	}

	return sum
}

func parseMul(mul string) int {
	parts := strings.FieldsFunc(mul, splitMul)
	leftNum, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	rightNum, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(err)
	}
	return leftNum * rightNum
}

func splitMul(r rune) bool {
	return r == '(' || r == ',' || r == ')'
}

func readInput(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var strs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		strs = append(strs, str)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strs
}
