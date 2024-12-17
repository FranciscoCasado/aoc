package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		numbers := strings.Split(scanner.Text(), "   ")

		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			continue
		}

		second, err := strconv.Atoi(numbers[1])
		if err != nil {
			continue
		}

		left = append(left, first)
		right = append(right, second)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No input filename provided")
		return
	}

	left, right, err := readInput(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	left_sorted := make([]int, len(left))
	copy(left_sorted, left)
	sort.Ints(left_sorted)

	right_sorted := make([]int, len(right))
	copy(right_sorted, right)
	sort.Ints(right_sorted)

	diff := make([]int, len(left))
	sum := 0

	for i := range left {
		diff[i] = abs(left_sorted[i] - right_sorted[i])
		sum += diff[i]
		fmt.Printf("[%3d]\t%v\t%v\t%v\n", i, left_sorted[i], right_sorted[i], diff[i])
	}
	fmt.Printf("Total: %v\n", sum)

}
