package main

import (
	"bufio"
	"fmt"
	"os"
)

type Xmas struct {
	word1, word2 string
}

func main() {
	file, _ := os.Open("4/input.txt")
	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}
	var xmases []Xmas
	for row := 1; row < len(matrix)-1; row++ {
		for column := 1; column < len(matrix[0])-1; column++ {
			xmases = append(xmases, Xmas{
				string([]rune{matrix[row-1][column-1], matrix[row][column], matrix[row+1][column+1]}),
				string([]rune{matrix[row-1][column+1], matrix[row][column], matrix[row+1][column-1]}),
			})
		}
	}

	var count int
	for _, xmas := range xmases {
		if isMas(xmas.word1) && isMas(xmas.word2) {
			count++
		}
	}
	fmt.Println(count)
}

func isMas(word string) bool {
	return word == "MAS" || word == "SAM"
}
