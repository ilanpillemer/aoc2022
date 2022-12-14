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
	state := [5]board{up, down, left, right, empty}
	for i := 0; i < (maxx*maxy)+1; i++ {
		lookup[i] = state
		state = next(state)
	}

	//	for i := 0; i < len(lookup); i++ {
	//		fmt.Println(i)
	//		display(lookup[i][4])
	//	}

	score := solve(acc{Pt(1, 0), 0})
	seen[fmt.Sprintf("%v%v", Pt(1, 0), empty)] = true
	fmt.Println("SCORE", score+1)
}

var seen = map[string]bool{}
var seen2 = map[string]bool{}

type acc struct {
	P Point
	S int
}

var exit = Pt(maxx-1, maxy)
var maxminute = -1

func solve(x acc) int {
	Q := []acc{x}
	for {
		p := Q[0]
//		if p.S > maxminute {
//			fmt.Println(p.S, p.P)
//			maxminute = p.S
//		}


		Q = Q[1:]
		if p.P == Pt(maxx-1, maxy) {
			return p.S
		}
		xs := moves(acc{p.P, p.S + 1})
		Q = append(Q, xs...)
	}

}

func moves(ac acc) []acc {
	z := ac.S % (maxx * maxy)
	empty := lookup[z][4]

	x := ac.P
	xs := []acc{}
	key := fmt.Sprintf("%v%v", x, empty)
	key2 := fmt.Sprintf("%v%v", x, ac.S)
	if !seen[key] && !seen2[key2] {
		xs = append(xs, acc{x, ac.S})
		seen[key] = true
		seen2[key2] = true
	}
	a := x.Add(Pt(1, 0))
	b := x.Add(Pt(-1, 0))
	c := x.Add(Pt(0, 1))
	d := x.Add(Pt(0, -1))
	options := []Point{a, b, c, d}
	if a == exit || b == exit || c == exit || d == exit {
		fmt.Println(ac.S, "!!!!", a)
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
