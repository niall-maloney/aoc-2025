package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	x, y int
}

func main() {
	s := 0
	var locations []Location

	file, err := os.Open("day9/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, ",")

		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])

		locations = append(locations, Location{x, y})
	}

	for i, a := range locations {
		for j, b := range locations {
			if i != j {
				area := math.Abs(float64(a.x-b.x+1)) * math.Abs(float64(a.y-b.y+1))
				if area > float64(s) {
					s = int(area)
				}
			}
		}
	}

	fmt.Printf("%d\n", s)
}
