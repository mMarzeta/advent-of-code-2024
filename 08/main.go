package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
		parts := strings.Split(line, "")

		table = append(table, parts)
	}

	return table, nil
}

type Point struct {
	x int
	y int
	v string
}

func (p1 Point) Equal(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func calcAntinodes(p1, p2 Point, bounce bool, table *[][]string) []Point {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	antinodes := []Point{
		{x: p1.x + dx, y: p1.y + dy},
		{x: p2.x - dx, y: p2.y - dy},
	}

	if !bounce {
		return antinodes
	}

	a1, a2 := antinodes[0], antinodes[1]
	for inBound(a1, table) || inBound(a2, table) {
		a1 = Point{x: a1.x + dx, y: a1.y + dy}
		a2 = Point{x: a2.x - dx, y: a2.y - dy}

		if inBound(a1, table) {
			antinodes = append(antinodes, a1)
		}
		if inBound(a2, table) {
			antinodes = append(antinodes, a2)
		}
	}

	return antinodes
}

func getCords(table [][]string) map[string][]Point {
	cords := make(map[string][]Point)

	for y, row := range table {
		for x, val := range row {
			if val != "." {
				p := Point{x, y, val}
				cords[val] = append(cords[val], p)
			}
		}
	}

	return cords
}

func inBound(p Point, table *[][]string) bool {
	t := *table
	return !(p.x < 0 || p.x >= len(t[0]) || p.y < 0 || p.y >= len(t))
}

func placeAntinode(table *[][]string, a Point, key string) bool {
	t := *table
	if !inBound(a, &t) {
		return false
	}
	if t[a.y][a.x] == "." {
		t[a.y][a.x] = "#"
		return true
	}
	if t[a.y][a.x] != key && t[a.y][a.x] != "#" {
		return true
	}
	return false
}

func calcUniqueAntinodes(table [][]string, bounce bool) int {
	cords := getCords(table)
	counter := 0
	placedAntinodes := make(map[string]bool)

	for key, points := range cords {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]
				antinodes := calcAntinodes(p1, p2, bounce, &table)
				for _, a := range antinodes {
					if placeAntinode(&table, a, key) && !placedAntinodes[fmt.Sprintf("%d,%d", a.x, a.y)] {
						placedAntinodes[fmt.Sprintf("%d,%d", a.x, a.y)] = true
						counter++
					}

				}
			}
		}
	}

	if !bounce {
		return counter
	}

	counter = 0
	for _, row := range table {
		for _, e := range row {
			if e != "." {
				counter++
			}
		}
	}

	return counter
}

func main() {
	// Solution to https://adventofcode.com/2024/day/8
	fmt.Println("Day 8: Resonant Collinearity")

	table, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	table2 := deepCopySlice(table)

	// Part1
	start := time.Now()
	sum := calcUniqueAntinodes(table, false)
	duration := time.Since(start)
	fmt.Printf("Part1: %d \n(Execution Time: %s)\n\n", sum, formatDuration(duration))

	// Part2
	start = time.Now()
	sum = calcUniqueAntinodes(table2, true)
	duration = time.Since(start)
	fmt.Printf("Part2: %d \n(Execution Time: %s)\n\n", sum, formatDuration(duration))
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

func deepCopySlice(original [][]string) [][]string {
	copy := make([][]string, len(original))
	for i, row := range original {
		copy[i] = append([]string(nil), row...)
	}
	return copy
}
