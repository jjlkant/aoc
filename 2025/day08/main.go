package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jjlkant/aoc/2025/utils"
)

type Coordinate struct {
	x int
	y int
	z int
}

type JunctionBoxPair struct {
	distance float64
	from     Coordinate
	to       Coordinate
}

func main() {
	lines, err := utils.ReadInputLines()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully read %d lines for Day 1.\n", len(lines))

	// --- PART 1 ---
	start := time.Now()
	resultPart1 := solvePart1(lines, 1000)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 Solution: %v (took %s)\n", resultPart1, elapsed)

	// --- PART 2 ---
	start = time.Now()
	resultPart2 := solvePart2(lines)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 Solution: %v (took %s)\n", resultPart2, elapsed)
}

// Calculate the Euclidean distance between two coordinates
func distance(from, to Coordinate) float64 {
	return math.Sqrt(math.Pow(float64(from.x-to.x), 2) +
		math.Pow(float64(from.y-to.y), 2) +
		math.Pow(float64(from.z-to.z), 2))
}

// Parse input lines into a slice of Coordinates
func parseCoordinatesAndCircuits(lines []string) ([]Coordinate, []map[Coordinate]struct{}) {
	coordinates := make([]Coordinate, len(lines))
	circuits := make([]map[Coordinate]struct{}, len(coordinates))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		coordinates[i] = Coordinate{x, y, z}
		circuits[i] = map[Coordinate]struct{}{coordinates[i]: {}}
	}
	return coordinates, circuits
}

// Generate all pairwise distances between coordinates
func calculateDistances(coordinates []Coordinate) []JunctionBoxPair {
	var distances []JunctionBoxPair
	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {
			distances = append(distances, JunctionBoxPair{
				distance: distance(coordinates[i], coordinates[j]),
				from:     coordinates[i],
				to:       coordinates[j],
			})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
	return distances
}

func processPairs(pairs []JunctionBoxPair, circuits []map[Coordinate]struct{}, maxIterations int) ([]map[Coordinate]struct{}, JunctionBoxPair) {
	// Process distances to join circuits
	nonEmptyCircuits := make([]map[Coordinate]struct{}, 0)
	var linkingPair JunctionBoxPair
	for i := 0; i < maxIterations && i < len(pairs); i++ {
		from := pairs[i].from
		to := pairs[i].to

		var fromCircuit, toCircuit map[Coordinate]struct{}
		fromIndex, toIndex := -1, -1

		// Find the circuits containing `from` and `to`
		for j, circuit := range circuits {
			if _, exists := circuit[from]; exists {
				fromCircuit = circuit
				fromIndex = j
			}
			if _, exists := circuit[to]; exists {
				toCircuit = circuit
				toIndex = j
			}
			// Break early if both circuits are found
			if fromCircuit != nil && toCircuit != nil {
				break
			}
		}

		// Merge circuits if `from` and `to` are in different circuits
		if fromIndex != toIndex {
			for coord := range toCircuit {
				fromCircuit[coord] = struct{}{}
			}
			// Remove the merged circuit
			circuits[toIndex] = nil
		}

		// Filter out empty circuits
		nonEmptyCircuits = make([]map[Coordinate]struct{}, 0)
		for _, circuit := range circuits {
			if len(circuit) > 0 {
				nonEmptyCircuits = append(nonEmptyCircuits, circuit)
			}
		}

		// If only one circuit remains, store the first connection and break
		if len(nonEmptyCircuits) == 1 {
			linkingPair = pairs[i]
			break
		}
	}

	return nonEmptyCircuits, linkingPair
}

// This function holds your Part 1 logic and is what you will test.
func solvePart1(lines []string, iterations int) any {
	coordinates, circuits := parseCoordinatesAndCircuits(lines)
	distances := calculateDistances(coordinates)

	nonEmptyCircuits, _ := processPairs(distances, circuits, iterations)

	sort.Slice(nonEmptyCircuits, func(i, j int) bool {
		return len(nonEmptyCircuits[i]) > len(nonEmptyCircuits[j])
	})

	return len(nonEmptyCircuits[0]) * len(nonEmptyCircuits[1]) * len(nonEmptyCircuits[2])
}

// This function holds your Part 2 logic.
func solvePart2(lines []string) any {
	coordinates, circuits := parseCoordinatesAndCircuits(lines)
	distances := calculateDistances(coordinates)

	_, linkingPair := processPairs(distances, circuits, len(distances))

	return linkingPair.from.x * linkingPair.to.x
}
