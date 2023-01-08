package main

import (
	"bufio"
	"fmt"
	. "image"

	"log"
	"os"
)

var grid = map[Point]bool{}
var part2 bool

func main() {
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		process(line, y)
		y++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	state := dcopy(grid)

	//part 1
	fmt.Println("part1")
	for i := 0; i < 10; i++ {
		state = step(state, i)
	}
	display(state)
	fmt.Println("part 2")
	//part 2
	state = dcopy(grid)
	part2 = true
	for i := 0; i >= 0; i++ {
		state = step(state, i)
	}

	display(state)
}

func exists(xs map[Point]bool, ps []Point) bool {
	for _, pt := range ps {
		if _, ok := xs[pt]; ok {
			return true
		}
	}
	return false
}

func dcopy(xs map[Point]bool) map[Point]bool {
	next := map[Point]bool{}
	for k, v := range xs {
		next[k] = v
	}
	return next
}

func step(xs map[Point]bool, count int) map[Point]bool {
	fixed := true
	next := dcopy(xs)
	dirs := [4]func(Point) []Point{north, south, west, east}
	change := [4]Point{Pt(0, -1), Pt(0, 1), Pt(-1, 0), Pt(1, 0)}
	proposed := map[Point]int{}
	proposer := map[Point]Point{}
	offset := count % 4
	for elf := range xs {
		surr := surrounding(elf)
		// only propose if an elf exists in your surround
		if exists(xs, surr) {
			for i := 0; i < 4; i++ {
				index := (offset + i) % 4
				if !exists(xs, dirs[index](elf)) {
					proposed[elf.Add(change[index])]++
					proposer[elf.Add(change[index])] = elf
					break
				}
			}
		}
	}

	toAdd := []Point{}
	toDelete := []Point{}
	for proposal, j := range proposed {
		if j == 1 {
			fixed = false
			toAdd = append(toAdd, proposal)
			toDelete = append(toDelete, proposer[proposal])
		}
	}
	for _, e := range toDelete {
		delete(next, e)
	}
	for _, e := range toAdd {
		_ = e
		next[e] = true
	}
	if fixed && part2 {
		fmt.Println("elves have stopped moving at", count+1)
		os.Exit(1)
	}

	return next
}

func surrounding(pt Point) []Point {
	all := []Point{}
	all = append(all, north(pt)...)
	all = append(all, south(pt)...)
	all = append(all, west(pt)...)
	all = append(all, east(pt)...)
	return all
}

func north(pt Point) []Point {
	n := Pt(pt.X, pt.Y-1)
	ne := Pt(pt.X+1, pt.Y-1)
	nw := Pt(pt.X-1, pt.Y-1)
	return []Point{n, ne, nw}
}

func south(pt Point) []Point {
	s := Pt(pt.X, pt.Y+1)
	se := Pt(pt.X+1, pt.Y+1)
	sw := Pt(pt.X-1, pt.Y+1)
	return []Point{s, se, sw}
}

func west(pt Point) []Point {
	w := Pt(pt.X-1, pt.Y)
	nw := Pt(pt.X-1, pt.Y-1)
	sw := Pt(pt.X-1, pt.Y+1)
	return []Point{w, nw, sw}
}

func east(pt Point) []Point {
	e := Pt(pt.X+1, pt.Y)
	ne := Pt(pt.X+1, pt.Y-1)
	se := Pt(pt.X+1, pt.Y+1)
	return []Point{e, ne, se}
}

func process(line string, y int) {
	for x, v := range line {
		if v == '#' {
			grid[Pt(x, y)] = true
		}
	}
}

func display(xs map[Point]bool) {
	empty := 0
	x1, y1, x2, y2 := minmax(xs)
	for y := y1; y < y2+1; y++ {
		for x := x1; x < x2+1; x++ {
			if xs[Pt(x, y)] {
				fmt.Printf("%c", '#')
				continue
			}
			fmt.Printf("%c", '.')
			empty++
		}
		fmt.Println()
	}
	fmt.Println("empty", empty)
}

func minmax(xs map[Point]bool) (x, y, x2, y2 int) {
	for k := range xs {
		x = min(x, k.X)
		x2 = max(x2, k.X)
		y = min(y, k.Y)
		y2 = max(y2, k.Y)
	}
	return x, y, x2, y2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
