package day8

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
)

type Coord struct {
	row int
	col int
}

type Antennas map[string][]Coord

func read(scanner *bufio.Scanner) [][]string {
	lines := make([][]string, 0)
	for scanner.Scan() {
		current := scanner.Text()
		lines = append(lines, strings.Split(current, ""))
	}

	return lines
}

func getAntennas(grid [][]string) Antennas {
	antennas := make(map[string][]Coord)

	for i, row := range grid {
		for j, value := range row {
			if value != "." {
				_, ok := antennas[value]
				if !ok {
					antennas[value] = []Coord{{row: i, col: j}}
				} else {
					antennas[value] = append(antennas[value], Coord{row: i, col: j})
				}
			}
		}
	}

	return antennas
}

func markAntinodes(grid [][]string, key string, locations []Coord) {
	for i, first := range locations {
		for _, second := range locations[i+1:] {
			rowDiff := second.row - first.row
			colDiff := second.col - first.col

			firstAntinodeRow := first.row - rowDiff
			firstAntinodeCol := first.col - colDiff
			secondAntinodeRow := second.row + rowDiff
			secondAntinodeCol := second.col + colDiff

			if firstAntinodeRow >= 0 && firstAntinodeRow < len(grid) && firstAntinodeCol >= 0 && firstAntinodeCol < len(grid[0]) {
				grid[firstAntinodeRow][firstAntinodeCol] = "#"
			}

			if secondAntinodeRow >= 0 && secondAntinodeRow < len(grid) && secondAntinodeCol >= 0 && secondAntinodeCol < len(grid[0]) {
				grid[secondAntinodeRow][secondAntinodeCol] = "#"
			}
		}
	}
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)
	antennas := getAntennas(input)
	var wg sync.WaitGroup

	for key, value := range antennas {
		wg.Add(1)

		go func(key string, value []Coord) {
			defer wg.Done()
			markAntinodes(input, key, value)
		}(key, value)
	}

	wg.Wait()

	total := 0
	for _, row := range input {
		for _, value := range row {
			if value == "#" {
				total++
			}
		}
	}
	fmt.Println(total)
}

func markInDirection(grid [][]string, initial, spacing Coord) {
	current := initial
	next := Coord{row: initial.row + spacing.row, col: initial.col + spacing.col}
	for next.row >= 0 && next.row < len(grid) && next.col >= 0 && next.col < len(grid[0]) {
		grid[next.row][next.col] = "#"
		current = next
		next = Coord{row: current.row + spacing.row, col: current.col + spacing.col}
	}
}

func markAntinodes2(grid [][]string, key string, locations []Coord) {
	for i, first := range locations {
		for _, second := range locations[i+1:] {
			rowDiff := second.row - first.row
			colDiff := second.col - first.col

			var wg sync.WaitGroup

			wg.Add(2)
			go func(grid [][]string, first, diff Coord) {
				defer wg.Done()
				markInDirection(grid, first, diff)
			}(grid, first, Coord{row: -rowDiff, col: -colDiff})

			go func(grid [][]string, first, diff Coord) {
				defer wg.Done()
				markInDirection(grid, first, diff)
			}(grid, first, Coord{row: rowDiff, col: colDiff})

			wg.Wait()
		}
	}
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)
	antennas := getAntennas(input)
	var wg sync.WaitGroup

	for key, value := range antennas {
		wg.Add(1)

		go func(key string, value []Coord) {
			defer wg.Done()
			markAntinodes2(input, key, value)
		}(key, value)
	}

	wg.Wait()

	total := 0
	for _, row := range input {
		for _, value := range row {
			if value != "." {
				total++
			}
		}
	}
	fmt.Println(total)
}
