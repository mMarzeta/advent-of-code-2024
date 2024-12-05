package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) (map[string]int, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var pageOrderingRule map[string]int = make(map[string]int)
	var pagesToProduce [][]int
	firstPart := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			firstPart = false
			continue
		}

		if firstPart {
			pageOrderingRule[line]++
		} else {
			var intList []int
			for _, part := range strings.Split(line, ",") {
				value, err := strconv.Atoi(part)
				if err == nil {
					intList = append(intList, value)
				}
			}
			pagesToProduce = append(pagesToProduce, intList)
		}
	}

	return pageOrderingRule, pagesToProduce, nil
}

func printQueue(pagesOrderingRule map[string]int, pagesToProduce [][]int) (int, [][]int) {
	counter := 0
	var incorrectPages [][]int

	for _, pages := range pagesToProduce {
		isCorrect := true
		for i := 0; i < len(pages)-1; i++ {
			curr, next := pages[i], pages[i+1]

			key := fmt.Sprintf("%d|%d", curr, next)
			_, ok := pagesOrderingRule[key]
			if !ok {
				isCorrect = false
				break
			}
		}

		if isCorrect {
			counter += pages[len(pages)/2]
		} else {
			incorrectPages = append(incorrectPages, pages)
		}
	}

	return counter, incorrectPages
}

func fixIncorrectPages(pagesOrderingRule map[string]int, incorrectPages [][]int) [][]int {
	var corrected [][]int

	for _, page := range incorrectPages {
		i, j := 0, 1
		for {
			if j == len(page) {
				i++
				j = i
			}
			if i == len(page)-1 {
				break
			}
			if i == j {
				j++
				continue
			}

			curr, next := page[i], page[j]

			comp := fmt.Sprintf("%d|%d", curr, next)
			_, ok := pagesOrderingRule[comp]
			if ok {
				j++
			} else {
				page[i] = next
				page[j] = curr
				j = i
			}
		}
		corrected = append(corrected, page)
	}
	return corrected
}

func main() {
	// Solution to https://adventofcode.com/2024/day/5
	fmt.Println("Day 5: Print Queue")

	pageOrderingRule, pagesToProduce, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	counter, incorrectPages := printQueue(pageOrderingRule, pagesToProduce)
	fmt.Printf("Part1: %d\n", counter)

	// Part2
	corrected := fixIncorrectPages(pageOrderingRule, incorrectPages)
	counter, incorrectPages = printQueue(pageOrderingRule, corrected)
	fmt.Printf("Part2: %d\n", counter)
}
