package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var table []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		for _, p := range parts {
			num, err := strconv.Atoi(p)
			if err != nil {
				num = -1
			}
			table = append(table, num)
		}

	}

	return table, nil
}

func digitCount(n int) int {
	return int(math.Floor(math.Log10(float64(n))) + 1)
}

func transformStone(stone int) (int, int, int) {
	if stone == 0 {
		return 1, 0, 0
	}

	digits := digitCount(stone)
	if digits%2 == 0 {
		splitPower := int(math.Pow10(digits / 2))

		leftHalf := stone / splitPower
		rightHalf := stone % splitPower

		return leftHalf, rightHalf, 1
	}

	return stone * 2024, 0, 2
}

func blink(cache map[int]int) map[int]int {
	tmpCache := make(map[int]int, 0)
	for stone, count := range cache {
		//		fmt.Printf("stone: %d, count: %d\n", stone, count)
		if count <= 0 {
			continue
		}
		cache[stone] = 0
		l, r, t := transformStone(stone)

		if t == 0 {
			//			fmt.Println("0->1")
			tmpCache[l] += count
		} else if t == 1 {
			//= count			fmt.Println("split")
			if count == 1 {
				tmpCache[l]++
				tmpCache[r]++

			} else {
				tmpCache[l] += count
				tmpCache[r] += count

			}
		} else if t == 2 {
			//			fmt.Println("*2024")
			tmpCache[l] += count
		}
	}

	for k, v := range tmpCache {
		cache[k] = v
	}
	return cache
}

func printStones(cache map[int]int) int {
	sum := 0
	for k, v := range cache {
		if v != 0 {
			fmt.Printf("%d:%d, ", k, v)
		}

		sum += v
	}
	fmt.Printf("\nsum: %d\n", sum)
	return sum
}

func solveStoneBlinking(initialStones []int, blinkCount int) int {
	stones := initialStones
	cache := make(map[int]int, 0)

	for _, stone := range stones {
		cache[stone]++
	}

	for i := 0; i < blinkCount; i++ {
		cache = blink(cache)
	}

	sum := 0
	for _, val := range cache {
		sum += val
	}
	return sum
}

func main() {
	// Solution to https://adventofcode.com/2024/day/11
	fmt.Println("Day 11: Plutonian Pebbles")

	arr, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	start := time.Now()
	finalStones := solveStoneBlinking(arr, 25)
	duration := time.Since(start)
	fmt.Printf("Part1: %d \n(Execution Time: %s)\n\n", finalStones, formatDuration(duration))

	// Part2
	start = time.Now()
	finalStones = solveStoneBlinking(arr, 75)
	duration = time.Since(start)
	fmt.Printf("Part2: %d \n(Execution Time: %s)\n\n", finalStones, formatDuration(duration))
}

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	}
	if d < time.Millisecond {
		return fmt.Sprintf("%.3f µs", float64(d.Nanoseconds())/1000)
	}
	if d < time.Second {
		return fmt.Sprintf("%.3f ms", float64(d.Microseconds())/1000)
	}
	return d.Round(time.Millisecond).String()
}
