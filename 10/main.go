package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	var table [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		row := make([]int, 0)
		for _, p := range parts {
			num, err := strconv.Atoi(p)
			if err != nil {
				num = -1
			}
			row = append(row, num)
		}

		table = append(table, row)
	}

	return table, nil
}

type Point struct {
	x, y, v int
}

func findStartingPoints(table [][]int) []Point {
	sp := make([]Point, 0)
	for y, row := range table {
		for x, num := range row {
			if num == 0 {
				sp = append(sp, Point{x, y, num})
			}
		}
	}
	return sp
}

func findValidMoves(t [][]int, cp Point) []Point {
	vm := make([]Point, 0)

	// up
	if cp.y != 0 {
		pu := Point{cp.x, cp.y - 1, t[cp.y-1][cp.x]}
		if pu.v-cp.v == 1 {
			vm = append(vm, pu)
		}
	}
	// down
	if cp.y != len(t)-1 {
		pd := Point{cp.x, cp.y + 1, t[cp.y+1][cp.x]}
		if pd.v-cp.v == 1 {
			vm = append(vm, pd)
		}
	}
	// left
	if cp.x != 0 {
		pl := Point{cp.x - 1, cp.y, t[cp.y][cp.x-1]}
		if pl.v-cp.v == 1 {
			vm = append(vm, pl)
		}
	}
	// right
	if cp.x != len(t[cp.y])-1 {
		pr := Point{cp.x + 1, cp.y, t[cp.y][cp.x+1]}
		if pr.v-cp.v == 1 {
			vm = append(vm, pr)
		}
	}
	return vm
}

func hike(table [][]int, currPoint Point, visitedTops map[Point]bool) int {
	if currPoint.v == 9 {
		_, ok := visitedTops[currPoint]
		if ok {
			return 0
		} else {
			visitedTops[currPoint] = true
			return 1
		}
	}

	validMoves := findValidMoves(table, currPoint)
	if len(validMoves) == 0 {
		return 0
	}

	sum := 0
	for _, move := range validMoves {
		sum += hike(table, move, visitedTops)
	}
	return sum
}

func hikeRating(table [][]int, currPoint Point, visitedTops map[Point]bool) int {
	if currPoint.v == 9 {
		return 1
	}

	validMoves := findValidMoves(table, currPoint)
	if len(validMoves) == 0 {
		return 0
	}

	visitedTops[currPoint] = true
	totalTrails := 0

	for _, move := range validMoves {
		if !visitedTops[move] {
			totalTrails += hikeRating(table, move, visitedTops)
		}
	}

	visitedTops[currPoint] = false
	return totalTrails
}

func calcRating(table [][]int) int {
	startingPoints := findStartingPoints(table)

	totalRating := 0
	for _, sp := range startingPoints {
		visitedTops := make(map[Point]bool)
		trails := hikeRating(table, sp, visitedTops)
		totalRating += trails
	}

	return totalRating
}

func printTable(table [][]int) {
	for _, row := range table {
		for _, num := range row {
			if num == -1 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// Solution to https://adventofcode.com/2024/day/10
	fmt.Println("Day 10: Hoof It")

	table, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part1
	start := time.Now()
	startingPoints := findStartingPoints(table)

	routes := 0
	for _, sp := range startingPoints {
		visitedTops := make(map[Point]bool, 0)
		tmp := hike(table, sp, visitedTops)
		routes += tmp
	}
	duration := time.Since(start)
	fmt.Printf("Part1: %d \n(Execution Time: %s)\n\n", routes, formatDuration(duration))

	// Part2
	start = time.Now()
	rating := calcRating(table)
	duration = time.Since(start)
	fmt.Printf("Part2: %d \n(Execution Time: %s)\n\n", rating, formatDuration(duration))
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
