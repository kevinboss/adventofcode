package main

import (
	"bufio"
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
	for i := 0; i < len(matrix)-3; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			words = append(
				words,
				string([]rune{matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]}),
				string([]rune{matrix[i][j], matrix[i][j], matrix[i][j], matrix[i][j]}),
				string([]rune{matrix[i][j], matrix[i][j+1], matrix[i][j+2], matrix[i][j+3]}),
			)
		}
	}
}
