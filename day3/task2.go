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
	length := 12

	file, err := os.Open("day3/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		joltageStrs := strings.Split(line, "")

		number := 0
		var digits []int
		for i, joltageStr := range joltageStrs {
			joltage, _ := strconv.Atoi(joltageStr)

			for len(digits) > 0 && digits[len(digits)-1] < joltage &&
				len(joltageStrs)-i+len(digits) > length {
				digits = digits[:len(digits)-1]
			}
			if len(digits) < length {
				digits = append(digits, joltage)
			}
		}

		for _, digit := range digits {
			number = number*10 + digit
		}
		sum += number
	}

	fmt.Printf("%d\n", sum)
}
