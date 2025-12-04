package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	file, err := os.Open("day3/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		joltageStrs := strings.Split(line, "")

		first := 0
		second := 0
		for i, joltageStr := range joltageStrs {
			joltage, _ := strconv.Atoi(joltageStr)
			if joltage > first && i < len(joltageStrs)-1 {
				first = joltage
				second = 0
			} else if joltage > second {
				second = joltage
			}
		}
		sum += first*10 + second
	}

	fmt.Printf("%d\n", sum)
}
