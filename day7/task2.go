package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := 0
	var m [][]int
	printm := false

	file, err := os.Open("day7/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		strs := strings.Split(line, "")

		m = append(m, make([]int, len(strs)))

		for j, str := range strs {
			if str == "S" {
				m[i][j] = 1
			}
			if i > 0 {
				prev := m[i-1][j]
				if prev > 0 {
					if str == "^" {
						m[i][j-1] += prev
						m[i][j+1] += prev
					} else {
						m[i][j] += prev
					}
				}
			}
		}
	}

	if printm {
		for _, r := range m {
			fmt.Printf("%v\n", r)
		}
	}

	for _, t := range m[len(m)-1] {
		s += t
	}
	fmt.Printf("%d\n", s)
}
