package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	var points []image.Point
	var rectangles []image.Rectangle

	file, err := os.Open("day9/task1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, ",")

		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		p := image.Point{X: x, Y: y}

		for _, q := range points {
			rectangles = append(rectangles, image.Rectangle{Min: p, Max: q}.Canon())
		}
		points = append(points, p)

		i++
	}

	s := 0
	for _, r := range rectangles {
		r.Max = r.Max.Add(image.Point{X: 1, Y: 1})
		area := r.Dx() * r.Dy()
		s = max(s, area)
	}

	fmt.Printf("%d\n", s)
}
