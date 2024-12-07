package day5

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Adjacency map[int][]int

type ParsedDay5Input struct {
	adjacency Adjacency
	updates   [][]int
}

func read(scanner *bufio.Scanner) ParsedDay5Input {
	adjacency := make(map[int][]int)
	for scanner.Scan() {
		current := scanner.Text()
		if current == "" {
			break
		}

		split := strings.Split(current, "|")
		node, _ := strconv.Atoi(split[1])
		neighbor, _ := strconv.Atoi(split[0])
		_, ok := adjacency[node]
		if !ok {
			adjacency[node] = []int{neighbor}
		} else {
			adjacency[node] = append(adjacency[node], neighbor)
		}
	}

	updates := make([][]int, 0)

	for scanner.Scan() {
		current := scanner.Text()
		split := strings.Split(current, ",")
		row := make([]int, 0)
		for _, value := range split {
			converted, _ := strconv.Atoi(value)
			row = append(row, converted)
		}

		updates = append(updates, row)
	}

	return ParsedDay5Input{adjacency: adjacency, updates: updates}
}

func (a Adjacency) isHigher(current, goal int) bool {
	neighbors, ok := a[current]
	if !ok {
		return false
	}

	if slices.Index(neighbors, goal) > -1 {
		return true
	}

	return false
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)
	total := 0
	for _, row := range input.updates {
		isValid := true
		for i, value := range row[:len(row)-1] {
			if input.adjacency.isHigher(value, row[i+1]) {
				isValid = false
				break
			}
		}

		if isValid {
			mid := len(row) / 2
			total += row[mid]
		}
	}
	fmt.Println(total)
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)
	invalid := make([][]int, 0)
	for _, row := range input.updates {
		isValid := true
		for i, value := range row[:len(row)-1] {
			if input.adjacency.isHigher(value, row[i+1]) {
				isValid = false
				break
			}
		}

		if !isValid {
			invalid = append(invalid, row)
		}
	}

	total := 0
	for _, row := range invalid {
		slices.SortFunc(row, func(a, b int) int {
			if input.adjacency.isHigher(a, b) {
				return 1
			}

			return -1
		})

		mid := len(row) / 2
		total += row[mid]
	}

	fmt.Println(total)
}
