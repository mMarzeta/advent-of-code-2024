package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var nums [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var row []int
		for _, number := range parts {
			num, err := strconv.Atoi(number)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", number)
			}
			row = append(row, num)
		}
		nums = append(nums, row)
	}

	return nums, nil
}

func isReportSafe(report []int) bool {
	minDiff := 1
	maxDiff := 3
	var increasing bool

	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		diff := int(math.Abs(float64(curr) - float64(next)))

		if (diff < minDiff) || (diff > maxDiff) {
			return false
		}

		if i == 0 {
			if curr == next {
				return false
			}
			increasing = curr < next
		}
		if (increasing && curr > next) || (!increasing && curr < next) {
			return false
		}
	}
	return true
}

func countSafeReports(reports [][]int, tryToFix bool) int {
	var safeCount int

	for _, report := range reports {
		isSafe := isReportSafe(report)
		if isSafe {
			safeCount++
		} else if tryToFix {
			for i := range report {
				portion := slices.Delete(slices.Clone(report), i, i+1)
				if isReportSafe(portion) {
					safeCount++
					break
				}
			}
		}
	}
	return safeCount
}

func main() {
	// Solution to https://adventofcode.com/2024/day/2

	reports, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	safeCount := countSafeReports(reports, false)
	fmt.Println(safeCount)

	// Part2
	safeCount = countSafeReports(reports, true)
	fmt.Println(safeCount)

}
