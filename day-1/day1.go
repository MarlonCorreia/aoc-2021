package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var digits []int

	for scanner.Scan() {
		digit, _ := strconv.Atoi(scanner.Text())
		digits = append(digits, digit)
	}

	result_1 := part1(digits)
	result_2 := part2(digits)

	fmt.Println(result_1)
	fmt.Println(result_2)
}

func part1(digits []int) int {
	result := 0

	for idx, digit := range digits {
		if idx == 0 {
			continue
		} else if digit > digits[idx-1] {
			result++
		}
	}
	return result
}

func part2(digits []int) int {
	result := 0
	var sums []int

	for i := 0; i < len(digits); i++ {
		if i+3 > len(digits) {
			break
		}
		cur := _sum(digits[i : i+3])
		sums = append(sums, cur)
	}

	for idx, sum := range sums {
		if idx == 0 {
			continue
		} else if sum > sums[idx-1] {
			result++
		}
	}

	return result
}

func _sum(numbs []int) int {
	var result int
	for _, val := range numbs {
		result += val
	}
	return result
}
