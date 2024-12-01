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
	var similarityScore int
	for _, left := range leftColumn {
		var occurrences int
		for _, right := range rightColumn {
			if left == right {
				occurrences++
			}
		}
		similarityScore += occurrences * left
	}
	fmt.Println(similarityScore)
}
