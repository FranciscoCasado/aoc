package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Multiplication struct {
	First  int
	Second int
}

func NewMultiplicationFromString(s string) Multiplication {
	s = strings.Replace(s, "mul(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	numbers := strings.Split(s, ",")
	first, _ := strconv.Atoi(numbers[0])
	second, _ := strconv.Atoi(numbers[1])
	return Multiplication{
		First:  first,
		Second: second,
	}
}

func (m Multiplication) Result() int {
	return m.First * m.Second
}

func findMultiplications(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return pattern.FindAllString(string(content), -1), nil
}

func main() {
	fmt.Println("AoC 2024 - Day 2")

	if len(os.Args) <= 1 {
		fmt.Println("No input filename provided")
		return
	}

	mults, err := findMultiplications(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var sum int = 0
	for _, mul := range mults {
		m := NewMultiplicationFromString(mul)
		sum += m.Result()
		fmt.Printf("%v = %v\n", mul, m.Result())
	}

	fmt.Printf("Sum of all valid mults: %v\n", sum)
}
