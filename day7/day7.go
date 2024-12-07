package day7

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func read(scanner *bufio.Scanner) map[int][]int {
	result := make(map[int][]int)
	for scanner.Scan() {
		current := scanner.Text()
		split := strings.Split(current, ": ")
		testValue, _ := strconv.Atoi(split[0])
		splitValues := strings.Fields(split[1])
		values := make([]int, len(splitValues))

		for i, value := range splitValues {
			converted, _ := strconv.Atoi(value)
			values[i] = converted
		}

		result[testValue] = values
	}

	return result
}

func calculate(values []int, goal, currentTotal, index int) bool {
	if index == len(values)-1 {
		if currentTotal+values[index] == goal || currentTotal*values[index] == goal {
			return true
		}
		return false
	}

	if calculate(values, goal, currentTotal+values[index], index+1) {
		return true
	}

	if calculate(values, goal, currentTotal*values[index], index+1) {
		return true
	}
	return false
}

func concat(value1, value2 int) int {
	exp := math.Floor(math.Log10(float64(value2))) + 1
	mult := int(math.Pow(10, exp))
	return value1*mult + value2
}

func calculate2(values []int, goal, currentTotal, index int) bool {
	if index == len(values)-1 {
		if currentTotal+values[index] == goal || currentTotal*values[index] == goal || concat(currentTotal, values[index]) == goal {
			return true
		}
		return false
	}

	if calculate2(values, goal, currentTotal+values[index], index+1) {
		return true
	}

	if calculate2(values, goal, currentTotal*values[index], index+1) {
		return true
	}

	if calculate2(values, goal, concat(currentTotal, values[index]), index+1) {
		return true
	}

	return false
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)
	total := 0
	for goal, values := range input {
		if calculate(values, goal, 0, 0) {
			total += goal
		}
	}

	fmt.Println(total)
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)
	total := 0
	for goal, values := range input {
		if calculate2(values, goal, 0, 0) {
			total += goal
		}
	}

	fmt.Println(total)
}
