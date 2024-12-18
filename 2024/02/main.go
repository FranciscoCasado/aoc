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

func (r Report) isSafe() bool {
	for i := range r.Levels {
		if i < 1 {
			continue
		}

		if abs(r.Levels[i]-r.Levels[i-1]) > 3 {
			return false
		}
		if abs(r.Levels[i]-r.Levels[i-1]) < 1 {
			return false
		}

		if i < 2 {
			continue
		}

		if (r.Levels[i]-r.Levels[i-1])*(r.Levels[i-1]-r.Levels[i-2]) < 0 {
			return false
		}
	}
	return true
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
	for _, report := range reports {
		if report.isSafe() {
			count++
		}
	}
	fmt.Printf("Total Safe Count: %v\n", count)

}
