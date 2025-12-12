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
	s := 0

	var points []image.Point
	var rectangles []image.Rectangle
	var edges []image.Rectangle

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

		if len(points) > 1 {
			edges = append(edges, image.Rectangle{Min: points[i-1], Max: points[i]}.Canon())
		}
		i++
	}
	edges = append(edges, image.Rectangle{Min: points[len(points)-1], Max: points[0]}.Canon())

loop:
	for _, r := range rectangles {
		r.Max = r.Max.Add(image.Point{X: 1, Y: 1})
		area := r.Dx() * r.Dy()

		for _, v := range edges {
			v.Max = v.Max.Add(image.Point{X: 1, Y: 1})
			if v.Overlaps(r.Inset(1)) {
				continue loop
			}
		}
		s = max(s, area)
	}

	fmt.Printf("%d\n", s)
}
