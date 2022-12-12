package main

import (
	"bufio"
	"fmt"
	. "image"
	"log"
	"os"
	"sort"
)

var grid = map[Point]rune{}
var maxx, maxy int
var S acc
var E Point
var XS = []acc{}
var seen = map[Point]bool{}

func main() {
	fmt.Println("d12")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		maxx = len(line)
		process(line, y)
		y++
		if y > maxy {
			maxy = y
		}
	}
	fmt.Println(S)
	fmt.Printf("part 2 has %d choices\n", len(XS))
	fmt.Println(E)
	//display(grid)
	//	steps := solve(grid)
	//fmt.Println("steps: ", steps)

	scores := []int{}
	for _, next := range XS {
		seen = map[Point]bool{}
		fmt.Println("solving for ", next)
		S = next
		seen[S.x] = true
		steps := solve(grid)
		scores = append(scores, steps)
		fmt.Println("steps: ", steps)
		//display(grid)
	}
	sort.Ints(scores)
	fmt.Printf("Most scenic score %d\n", scores[0])
}

type acc struct {
	x Point
	c int
}

func solve(xs map[Point]rune) int {
	queue := []acc{}
	for {
		if S.x == E {
			break
		}
		queue = append(queue, getMoves(S)...)
		S = queue[0]
		queue = queue[1:]

	}

	return S.c
}

func getMoves(e acc) []acc {
	moves := []acc{}
	up := e.x.Add(Pt(0, 1))
	down := e.x.Add(Pt(0, -1))
	left := e.x.Add(Pt(-1, 0))
	right := e.x.Add(Pt(1, 0))
	if valid(e.x, up) {
		moves = append(moves, acc{up, e.c + 1})
		seen[up] = true
	}
	if valid(e.x, down) {
		moves = append(moves, acc{down, e.c + 1})
		seen[down] = true
	}
	if valid(e.x, left) {
		moves = append(moves, acc{left, e.c + 1})
		seen[left] = true
	}
	if valid(e.x, right) {
		moves = append(moves, acc{right, e.c + 1})
		seen[right] = true
	}
	return moves
}

func valid(e Point, f Point) bool {
	if seen[f] {
		return false
	}
	if f.X < 0 || f.Y < 0 {
		return false
	}
	a := grid[e]
	b := grid[f]
	if b-1 <= a {
		return true
	}
	return false
}

func process(line string, y int) {
	for x, v := range line {
		if v == 'S' {
			grid[Pt(x, y)] = 'a'
			S = acc{Pt(x, y), 0}
			XS = append(XS, acc{Pt(x, y), 0})
			continue
		}
		if v == 'E' {
			grid[Pt(x, y)] = 'z'
			E = Pt(x, y)
			continue
		}
		grid[Pt(x, y)] = v
		if grid[Pt(x, y)] == 'a' {
			XS = append(XS, acc{Pt(x, y), 0})
		}
	}
}

func display(xs map[Point]rune) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			v := grid[Pt(x, y)]
			if S.x == Pt(x, y) {
				fmt.Printf("%c", 'S')
				continue
			}
			if E == Pt(x, y) {
				fmt.Printf("%c", 'E')
				continue
			}
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}
