package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := 0
	var sums [][]string

	file, err := os.Open("day6/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)
		sums = append(sums, strs)
	}

	for i := 0; i < len(sums[0]); i++ {
		current := 0
		for j := 0; j < len(sums)-1; j++ {
			e, _ := strconv.Atoi(sums[j][i])
			switch sums[len(sums)-1][i] {
			case "+":
				current += e
			case "-":
				current -= e
			case "*":
				if current == 0 {
					current = 1
				}
				current *= e
			case "/":
				if current == 0 {
					current = 1
				}
				current /= e
			}
		}
		s += current
	}

	fmt.Printf("%d\n", s)
}
