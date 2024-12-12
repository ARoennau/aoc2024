package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ARoennau/aoc2024/day1"
	"github.com/ARoennau/aoc2024/day10"
	"github.com/ARoennau/aoc2024/day11"
	"github.com/ARoennau/aoc2024/day2"
	"github.com/ARoennau/aoc2024/day3"
	"github.com/ARoennau/aoc2024/day4"
	"github.com/ARoennau/aoc2024/day5"
	"github.com/ARoennau/aoc2024/day6"
	"github.com/ARoennau/aoc2024/day7"
	"github.com/ARoennau/aoc2024/day8"
)

func main() {
	number := flag.Int("n", 1, "The number of the problem (parts 1 and 2 are separate, so there will be a total of 50)")
	isTest := flag.Bool("test", true, "Indicates whether to use the test data")
	flag.Parse()

	fileName := strconv.Itoa(*number)
	if *isTest {
		fileName += "-test"
	}

	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filePath := fmt.Sprintf("%s/input/%s.txt", basePath, fileName)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch *number {
	case 1:
		day1.Part1(scanner)
	case 2:
		day1.Part2(scanner)
	case 3:
		day2.Part1(scanner)
	case 4:
		day2.Part2(scanner)
	case 5:
		day3.Part1(scanner)
	case 6:
		day3.Part2(scanner)
	case 7:
		day4.Part1(scanner)
	case 8:
		day4.Part2(scanner)
	case 9:
		day5.Part1(scanner)
	case 10:
		day5.Part2(scanner)
	case 11:
		day6.Part1(scanner)
	case 12:
		day6.Part2(scanner)
	case 13:
		day7.Part1(scanner)
	case 14:
		day7.Part2(scanner)
	case 15:
		day8.Part1(scanner)
	case 16:
		day8.Part2(scanner)
		// Accidentally deleted day 9 and I'm not going to redo it
	case 19:
		day10.Part1(scanner)
	case 20:
		day10.Part2(scanner)
	case 21:
		day11.Part1(scanner)
	case 22:
		day11.Part2(scanner)
	}
}
