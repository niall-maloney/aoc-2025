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

	file, err := os.Open("day2/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, rng := range ranges {
			boundaries := strings.Split(rng, "-")
			start, _ := strconv.Atoi(boundaries[0])
			end, _ := strconv.Atoi(boundaries[1])
			for i := start; i <= end; i++ {
				id := strconv.Itoa(i)
				if id[:len(id)/2] == id[len(id)/2:] {
					sum += i
				}
			}
		}
	}

	fmt.Printf("%d\n", sum)
}
