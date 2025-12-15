package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	s := 0

	file, err := os.Open("day10/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		man := bytes.Split(line, []byte(" "))

		lInstrct := man[0][1 : len(man[0])-1]
		light := 0
		for i, l := range lInstrct {
			if l == '#' {
				light |= 1 << i
			}
		}

		bInstrct := man[1 : len(man)-1]
		buttons := make([]uint, len(bInstrct))
		for i, b := range bInstrct {
			b = b[1 : len(b)-1]
			for s := range bytes.SplitSeq(b, []byte(",")) {
				buttons[i] |= uint(1 << conv(s))
			}
		}

		jInstrct := man[len(man)-1]
		jInstrct = jInstrct[1 : len(jInstrct)-1]
		jBytes := bytes.Split(jInstrct, []byte(","))
		joltages := make([]int, len(jBytes))
		for i, j := range jBytes {
			joltages[i] = conv(j)
		}

		width := 0
		for _, b := range buttons {
			width = max(width, bits.Len(b))
		}

		bitmasks := slices.Repeat([]int{-1}, 1<<width)
		bitmasks[0] = 0

		q := []uint{0}
		for i := 0; i < len(q); i++ {
			current := q[i]
			for _, b := range buttons {
				next := current ^ b

				if bitmasks[next] != -1 {
					continue
				}

				bitmasks[next] = bitmasks[current] + 1
				q = append(q, next)
			}
		}
		s += bitmasks[light]
	}

	fmt.Printf("%d\n", s)
}

func conv(b []byte) int {
	n := 0
	for _, x := range b {
		n = 10*n + int(x-'0')
	}
	return n
}
