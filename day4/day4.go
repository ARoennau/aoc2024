package day4

import (
	"bufio"
	"fmt"
	"strings"
)

func read(scanner *bufio.Scanner) [][]string {
	lines := make([][]string, 0)
	for scanner.Scan() {
		current := scanner.Text()
		lines = append(lines, strings.Split(current, ""))
	}

	return lines
}

const word string = "XMAS"

func checkOneDirectionForString(grid [][]string, direction []int, previous []int, searchLetterIndex int) int {
	newRow := previous[0] + direction[0]
	newCol := previous[1] + direction[1]
	if newRow >= len(grid) || newRow < 0 || newCol >= len(grid[0]) || newCol < 0 || grid[newRow][newCol] != string(word[searchLetterIndex]) {
		return 0
	}

	if searchLetterIndex == len(word)-1 {
		return 1
	}

	return checkOneDirectionForString(grid, direction, []int{newRow, newCol}, searchLetterIndex+1)
}

func getLocationValue(grid [][]string, location []int) string {
	return grid[location[0]][location[1]]
}

func checkXmasString(grid [][]string, location []int) int {
	if location[0] <= 0 || location[0] >= len(grid)-1 || location[1] <= 0 || location[1] >= len(grid[0])-1 {
		return 0
	}

	upperLeft := []int{location[0] - 1, location[1] - 1}
	upperRight := []int{location[0] - 1, location[1] + 1}
	bottomLeft := []int{location[0] + 1, location[1] - 1}
	bottomRight := []int{location[0] + 1, location[1] + 1}

	diagonal1Valid := (getLocationValue(grid, upperLeft) == "M" && getLocationValue(grid, bottomRight) == "S") || (getLocationValue(grid, upperLeft) == "S" && getLocationValue(grid, bottomRight) == "M")
	diagonal2Valid := (getLocationValue(grid, upperRight) == "M" && getLocationValue(grid, bottomLeft) == "S") || (getLocationValue(grid, upperRight) == "S" && getLocationValue(grid, bottomLeft) == "M")

	if diagonal1Valid && diagonal2Valid {
		return 1
	}

	return 0
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)

	total := 0
	for row, rowString := range input {
		for col, value := range rowString {
			if value == string(word[0]) {
				for i := -1; i < 2; i++ {
					for j := -1; j < 2; j++ {
						if i == 0 && j == 0 {
							continue
						}
						total += checkOneDirectionForString(input, []int{i, j}, []int{row, col}, 1)
					}
				}
			}
		}
	}
	fmt.Println(total)
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)

	total := 0
	for row, rowString := range input {
		for col, value := range rowString {
			if value == "A" {
				total += checkXmasString(input, []int{row, col})
			}
		}
	}
	fmt.Println(total)
}
