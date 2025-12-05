package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day4/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := populateMatrix(file)
	convolution := kernelConvolution(matrix)
	product(convolution, matrix)
	mask(convolution, func(n int) int {
		if n >= 4 {
			return n
		}
		return 0
	})
	filter(convolution)

	s := sum(matrix) - sum(convolution)
	fmt.Printf("%d\n", s)
}

func sum(matrix [][]int) int {
	s := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			s += matrix[i][j]
		}
	}
	return s
}

func filter(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 1
			}
		}
	}
}

func mask(matrix [][]int, m func(int) int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = m(matrix[i][j])
		}
	}
}

func product(a [][]int, b [][]int) {
	for i := 0; i < len(a); i++ {
		row := a[i]
		for j := 0; j < len(row); j++ {
			a[i][j] = a[i][j] * b[i][j]
		}
	}
}

func kernelConvolution(matrix [][]int) [][]int {
	convolution := createMatrix(len(matrix), len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row); j++ {
			convolution[i][j] = sumNeighbours(matrix, i, j)
		}
	}
	return convolution
}

func sumNeighbours(matrix [][]int, row int, column int) int {
	sum := 0
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i >= 0 && j >= 0 &&
				i < len(matrix) && j < len(matrix[0]) &&
				(i != row || j != column) {
				sum += matrix[i][j]
			}
		}
	}
	return sum
}

func populateMatrix(file *os.File) [][]int {
	var matrix [][]int
	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		positions := strings.Split(line, "")

		matrix = append(matrix, make([]int, len(positions)))

		for i, position := range positions {
			switch position {
			case ".":
				matrix[index][i] = 0
			case "@":
				matrix[index][i] = 1
			}
		}
		index++
	}
	return matrix
}

func createMatrix(x int, y int) [][]int {
	matrix := make([][]int, x)
	for i := 0; i < y; i++ {
		matrix[i] = make([]int, y)
	}
	return matrix
}
