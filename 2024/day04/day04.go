package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Advent of Code: 2024 - Day 04 - Part 1")
	wordSearch := readInput("example.txt")
	xmasCount := processWordSearch(wordSearch)
	fmt.Printf("  Total 'XMAS' occurrences: %d\n", xmasCount)
	wordSearch = readInput("input.txt")
	xmasCount = processWordSearch(wordSearch)
	fmt.Printf("  Total 'XMAS' occurrences: %d\n", xmasCount)
}

func partTwo() {
	fmt.Println("Advent of Code: 2024 - Day 04 - Part 2")
	wordSearch := readInput("example.txt")
	xMasCount := processCrossWordSearch(wordSearch)
	fmt.Printf("  Total 'X-MAS' occurrences: %d\n", xMasCount)
	wordSearch = readInput("input.txt")
	xMasCount = processCrossWordSearch(wordSearch)
	fmt.Printf("  Total 'X-MAS' occurrences: %d\n", xMasCount)
}

func processWordSearch(runeGrid [][]rune) int {
	var sum int

	for i, runeRow := range runeGrid {
		for j := range runeRow {
			for _, d := range Directions.List() {
				sum += validateWordLine(runeGrid, []rune("XMAS"), i, j, d)
			}
		}
	}

	return sum
}

func validateWordLine(crossword [][]rune, word []rune, rowIdx int, colIdx int, direction *Direction) int {
	if (rowIdx < 0) || (rowIdx >= len(crossword)) || (colIdx < 0) || (colIdx >= len(crossword[0])) {
		return 0
	}
	if crossword[rowIdx][colIdx] != word[0] {
		return 0
	}
	if len(word) == 1 {
		return 1
	}
	return validateWordLine(crossword, word[1:], rowIdx+direction.v, colIdx+direction.h, direction)
}

func processCrossWordSearch(runeGrid [][]rune) int {
	var sum int
	directions := []*Direction{Directions.North, Directions.East, Directions.South, Directions.West}

	for i, runeRow := range runeGrid {
		for j := range runeRow {
			for _, d := range directions {
				sum += validateCrossWord(runeGrid, []rune("MAS"), i, j, d)
			}
		}
	}

	return sum
}

func validateCrossWord(crossword [][]rune, word []rune, rowIdx int, colIdx int, direction *Direction) int {
	if (rowIdx > len(crossword)-len(word)) || (colIdx > len(crossword[0])-len(word)) {
		return 0
	}
	var matchOrder [5]rune
	matchOrder[2] = word[1]
	switch direction {
	case Directions.North:
		matchOrder[0] = word[0]
		matchOrder[1] = word[0]
		matchOrder[3] = word[2]
		matchOrder[4] = word[2]
	case Directions.East:
		matchOrder[0] = word[2]
		matchOrder[1] = word[0]
		matchOrder[3] = word[2]
		matchOrder[4] = word[0]
	case Directions.South:
		matchOrder[0] = word[2]
		matchOrder[1] = word[2]
		matchOrder[3] = word[0]
		matchOrder[4] = word[0]
	case Directions.West:
		matchOrder[0] = word[0]
		matchOrder[1] = word[2]
		matchOrder[3] = word[0]
		matchOrder[4] = word[2]
	default:
		return 0
	}
	if (crossword[rowIdx][colIdx] != matchOrder[0]) || (crossword[rowIdx][colIdx+2] != matchOrder[1]) || (crossword[rowIdx+1][colIdx+1] != matchOrder[2]) || (crossword[rowIdx+2][colIdx] != matchOrder[3]) || (crossword[rowIdx+2][colIdx+2] != matchOrder[4]) {
		return 0
	}
	return 1
}

var Directions = newDirectionRegistry()

type Direction struct {
	h int
	v int
}

func newDirectionRegistry() *directionRegistry {

	north := &Direction{0, -1}
	northEast := &Direction{1, -1}
	east := &Direction{1, 0}
	southEast := &Direction{1, 1}
	south := &Direction{0, 1}
	southWest := &Direction{-1, 1}
	west := &Direction{-1, 0}
	northWest := &Direction{-1, -1}

	return &directionRegistry{
		North:      north,
		NorthEast:  northEast,
		East:       east,
		SouthEast:  southEast,
		South:      south,
		SouthWest:  southWest,
		West:       west,
		NorthWest:  northWest,
		directions: []*Direction{north, northEast, east, southEast, south, southWest, west, northWest},
	}
}

type directionRegistry struct {
	North      *Direction
	NorthEast  *Direction
	East       *Direction
	SouthEast  *Direction
	South      *Direction
	SouthWest  *Direction
	West       *Direction
	NorthWest  *Direction
	directions []*Direction
}

func (d *directionRegistry) List() []*Direction {
	return d.directions
}

func readInput(fname string) [][]rune {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var runeGrid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		runeArray := []rune(str)
		runeGrid = append(runeGrid, runeArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return runeGrid
}
