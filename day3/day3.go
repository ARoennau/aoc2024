package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func read(scanner *bufio.Scanner) string {
	line := ""
	for scanner.Scan() {
		current := scanner.Text()
		line += current
	}

	return line
}

func Part1(scanner *bufio.Scanner) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	line := read(scanner)

	total := int64(0)
	found := re.FindAllStringSubmatch(line, -1)
	for _, match := range found {
		val1, _ := strconv.Atoi(match[1])
		val2, _ := strconv.Atoi(match[2])
		total += int64(val1) * int64(val2)
	}

	fmt.Println(total)
}

func Part2(scanner *bufio.Scanner) {
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)
	line := read(scanner)

	total := int64(0)
	found := re.FindAllStringSubmatch(line, -1)
	enabled := true
	for _, match := range found {
		if match[0] == "don't()" {
			enabled = false
		} else if match[0] == "do()" {
			enabled = true
		} else if enabled {
			val1, _ := strconv.Atoi(match[2])
			val2, _ := strconv.Atoi(match[3])
			total += int64(val1) * int64(val2)
		}
	}

	fmt.Println(total)
}
