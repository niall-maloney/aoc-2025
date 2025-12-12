package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x, y, z int
}

type Connection struct {
	length float64
	start  int
	end    int
}

type UnionFind struct {
	parent []int
	size   []int
	count  int
}

func main() {
	var junctionBoxes []JunctionBox
	take := 1000
	file, err := os.Open("day8/task1.txt")
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
		z, _ := strconv.Atoi(strs[2])

		junctionBoxes = append(junctionBoxes, JunctionBox{x, y, z})
	}

	uf := newUnionFind(junctionBoxes)

	var connections []Connection
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			d := distance(junctionBoxes[i], junctionBoxes[j])
			connections = append(connections, Connection{d, i, j})
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].length < connections[j].length
	})

	connections = connections[:take]
	for _, c := range connections {
		uf.union(c.start, c.end)
	}

	sort.Slice(uf.size, func(i, j int) bool {
		return uf.size[i] > uf.size[j]
	})

	m := make(map[int]int, len(junctionBoxes))
	for i := 0; i < len(junctionBoxes); i++ {
		m[uf.find(i)] += 1
	}

	circuits := make([]int, 0, len(m))
	for _, v := range m {
		circuits = append(circuits, v)
	}

	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i] > circuits[j]
	})

	s := 1
	for _, c := range circuits[:3] {
		s *= c
	}

	fmt.Printf("%d\n", s)
}

func newUnionFind(junctionBoxes []JunctionBox) *UnionFind {
	parents := make([]int, len(junctionBoxes))
	size := make([]int, len(junctionBoxes))
	for i := 0; i < len(junctionBoxes); i++ {
		parents[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parents,
		size:   size,
		count:  len(junctionBoxes),
	}
}

func (uf *UnionFind) find(x int) int {
	for x != uf.parent[x] {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX == rootY {
		return
	}

	if uf.size[rootX] > uf.size[rootY] {
		uf.parent[rootY] = rootX
		uf.size[rootX] += 1
	} else {
		uf.parent[rootX] = rootY
		uf.size[rootY] += 1
	}

	uf.count -= 1
}

func distance(p1, p2 JunctionBox) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) +
		math.Pow(float64(p2.y-p1.y), 2) +
		math.Pow(float64(p2.z-p1.z), 2))
}
