package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Levels []int
}

func NewReport(content string) Report {
	var levels []int
	for _, str := range strings.Split(content, " ") {
		num, _ := strconv.Atoi(str)
		levels = append(levels, num)
	}

	return Report{
		Levels: levels,
	}
}

func readInput(filename string) ([]Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := []Report{}
	for scanner.Scan() {
		reports = append(reports, NewReport(scanner.Text()))
	}

	return reports, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func testArraySafe(arr []int) bool {
	for i := range arr {
		if i < 1 {
			continue
		}

		if abs(arr[i]-arr[i-1]) > 3 {
			return false
		}
		if abs(arr[i]-arr[i-1]) < 1 {
			return false
		}

		if i < 2 {
			continue
		}

		if (arr[i]-arr[i-1])*(arr[i-1]-arr[i-2]) < 0 {
			return false
		}
	}
	return true
}

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
func (r Report) isSafe() bool {
	return testArraySafe(r.Levels)
}

func (r Report) isDampenedSafe() bool {
	if r.isSafe() {
		return true
	}

	for i := range r.Levels {
		dampened := make([]int, len(r.Levels))
		copy(dampened, r.Levels)
		dampened = removeIndex(dampened, i)

		if testArraySafe(dampened) {
			return true
		}

	}
	return false
}

func main() {
	fmt.Println("AoC 2024 - Day 2")

	if len(os.Args) <= 1 {
		fmt.Println("No input filename provided")
		return
	}

	reports, err := readInput(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var count = 0
	var count_dampened = 0
	for _, report := range reports {
		if report.isSafe() {
			count++
		}
		if report.isDampenedSafe() {
			count_dampened++
		}
	}
	fmt.Printf("Total Safe Count: %v\n", count)
	fmt.Printf("Total Dampened Safe Count: %v\n", count_dampened)

}
