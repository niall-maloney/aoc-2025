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
	maxLen := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "")
		sums = append(sums, strs)
		maxLen = max(maxLen, len(strs))
	}

	pad(sums, maxLen)

	currNum := 0
	nums := make([]int, len(sums))
	for clmn := maxLen - 1; clmn >= 0; clmn-- {
		rowRemainingCount := len(sums) - 1
		for row := 0; row < len(sums); row++ {
			elem := sums[row][clmn]
			switch elem {
			case " ":
				if rowRemainingCount == 0 {
					nums = make([]int, len(sums))
					currNum = 0
				} else {
					rowRemainingCount--
				}
			case "+":
				s += add(nums)
			case "*":
				s += product(nums)
			default:
				digit, _ := strconv.Atoi(elem)
				nums[currNum] = nums[currNum]*10 + digit
			}
		}
		currNum++
	}
	fmt.Printf("%d\n", s)
}

func add(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

func product(numbers []int) int {
	result := 1
	for _, n := range numbers {
		if n > 0 {
			result *= n
		}
	}
	return result
}

func pad(arr [][]string, size int) {
	for r := range arr {
		if len(arr[r]) < size {
			for i := len(arr[r]); i < size; i++ {
				arr[r] = append(arr[r], " ")
			}
		}
	}
}
