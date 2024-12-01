package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var ls []int
	var rs []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[0])
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[1])
		}

		ls = append(ls, left)
		rs = append(rs, right)
	}

	return ls, rs, nil
}

func calcDistance(ls []int, rs []int) (int, error) {
	var err error
	if len(ls) != len(rs) {
		fmt.Println("slices are not equal")
		return 0, err
	}

	// edge case when list is empty
	if len(ls) == 1 {
		return 0, nil
	}

	sort.Ints(ls)
	sort.Ints(rs)
	var distance int
	for i := 0; i < len(ls); i++ {
		distance += int(math.Abs(float64(ls[i] - rs[i])))
	}

	return distance, nil
}

func countOccurencies(l []int) map[int]int {
	occurencies := make(map[int]int)
	for _, item := range l {
		occurencies[item]++
	}
	return occurencies
}

func calcSimilarity(ls []int, rs []int) int {
	occurencies := countOccurencies(rs)
	var similarity int

	for _, item := range ls {
		similarity += item * occurencies[item]
	}

	return similarity
}

func main() {
	ls, rs, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	distance, err := calcDistance(ls, rs)
	if err != nil {
		log.Fatalf("Error calculating distance: %v", err)
	}

	fmt.Println(distance)

	similarity := calcSimilarity(ls, rs)
	fmt.Println(similarity)
}
