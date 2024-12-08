package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("4/input.txt")
	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}
	var words []string
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if row < len(matrix)-3 {
				words = append(words, string([]rune{matrix[row][column], matrix[row+1][column], matrix[row+2][column], matrix[row+3][column]}))
			}
			if column < len(matrix[0])-3 {
				words = append(words, string([]rune{matrix[row][column], matrix[row][column+1], matrix[row][column+2], matrix[row][column+3]}))
			}
			if column >= 3 && row < len(matrix)-3 {
				words = append(words, string([]rune{matrix[row][column], matrix[row+1][column-1], matrix[row+2][column-2], matrix[row+3][column-3]}))
			}
			if row < len(matrix)-3 && column < len(matrix[0])-3 {
				words = append(words, string([]rune{matrix[row][column], matrix[row+1][column+1], matrix[row+2][column+2], matrix[row+3][column+3]}))
			}
		}
	}

	var count int
	for _, word := range words {
		if word == "XMAS" || word == "SAMX" {
			count++
		}
	}
	fmt.Println(count)
}
