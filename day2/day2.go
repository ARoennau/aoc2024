package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func read(scanner *bufio.Scanner) [][]int64 {
	lines := make([][]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		current := make([]int64, 0)
		for _, stringNum := range split {
			num, _ := strconv.Atoi(stringNum)
			current = append(current, int64(num))
		}
		lines = append(lines, current)
	}

	return lines
}

func Part1(scanner *bufio.Scanner) {
	lines := read(scanner)
	safeCount := 0
	for _, line := range lines {
		isSafe := true

		if line[0] == line[1] {
			continue
		}

		isIncreasing := line[1] > line[0]

		for i, current := range line[:len(line)-1] {
			next := line[i+1]
			var difference int64

			if isIncreasing {
				difference = next - current
			} else {
				difference = current - next
			}

			if difference < 1 || difference > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount += 1
		}
	}
	fmt.Println(safeCount)
}

func Part2(scanner *bufio.Scanner) {
	lines := read(scanner)
	safeCount := 0
	for _, line := range lines {
		badIndex := findFirstBadIndex(line)
		if badIndex == len(line) {
			safeCount++
		} else {
			// Too many edge cases, decided to brute force
			if testIfRemovalWorks(line, badIndex-1) ||
				testIfRemovalWorks(line, badIndex) ||
				testIfRemovalWorks(line, badIndex+1) {
				safeCount++
			}
		}
	}
	fmt.Println(safeCount)
}

func testIfRemovalWorks(line []int64, index int) bool {
	if index < 0 || index >= len(line) {
		return false
	}

	updated := make([]int64, len(line))
	copy(updated, line)
	updated = append(updated[:index], updated[index+1:]...)
	badIndex := findFirstBadIndex(updated)
	return badIndex == len(updated)
}

func findFirstBadIndex(line []int64) int {
	if line[0] == line[1] {
		return 0
	}

	isIncreasing := line[1] > line[0]

	for i := 0; i < len(line)-1; i++ {
		current := line[i]
		next := line[i+1]
		var difference int64

		if isIncreasing {
			difference = next - current
		} else {
			difference = current - next
		}

		if difference < 1 || difference > 3 {
			return i
		}
	}
	return len(line)
}
