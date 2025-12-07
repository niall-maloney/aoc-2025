package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := 0
	var m [][]string
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

		m = append(m, make([]string, len(strs)))

		for j, str := range strs {
			if m[i][j] == "" {
				m[i][j] = str
			}
			if i > 0 {
				prev := m[i-1][j]
				if prev == "S" || prev == "|" {
					if str == "^" {
						if m[i][j-1] != "|" {
							m[i][j-1] = "|"
						}
						if m[i][j+1] != "|" {
							m[i][j+1] = "|"
						}
						s += 1
					} else {
						m[i][j] = "|"
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
	fmt.Printf("%d\n", s)
}
