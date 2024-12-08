package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, _ := os.ReadFile("3/input.txt")
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := r.FindAllStringSubmatch(string(content), -1)
	var total int
	var enabled = true
	for _, match := range matches {
		switch {
		case match[0] == "do()":
			enabled = true
		case match[0] == "don't()":
			enabled = false
		}
		if enabled {
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])
			total += right * left
		}
	}
	fmt.Println(total)
}
