package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, _ := os.ReadFile("3/input.txt")
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(string(content), -1)
	var total int
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		total += right * left
	}
	fmt.Println(total)
}
