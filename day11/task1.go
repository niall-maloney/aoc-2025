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

	s := countPath("you", instructions)

	fmt.Printf("%d\n", s)
}

func countPath(current string, instructions map[string]Instruction) int {
	if current == "out" {
		return 1
	}

	instruction, exists := instructions[current]
	if !exists {
		return 0
	}

	total := 0

	for _, next := range instruction.outputs {
		total += countPath(next, instructions)
	}

	return total
}
