package day11

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func read(scanner *bufio.Scanner) []string {
	scanner.Scan()
	line := scanner.Text()
	return strings.Split(line, " ")
}

func trimZeroes(stone string) string {
	toInt, _ := strconv.Atoi(stone)
	return strconv.Itoa(toInt)
}

func blink(stones []string) []string {
	result := make([]string, 0)
	for _, stone := range stones {
		if stone == "0" {
			result = append(result, "1")
		} else if len(stone)%2 == 0 {
			mid := len(stone) / 2
			result = append(result, trimZeroes(stone[:mid]), trimZeroes(stone[mid:]))
		} else {
			numberVal, _ := strconv.Atoi(stone)
			newStone := strconv.Itoa(numberVal * 2024)
			result = append(result, newStone)
		}
	}
	return result
}

func Part1(scanner *bufio.Scanner) {
	input := read(scanner)
	current := input
	for i := 0; i < 25; i++ {
		blinked := blink(current)
		current = blinked
	}
	fmt.Println(len(current))
}

// func blink2(seen map[string][]string, stones []string) []string {
// 	result := make([]string, 0)
// 	for _, stone := range stones {
// 		memo, ok := seen[stone]
// 		if ok {
// 			result = append(result, memo...)
// 		} else if len(stone)%2 == 0 {
// 			mid := len(stone) / 2
// 			left := trimZeroes(stone[:mid])
// 			right := trimZeroes(stone[mid:])

// 			result = append(result, left, right)
// 			seen[stone] = []string{left, right}
// 		} else {
// 			numberVal, _ := strconv.Atoi(stone)
// 			newStone := strconv.Itoa(numberVal * 2024)
// 			result = append(result, newStone)
// 			seen[stone] = []string{newStone}
// 		}
// 	}
// 	return result
// }

func getNumberOfDigits(value int) int {
	return int(math.Floor(math.Log10(float64(value)))) + 1
}

func getSplitDigits(value, numDigits int) []int {
	rightTotal := 0
	mid := numDigits / 2
	for i := 0; i < mid; i++ {
		digit := value % 10
		rightTotal += int(math.Pow(10, float64(i))) * digit
		value /= 10
	}

	leftTotal := 0
	for i := 0; i < mid; i++ {
		digit := value % 10
		leftTotal += int(math.Pow(10, float64(i))) * digit
		value /= 10
	}
	return []int{leftTotal, rightTotal}
}

func blink2(seen map[int][]int, stones []int) []int {
	result := make([]int, 0)
	for _, stone := range stones {
		memo, ok := seen[stone]
		if ok {
			result = append(result, memo...)
		} else {
			numberOfDigits := getNumberOfDigits(stone)
			if numberOfDigits%2 == 0 {
				split := getSplitDigits(stone, numberOfDigits)
				result = append(result, split[0], split[1])
				seen[stone] = []int{split[0], split[1]}
			} else {
				newStone := stone * 2024
				result = append(result, newStone)
				seen[stone] = []int{newStone}
			}
		}
	}
	return result
}

type Node struct {
	val  int
	next *Node
}

func (node *Node) Length() int {
	i := 0
	current := node
	for current != nil {
		i++
		current = current.next
	}
	return i
}

func (node *Node) String() string {
	ints := make([]string, 0)
	current := node
	for current != nil {
		converted := strconv.Itoa(current.val)
		ints = append(ints, converted)
		current = current.next
	}
	return strings.Join(ints, ", ")
}

func blink2List(seen map[int][]int, stones *Node) *Node {
	current := stones
	for current != nil {
		stone := *current
		memo, ok := seen[stone.val]
		if ok {
			current.val = memo[0]
			if len(memo) == 2 {
				nextPointer := stone.next
				newStone := &Node{val: memo[1], next: nextPointer}
				current.next = newStone
				current = nextPointer
				continue
			}
			current = current.next
		} else {
			numberOfDigits := getNumberOfDigits(stone.val)
			if numberOfDigits%2 == 0 {
				split := getSplitDigits(stone.val, numberOfDigits)
				nextPointer := stone.next
				newStone := &Node{val: split[1], next: nextPointer}
				current.val = split[0]
				current.next = newStone
				seen[stone.val] = []int{split[0], split[1]}
				current = nextPointer
			} else {
				newStone := stone.val * 2024
				current.val = newStone
				seen[stone.val] = []int{newStone}
				current = current.next
			}
		}
	}
	return stones
}

func Part2(scanner *bufio.Scanner) {
	input := read(scanner)
	ints := make([]int, len(input))
	for i, str := range input {
		converted, _ := strconv.Atoi(str)
		ints[i] = converted
	}

	seen := make(map[int][]int)
	seen[0] = []int{1}
	timesToBlink := 50
	total := 0
	for _, value := range ints {
		current := &Node{val: value}
		for i := 0; i < timesToBlink; i++ {
			blinked := blink2List(seen, current)
			current = blinked
		}
		total += current.Length()
	}
	// for _, value := range ints[:1] {
	// 	current := []int{value}
	// 	for i := 0; i < timesToBlink; i++ {
	// 		fmt.Println(i)
	// 		blinked := blink2(seen, current)
	// 		current = blinked
	// 	}
	// 	total += len(current)
	// }
	// fmt.Println(len(seen))
	fmt.Println(total)
}
