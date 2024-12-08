package day6

import (
	"bufio"
	"fmt"
	"strings"
)

type Day6Grid struct {
	grid            [][]string
	direction       string
	currentLocation []int
}

var directions = map[string][]int{
	"left":  {0, -1},
	"right": {0, 1},
	"up":    {-1, 0},
	"down":  {1, 0},
}

func read(scanner *bufio.Scanner) [][]string {
	lines := make([][]string, 0)
	for scanner.Scan() {
		current := scanner.Text()
		lines = append(lines, strings.Split(current, ""))
	}

	return lines
}

func getNextDirection(current string) string {
	switch current {
	case "left":
		return "up"
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	}
	return ""
}

func (grid *Day6Grid) clear() {
	for i, row := range grid.grid {
		for j, value := range row {
			if value == "X" {
				grid.grid[i][j] = "."
			}
		}
	}
}

func (grid *Day6Grid) move() {
	direction := directions[grid.direction]
	next := []int{grid.currentLocation[0] + direction[0], grid.currentLocation[1] + direction[1]}

	if next[0] >= 0 && next[0] < len(grid.grid) && next[1] >= 0 && next[1] < len(grid.grid[0]) && grid.grid[next[0]][next[1]] == "#" {
		grid.direction = getNextDirection(grid.direction)
		return
	}

	grid.grid[grid.currentLocation[0]][grid.currentLocation[1]] = "X"
	grid.currentLocation = next
}

func findStart(input [][]string) []int {
	for i, row := range input {
		for j, value := range row {
			if value == "^" {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)
	grid := Day6Grid{grid: input, currentLocation: findStart(input), direction: "up"}
	for grid.currentLocation[0] >= 0 && grid.currentLocation[0] < len(input) && grid.currentLocation[1] >= 0 && grid.currentLocation[1] < len(input[0]) {
		grid.move()
	}

	total := 0
	for _, row := range grid.grid {
		for _, value := range row {
			if value == "X" {
				total++
			}
		}
	}

	fmt.Println(total)
}

func (grid *Day6Grid) move2(seen map[string]bool) bool {
	currentKey := fmt.Sprintf("%v-%v", grid.currentLocation, grid.direction)
	_, ok := seen[currentKey]
	if ok {
		return true
	}
	seen[currentKey] = true

	direction := directions[grid.direction]
	next := []int{grid.currentLocation[0] + direction[0], grid.currentLocation[1] + direction[1]}

	if next[0] >= 0 && next[0] < len(grid.grid) && next[1] >= 0 && next[1] < len(grid.grid[0]) && grid.grid[next[0]][next[1]] == "#" {
		grid.direction = getNextDirection(grid.direction)
		return false
	}

	grid.grid[grid.currentLocation[0]][grid.currentLocation[1]] = "X"
	grid.currentLocation = next
	return false
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)
	total := 0
	start := findStart(input)
	grid := Day6Grid{grid: input, currentLocation: start, direction: "up"}

	for i, row := range input {
		for j, value := range row {
			if value == "^" || value == "#" {
				continue
			}

			grid.currentLocation = start
			grid.direction = "up"

			grid.grid[i][j] = "#"

			seen := make(map[string]bool)

			valid := true
			for grid.currentLocation[0] >= 0 && grid.currentLocation[0] < len(grid.grid) && grid.currentLocation[1] >= 0 && grid.currentLocation[1] < len(grid.grid[0]) {
				if grid.move2(seen) {

					valid = false
					break
				}
			}

			grid.grid[i][j] = "."
			grid.clear()

			if !valid {
				total++
			}
		}
	}

	fmt.Println(total)
}
