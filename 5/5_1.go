package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule1 struct {
	left, right int
}

func main() {
	file, _ := os.Open("5/input.txt")
	scanner := bufio.NewScanner(file)
	var rules []Rule1
	var updates [][]int
	section1 := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			section1 = false
		} else if section1 {
			parts := strings.Split(line, "|")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule1{
				left:  left,
				right: right,
			})
		} else {
			parts := strings.Split(line, ",")
			var segments []int
			for _, part := range parts {
				segment, _ := strconv.Atoi(part)
				segments = append(segments, segment)
			}
			updates = append(updates, segments)
		}
	}

	total := 0
	for _, update := range updates {
		ruleViolated := false
		previous := []int{update[0]}
		for i := 1; i < len(update); i++ {
			left := update[i]
			for _, rule := range rules {
				if rule.left == left {
					for _, p := range previous {
						if p == rule.right {
							ruleViolated = true
						}
					}
				}
			}
			previous = append(previous, left)
		}
		if !ruleViolated {
			middle := update[len(update)/2]
			total += middle
		}
	}

	fmt.Println(total)
}
