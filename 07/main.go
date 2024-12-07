package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) ([]int, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var results []int
	var numbers [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid num of parts: %v", parts)
		}

		res, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number part1: %s", parts[0])
		}
		results = append(results, res)

		var intSlice []int
		for _, num := range strings.Split(parts[1], " ") {
			if num == "" {
				continue
			}

			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number part2: %s", num)
			}
			intSlice = append(intSlice, val)
		}
		numbers = append(numbers, intSlice)
	}

	return results, numbers, nil
}
func generatePermutations(n int, ops []string) []string {
	var result []string

	var backtrack func(current string, depth int)
	backtrack = func(current string, depth int) {
		if depth == n-1 {
			result = append(result, current)
			return
		}
		for _, op := range ops {
			backtrack(current+op, depth+1)
		}
	}

	backtrack("", 0)

	return result
}

func findOperators(results []int, numbers [][]int, ops []string) (int, error) {
	if len(results) != len(numbers) {
		return -1, fmt.Errorf("not equal length")
	}

	sum := 0
	opPerms := make(map[int][]string)

	for i := 0; i < len(results); i++ {
		res := results[i]
		nums := numbers[i]
		perms, ok := opPerms[len(nums)]
		if !ok {
			perms = generatePermutations(len(nums), ops)
			opPerms[len(nums)] = perms
		}

		for _, ops := range perms {
			tmpRes := 0
			tmp := 0

			for j := 0; j < len(ops); j++ {
				l, r := 0, 0
				if j == 0 {
					l, r = nums[j], nums[j+1]
				} else {
					l, r = tmpRes, nums[j+1]
				}
				op := string(ops[j])

				switch op {
				case "+":
					tmp = l + r
					tmpRes = tmp
				case "*":
					tmp = l * r
					tmpRes = tmp
				case "|":
					var err error
					tmp, err = strconv.Atoi(fmt.Sprintf("%d%d", l, r))
					if err != nil {
						log.Panicf("cant concatenate ints %v", err)
					}

					tmpRes = tmp
				}
			}
			if tmpRes == res {
				sum += res
				break
			}

		}
	}

	return sum, nil
}
func main() {
	// Solution to https://adventofcode.com/2024/day/7
	fmt.Println("Day 7: Bridge Repair")

	results, numbers, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	ops := []string{"+", "*"}
	sum, _ := findOperators(results, numbers, ops)
	fmt.Printf("Part1: %d\n", sum)

	// Part2
	ops = []string{"+", "*", "|"}
	sum, _ = findOperators(results, numbers, ops)
	fmt.Printf("Part2: %d\n", sum)
}
