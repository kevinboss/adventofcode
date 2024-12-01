package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("1/input.txt")

	var leftColumn []int
	var rightColumn []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		leftColumn = append(leftColumn, left)
		rightColumn = append(rightColumn, right)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)
	var distance int
	for i, left := range leftColumn {
		var right = rightColumn[i]
		var diff = left - right
		if diff > 0 {
			distance += diff
		} else {
			distance -= diff
		}
	}
	fmt.Println(distance)
}
