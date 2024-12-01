package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
		Part1(scanner)
	case 2:
		Part2(scanner)
	}
}
