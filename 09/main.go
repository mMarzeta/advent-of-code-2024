package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		for _, digit := range line {
			num := int(digit - '0')
			table = append(table, num)
		}
	}

	return table, nil
}

func printLine(line []int) {
	for _, e := range line {
		if e == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", e)
		}
	}
	fmt.Println()
}

func compact(line []int) []int {
	blocks := make([]int, 0)
	id := 0

	for i, e := range line {
		if i%2 == 0 {
			for j := 0; j < e; j++ {
				blocks = append(blocks, id)
			}
			id++
		} else {
			for j := 0; j < e; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	i, j := 0, len(blocks)-1
	var l, r int

	for i <= j {
		if blocks[i] == -1 {
			l = i
		} else {
			i++
			continue
		}

		if blocks[j] != -1 {
			r = j
		} else {
			j--
			continue
		}

		blocks[l], blocks[r] = blocks[r], blocks[l]
	}

	return blocks
}

func findEmptyPlace(line []int, start int) (int, int, bool) {
	size := 0
	inEmpty := false
	for i := start; i < len(line); i++ {
		if line[i] == -1 {
			if !inEmpty {
				start = i
				inEmpty = true
				size++
			} else {
				size++
			}
		}
		if line[i] != -1 && inEmpty {
			break
		}
	}

	return start, size, !inEmpty
}

func findDataBlock(line []int, sStart int) (int, int) {
	start, size := 0, 0
	id := -1
	for i := sStart; i >= 0; i-- {
		if line[i] == -1 {
			continue
		}
		if id == -1 {
			id = line[i]
			start = i
		}
		if line[i] != id {
			break
		}
		if id == line[i] {
			size++
		}
	}

	start = start - size + 1
	return start, size
}

func swap(blocks *[]int, eS, bS, bL int) {
	b := *blocks
	for i := 0; i < bL; i++ {
		b[eS+i] = b[bS+i]
		b[bS+i] = -1
	}
}

func compactWholeFiles(line []int) []int {
	blocks := make([]int, 0)
	id := 0

	for i, e := range line {
		if i%2 == 0 {
			for j := 0; j < e; j++ {
				blocks = append(blocks, id)
			}
			id++
		} else {
			for j := 0; j < e; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	i, j := 0, len(blocks)-1

	for {
		bStart, bSize := findDataBlock(blocks, j)
		j = bStart - 1
		if j <= 0 {
			break
		}

		for {
			eStart, eSize, end := findEmptyPlace(blocks, i)
			if end {
				i = 0
				break
			}
			i = eStart + eSize
			if bSize <= eSize && eStart < bStart {
				swap(&blocks, eStart, bStart, bSize)
				i = 0
				break
			}
		}
	}

	return blocks
}
func calcCheckSum(compacted []int) int {
	checkSum := 0
	for i, e := range compacted {
		if e != -1 {
			checkSum += i * e
		}
	}
	return checkSum
}

func main() {
	// Solution to https://adventofcode.com/2024/day/9
	fmt.Println("Day 9: Disk Fragmenter")

	line, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	start := time.Now()
	compact := compact(line)
	checkSum := calcCheckSum(compact)
	duration := time.Since(start)
	fmt.Printf("Part1: %d \n(Execution Time: %s)\n\n", checkSum, formatDuration(duration))

	// Part2
	start = time.Now()
	compact = compactWholeFiles(line)
	checkSum = calcCheckSum(compact)
	duration = time.Since(start)
	fmt.Printf("Part2: %d \n(Execution Time: %s)\n\n", checkSum, formatDuration(duration))
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
