package main

import (
	"fmt"
	. "image"

	"os"
	"strconv"
	"strings"
)

type board map[Point]rune

var maxx, maxy int
var lookup = map[int][5]board{}
var walls = board{}

func main() {
	tiles := board{}
	up := board{}
	down := board{}
	left := board{}
	right := board{}
	empty := board{}
	fmt.Println("d22")
	b, _ := os.ReadFile("INPUT")
	grid := strings.Split(string(b), "\n")
	//parse to get all tiles
	for i, t := range grid {
		for j, v := range t {
			if v != ' ' {
				maxx = max(maxx, j)
				maxy = max(maxy, i)
				e := v
				tiles[Pt(j, i)] = e
				switch e {
				case '^':
					up[Pt(j, i)] = e
				case 'v':
					down[Pt(j, i)] = e
				case '>':
					right[Pt(j, i)] = e
				case '<':
					left[Pt(j, i)] = e
				case '.':
					empty[Pt(j, i)] = 'o'
				case '#':
					walls[Pt(j, i)] = e
				default:
					panic("oy")
				}
			}
		}
	}
	//parse to link tiles
	fmt.Println(maxx, maxy)
	display(tiles)
	lookup[0] = [5]board{up, down, left, right, empty}
	fmt.Println(maxx, maxy)
	exit = Pt(maxx-1, maxy)
	score := solve(acc{Pt(1, 0), 0})
	seen[fmt.Sprintf("%v%v", Pt(1, 0), empty)] = true
	fmt.Println("SCORE", score)
}

var seen = map[string]bool{}
var seen2 = map[string]bool{}
var part1 = false

func getLookup(t int) [5]board {
	if l, ok := lookup[t]; ok {
		return l
	}

	state := next(lookup[t-1])
	lookup[t] = state
	return lookup[t]

}

type acc struct {
	P Point
	S int
}

var exit = Pt(maxx-1, maxy)
var count = 0
var maxminute = -1

func solve(x acc) int {
	Q := []acc{x}
	for {
		p := Q[0]
		Q = Q[1:]
		if p.P == exit && part1 {
			fmt.Println("found", p.P, p.S)
			return p.S
		}
		if p.P == exit && count == 2 {
			fmt.Println("found", p.P, p.S)
			return p.S
		} else if p.P == exit && count == 0 {
			exit = Pt(1, 0)
			count++
			fmt.Println("one", p.P, p.S)
			seen = map[string]bool{}
			seen2 = map[string]bool{}
			Q = []acc{}
		} else if p.P == exit && count == 1 {
			exit = Pt(maxx-1, maxy)
			count++
			fmt.Println("two", p.P, p.S)
			seen = map[string]bool{}
			seen2 = map[string]bool{}
			Q = []acc{}
		}

		xs := moves(acc{p.P, p.S + 1})
		Q = append(Q, xs...)
	}

}

func moves(ac acc) []acc {
	empty := getLookup(ac.S)[4]

	x := ac.P
	xs := []acc{}
	a := x.Add(Pt(1, 0))
	b := x.Add(Pt(-1, 0))
	c := x.Add(Pt(0, 1))
	d := x.Add(Pt(0, -1))
	e := x
	options := []Point{a, b, c, d, e}
	for _, y := range options {
		if y == exit {
			fmt.Println(ac.S, "!!!!", y)
		}
	}
	for _, y := range options {
		if _, ok := empty[y]; ok {
			key := fmt.Sprintf("%v%v", y, empty)
			key2 := fmt.Sprintf("%v%v", y, ac.S)
			if !seen[key] && !seen2[key2] {
				xs = append(xs, acc{y, ac.S})
			}
			seen[key] = true
			seen2[key2] = true
		}
	}

	return xs
}

func next(b [5]board) [5]board {
	pt := map[rune]Point{'^': Pt(0, -1), 'v': Pt(0, 1), '>': Pt(1, 0), '<': Pt(-1, 0)}
	f := func(x board, d rune) board {
		y := board{}
		for k, v := range x {
			if v == d {
				w := k.Add(pt[d])
				if w.Y == 0 {
					w.Y = maxy - 1
				}
				if w.Y == maxy {
					w.Y = 1
				}
				if w.X == 0 {
					w.X = maxx - 1
				}
				if w.X == maxx {
					w.X = 1
				}
				y[w] = v
			}
		}
		return y
	}
	up := f(b[0], '^')
	down := f(b[1], 'v')
	left := f(b[2], '<')
	right := f(b[3], '>')
	empty := empties(up, down, left, right)
	z := [5]board{up, down, left, right, empty}
	return z
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func atoi(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}

func display(xs board) {
	for y := 0; y < maxy+1; y++ {
		for x := 0; x < maxx+1; x++ {
			if e, ok := xs[Pt(x, y)]; ok {
				fmt.Printf("%c", e)
				continue
			}
			fmt.Printf("%c", '.')
		}
		fmt.Println()
	}
}

func empties(a, b, c, d board) board {
	e := board{}
	for y := 0; y < maxy+1; y++ {
		for x := 0; x < maxx+1; x++ {
			_, i := a[Pt(x, y)]
			_, j := b[Pt(x, y)]
			_, k := c[Pt(x, y)]
			_, l := d[Pt(x, y)]
			_, w := walls[Pt(x, y)]
			if i || j || k || l || w {
				continue
			}
			e[Pt(x, y)] = 'o'
		}
	}
	return e
}
