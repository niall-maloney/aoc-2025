package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	lower int
	upper int
}

func main() {
	s := 0
	var ranges []Range

	file, err := os.Open("day5/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "-")
		if len(strs) == 2 {
			lower, _ := strconv.Atoi(strs[0])
			upper, _ := strconv.Atoi(strs[1])
			ranges = append(ranges, Range{lower, upper})
		}
		if len(strs) == 1 {
			ingredient, _ := strconv.Atoi(strs[0])
			fresh := false
			for _, r := range ranges {
				if r.lower <= ingredient && ingredient <= r.upper {
					fresh = true
				}
			}
			if fresh {
				s++
			}
		}
	}

	fmt.Printf("%d\n", s)
}
