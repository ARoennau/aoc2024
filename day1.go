package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day1Arrays struct {
	left  []int64
	right []int64
}

func readArrays(scanner *bufio.Scanner) Day1Arrays {
	left := make([]int64, 0)
	right := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		leftValue, _ := strconv.Atoi(split[0])
		rightValue, _ := strconv.Atoi(split[1])
		left = append(left, int64(leftValue))
		right = append(right, int64(rightValue))
	}

	return Day1Arrays{left, right}
}

func Part1(scanner *bufio.Scanner) {
	arrays := readArrays(scanner)
	left := arrays.left
	right := arrays.right

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	total := 0
	for i := range left {
		total += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(total)
}

func Part2(scanner *bufio.Scanner) {
	arrays := readArrays(scanner)
	left := arrays.left
	right := arrays.right

	counts := make(map[int64]int64)

	for _, number := range right {
		_, ok := counts[number]
		if ok {
			counts[number] += 1
		} else {
			counts[number] = 1
		}
	}

	var total int64

	for _, number := range left {
		count := counts[number]
		total += number * count
	}

	fmt.Println(total)
}
