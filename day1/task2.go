package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dial := 50
	password := 0

	file, err := os.Open("day1/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		characters := strings.Split(line, "")
		direction := characters[0]
		amount, errAtoi := strconv.Atoi(strings.Join(characters[1:], ""))
		if errAtoi != nil {
			fmt.Printf("%v\n", errAtoi)
			return
		}

		switch direction {
		case "R":
			for i := 0; i < amount; i++ {
				dial++
				check(dial, &password)
			}
		case "L":
			for i := 0; i < amount; i++ {
				dial--
				check(dial, &password)
			}
		}
	}

	fmt.Printf("%d\n", password)
}

func check(position int, password *int) {
	if math.Mod(math.Abs(float64(position)), 100) == 0 {
		*password++
	}
}
