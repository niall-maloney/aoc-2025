package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	lower int
	upper int
}

func main() {
	s := 0
	var ranges []*Range

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
			ranges = append(ranges, &Range{lower, upper})
		}
	}
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].lower < ranges[j].lower })

	var fresh []*Range
	for _, r := range ranges {
		shouldAppend := true
		for _, f := range fresh {
			if f.lower <= r.lower && r.lower <= f.upper ||
				f.lower <= r.upper && r.upper <= f.upper {
				f.lower = min(f.lower, r.lower)
				f.upper = max(f.upper, r.upper)
				shouldAppend = false
			}
		}
		if shouldAppend {
			fresh = append(fresh, r)
		}
	}

	for _, f := range fresh {
		s += f.upper - f.lower + 1
	}

	fmt.Printf("%d\n", s)
}
