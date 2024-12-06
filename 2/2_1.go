package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2/input.txt")
	scanner := bufio.NewScanner(file)
	var safeLevels = 0
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		var numbers []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers = append(numbers, num)
		}
		var safeLevel = true
		for i := 0; i < len(numbers)-1; i++ {
			left := numbers[i]
			right := numbers[i+1]
			diff := left - right
			var absDiff int
			if diff > 0 {
				absDiff = diff
			} else {
				absDiff = -diff
			}
			if absDiff < 1 || absDiff > 3 {
				safeLevel = false
			}
		}
		for i := 1; i < len(numbers)-1; i++ {
			left := numbers[i-1]
			middle := numbers[i]
			right := numbers[i+1]
			if (left < middle && middle > right) || (left > middle && middle < right) {
				safeLevel = false
			}
		}
		if safeLevel {
			safeLevels++
		}
	}
	fmt.Println(safeLevels)
}
