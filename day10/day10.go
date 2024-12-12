package day10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type TrailMap struct {
	grid   [][]int
	counts [][]int
}

func read(scanner *bufio.Scanner) [][]int {
	lines := make([][]int, 0)
	for scanner.Scan() {
		current := scanner.Text()
		split := strings.Split(current, "")
		line := make([]int, len(split))
		for i, value := range split {
			converted, _ := strconv.Atoi(value)
			line[i] = converted
		}
		lines = append(lines, line)
	}

	return lines
}

func (trailMap *TrailMap) initializeCounts() {
	rows := len(trailMap.grid)
	cols := len(trailMap.grid[0])
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		for j := 0; j < cols; j++ {
			if trailMap.grid[i][j] == 9 {
				row[j] = 1
			} else {
				row[j] = -1
			}
		}
		result[i] = row
	}
	trailMap.counts = result
}

func (trailMap *TrailMap) numberOfFullTrailsFromPoint(row, col, valToSearchFor int) int {
	count := 0
	for i := -1; i <= 1; i += 2 {
		newRow := row + i
		if newRow >= 0 && newRow < len(trailMap.grid) && trailMap.grid[newRow][col] == valToSearchFor {
			if trailMap.counts[newRow][col] > -1 {
				count += trailMap.counts[newRow][col]
			} else {
				count += trailMap.numberOfFullTrailsFromPoint(newRow, col, valToSearchFor+1)
			}
		}

		newCol := col + i
		if newCol >= 0 && newCol < len(trailMap.grid[0]) && trailMap.grid[row][newCol] == valToSearchFor {
			if trailMap.counts[row][newCol] > -1 {
				count += trailMap.counts[row][newCol]
			} else {
				count += trailMap.numberOfFullTrailsFromPoint(row, newCol, valToSearchFor+1)
			}
		}
	}

	trailMap.counts[row][col] = count

	return count
}

func (trailMap *TrailMap) numberOfEndpointsReachable(row, col int) int {
	seen := make(map[string]bool)

	queue := [][]int{[]int{row, col, 1}}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		currRow, currCol, valToSearchFor := next[0], next[1], next[2]

		for i := -1; i <= 1; i += 2 {
			newRow := currRow + i
			if newRow >= 0 && newRow < len(trailMap.grid) && trailMap.grid[newRow][currCol] == valToSearchFor {
				if trailMap.grid[newRow][currCol] == 9 {
					key := fmt.Sprintf("%d-%d", newRow, currCol)
					_, ok := seen[key]
					if !ok {
						seen[key] = true
					}
				} else {
					queue = append(queue, []int{newRow, currCol, valToSearchFor + 1})
				}
			}

			newCol := currCol + i
			if newCol >= 0 && newCol < len(trailMap.grid[0]) && trailMap.grid[currRow][newCol] == valToSearchFor {
				if trailMap.grid[currRow][newCol] == 9 {
					key := fmt.Sprintf("%d-%d", currRow, newCol)
					_, ok := seen[key]
					if !ok {
						seen[key] = true
					}
				} else {
					queue = append(queue, []int{currRow, newCol, valToSearchFor + 1})
				}
			}
		}
	}

	return len(seen)
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)
	trailMap := &TrailMap{grid: input}
	trailMap.initializeCounts()
	total := 0
	for i, row := range trailMap.grid {
		for j, val := range row {
			if val == 0 {
				count := trailMap.numberOfEndpointsReachable(i, j)
				total += count
			}
		}
	}
	fmt.Println(total)
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)
	trailMap := &TrailMap{grid: input}
	trailMap.initializeCounts()
	total := 0
	for i, row := range trailMap.grid {
		for j, val := range row {
			if val == 0 {
				count := trailMap.numberOfFullTrailsFromPoint(i, j, 1)
				total += count
			}
		}
	}
	fmt.Println(total)
}
