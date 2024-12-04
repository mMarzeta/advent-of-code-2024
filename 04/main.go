package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var letters [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []string
		for _, letter := range line {
			row = append(row, string(letter))
		}
		letters = append(letters, row)
	}

	return letters, nil
}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func findHorizontally(puzzle [][]string, word string, revWord string) int {
	var counter int

	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[0])-len(word)+1; col++ {
			tmp := strings.Join(puzzle[row][col:col+len(word)], "")
			if tmp == word || tmp == revWord {
				counter++
			}
		}
	}
	return counter
}

func findDiagonally(puzzle [][]string, word string, revWord string) int {
	var counter int

	// Diagonal \
	//           \
	for row := 0; row < len(puzzle)-len(word)+1; row++ {
		for col := 0; col < len(puzzle[0])-len(word)+1; col++ {
			diag := ""
			for i := 0; i < len(word); i++ {
				diag += puzzle[row+i][col+i]
			}
			if diag == word || diag == revWord {
				counter++
			}
		}
	}

	// Diagonal   /
	//           /
	for row := len(word) - 1; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[0])-len(word)+1; col++ {
			diag := ""
			for i := 0; i < len(word); i++ {
				diag += puzzle[row-i][col+i]
			}
			if diag == word || diag == revWord {
				counter++
			}
		}
	}

	return counter
}

func findVertically(puzzle [][]string, word string, revWord string) int {
	var counter int

	for row := 0; row < len(puzzle)-len(word)+1; row++ {
		for col := 0; col < len(puzzle[0]); col++ {
			vertice := ""
			for i := 0; i < len(word); i++ {
				vertice += puzzle[row+i][col]
			}
			if vertice == word || vertice == revWord {
				counter++
			}
		}
	}
	return counter
}

func findWords(puzzle [][]string, word string) int {
	var counter int
	revWord := ReverseString(word)

	counter += findHorizontally(puzzle, word, revWord)
	counter += findVertically(puzzle, word, revWord)
	counter += findDiagonally(puzzle, word, revWord)

	return counter
}

func findCrosses(puzzle [][]string, word string) int {
	revWord := ReverseString(word)
	halfWordLen := len(word) / 2
	var counter int

	for row := halfWordLen; row < len(puzzle)-halfWordLen; row++ {
		for col := halfWordLen; col < len(puzzle[0])-halfWordLen; col++ {
			diag1, diag2 := "", ""
			for i := -1; i < len(word)-1; i++ {
				diag1 += puzzle[row-i][col+i]
				diag2 += puzzle[row+i][col+i]
			}
			if (diag1 == word || diag1 == revWord) && (diag2 == word || diag2 == revWord) {
				counter++
			}
		}
	}
	return counter
}

func main() {
	// Solution to https://adventofcode.com/2024/day/4
	fmt.Println("Day 4: Ceres Search")

	puzzle, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	counter := findWords(puzzle, "XMAS")
	fmt.Printf("Part1: %d\n", counter)

	// Part2
	counter = findCrosses(puzzle, "MAS")
	fmt.Printf("Part2: %d\n", counter)
}
