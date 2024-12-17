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

func distance(left []int, right []int) int {
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
	}
	return sum
}

func similarity(left []int, right []int) int {
	numbers := make(map[int]int)

	for _, value := range left {
		numbers[value] = 0
	}

	for _, value := range right {
		if _, exists := numbers[value]; exists {
			numbers[value] += 1
		}
	}

	var sim = 0

	for key, value := range numbers {
		sim += key * value
		fmt.Printf("[%d]\t%d\t%d\n", key, value, sim)
	}

	return sim
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

	dist := distance(left, right)
	sim := similarity(left, right)

	fmt.Printf("Distance: %v\n", dist)
	fmt.Printf("Similarity: %v\n", sim)

}
