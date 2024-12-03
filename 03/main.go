package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInput(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var parts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts = append(parts, line)
	}
	return parts, nil
}

func getUncorrupted(parts []string, includeDos bool) []string {
	uncorrupted := make([]string, 0)
	var re *regexp.Regexp

	if !includeDos {
		re = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	} else {

		re = regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(don't\(\))|(do\(\))`)
	}

	for _, part := range parts {
		uncorrupted = append(uncorrupted, re.FindAllString(part, -1)...)
	}

	return uncorrupted
}

func calculate(parts []string, includeDos bool) int {
	var res int
	re := regexp.MustCompile(`(\d{1,3})`)
	process := true

	for _, part := range parts {
		if includeDos {
			if part == "do()" {
				process = true
				continue
			}
			if part == "don't()" {
				process = false
				continue
			}

			if !process {
				continue
			}
		}

		matches := re.FindAllStringSubmatch(part, -1)

		l, err := strconv.Atoi(matches[0][0])
		if err != nil {
			log.Fatal("can't convert to int %w", err)
		}
		r, err := strconv.Atoi(matches[1][1])
		if err != nil {
			log.Fatal("can't convert to int %w", err)
		}

		res += l * r
	}

	return res
}

func main() {
	// Solution to https://adventofcode.com/2024/day/3
	fmt.Println("Day 3: Mull It Over")

	parts, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	// part 1
	uncorrupted := getUncorrupted(parts, false)
	res := calculate(uncorrupted, false)
	fmt.Printf("Part1: %d\n", res)

	// part 2
	uncorrupted = getUncorrupted(parts, true)
	res = calculate(uncorrupted, true)
	fmt.Printf("Part2: %d\n", res)
}
