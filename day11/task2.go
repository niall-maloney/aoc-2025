package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	input   string
	outputs []string
}

func main() {
	file, err := os.Open("day11/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	instructions := map[string]Instruction{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, ":")
		input := strings.TrimSpace(strs[0])
		outputs := strings.Split(strings.TrimSpace(strs[1]), " ")
		instructions[input] = Instruction{input, outputs}
	}

	a := countPath("svr", "fft", instructions, map[string]int{})
	a *= countPath("fft", "dac", instructions, map[string]int{})
	a *= countPath("dac", "out", instructions, map[string]int{})

	b := countPath("svr", "dac", instructions, map[string]int{})
	b *= countPath("dac", "fft", instructions, map[string]int{})
	b *= countPath("fft", "out", instructions, map[string]int{})

	fmt.Printf("%d\n", a+b)
}

func countPath(current, target string, instructions map[string]Instruction, cache map[string]int) int {
	count, exists := cache[current]
	if exists {
		return count
	}

	if current == target {
		return 1
	}

	instruction, exists := instructions[current]
	if !exists {
		return 0
	}

	total := 0
	for _, next := range instruction.outputs {
		total += countPath(next, target, instructions, cache)
	}

	cache[current] = total

	return total
}
