package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var table [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		table = append(table, row)
	}

	return table, nil
}

func getStartingPosition(table [][]string) (int, int, error) {
	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			if table[y][x] == "^" {
				return x, y, nil
			}
		}
	}

	return -1, -1, fmt.Errorf("can't find starting point")
}

type Direction int

const (
	south Direction = iota
	east
	north
	west
)

func rotate90(direction Direction) Direction {
	if direction == west {
		return south
	}

	return direction + 1
}

func makeMove(direction Direction, x int, y int) (int, int) {
	switch direction {
	case south:
		return x, y - 1
	case north:
		return x, y + 1
	case east:
		return x + 1, y
	case west:
		return x - 1, y
	}
	return x, y
}
func getTrailTable(table [][]string, startX int, startY int, d Direction) ([][]string, bool) {
	x, y := startX, startY
	visitedStates := make(map[string]int)

	for {
		key := fmt.Sprintf("%d|%d|%d", x, y, d)

		// detect cycle
		if count, exists := visitedStates[key]; exists {
			if count >= 2 {
				return table, true
			}
			visitedStates[key]++
		} else {
			visitedStates[key] = 1
		}

		nextX, nextY := makeMove(d, x, y)

		// boundaries
		if nextX < 0 || nextX >= len(table[0]) || nextY < 0 || nextY >= len(table) {
			table[y][x] = "x"
			break
		}

		// collision detection
		if table[nextY][nextX] == "#" || table[nextY][nextX] == "O" {
			d = rotate90(d)
			continue
		}

		// Mark trail and move
		table[y][x] = "x"
		x, y = nextX, nextY
	}

	table[startY][startX] = "^"
	return table, false
}

func makeCycles(table [][]string, sx int, sy int, d Direction) int {
	obstacles := 0
	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[0]); x++ {
			tmpTable := make([][]string, len(table))
			for i := range table {
				tmpTable[i] = make([]string, len(table[i]))
				copy(tmpTable[i], table[i])
			}

			curr := table[y][x]
			if curr != "^" {
				tmpTable[y][x] = "O"
				_, cycle := getTrailTable(tmpTable, sx, sy, d)
				if cycle {
					obstacles++
				}
			}

		}
	}
	return obstacles
}

func calcTrail(table [][]string) int {
	trail := 0
	for _, row := range table {
		for _, v := range row {
			if v == "x" {
				trail++
			}
		}
	}
	return trail + 1 // to cound also starting point
}

func main() {
	// Solution to https://adventofcode.com/2024/day/6
	fmt.Println("Day 6: Guard Gallivant")

	table, err := readInput("input_short.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	// Part1
	sx, sy, err := getStartingPosition(table)
	trailedTable, _ := getTrailTable(table, sx, sy, south)
	trails := calcTrail(trailedTable)
	fmt.Printf("Part1: %d\n", trails)

	// Part2
	obstructions := makeCycles(table, sx, sy, south)
	fmt.Printf("Part2: %d\n", obstructions)
}
