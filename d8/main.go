package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pair struct {
	x int
	y int
}

var m = map[Pair]int{}
var counts = map[Pair]int{}
var upcounts = map[Pair]int{}
var downcounts = map[Pair]int{}
var leftcounts = map[Pair]int{}
var rightcounts = map[Pair]int{}
var sceniccounts = map[Pair]int{}
var maxx int
var maxy int

func main() {
	fmt.Println("d8")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if maxx == 0 {
			maxx = len(line)
		}
		process(line, y)
		y++
		maxy++
	}
	top()
	bottom()
	left()
	right()
	fmt.Println("part1: ", visible(counts))
	lookaround(m)
	calc(m)
	fmt.Println("part2: ", max(sceniccounts))
}

func process(line string, y int) {
	for x, v := range line {
		n, _ := strconv.Atoi(string(v))
		m[Pair{x, y}] = n
		counts[Pair{x, y}] = 0
	}
}

func top() {
	for x := 0; x < maxx; x++ {
		view := m[Pair{x, 0}]
		counts[Pair{x, 0}]++
		for y := 1; y < maxy; y++ {
			next := m[Pair{x, y}]
			if next > view {
				counts[Pair{x, y}]++
				view = next
				continue
			}
		}
	}
}

func bottom() {
	for x := 0; x < maxx; x++ {
		view := m[Pair{x, maxy - 1}]
		counts[Pair{x, maxy - 1}]++
		for y := maxy - 2; y >= 0; y-- {
			next := m[Pair{x, y}]
			if next > view {
				counts[Pair{x, y}]++
				view = next
				continue
			}
		}
	}
}

func left() {
	for y := 0; y < maxy; y++ {
		view := m[Pair{0, y}]
		counts[Pair{0, y}]++
		for x := 1; x < maxx; x++ {
			next := m[Pair{x, y}]
			if next > view {
				counts[Pair{x, y}]++
				view = next
				continue
			}
		}
	}
}

func right() {
	for y := 0; y < maxy; y++ {
		view := m[Pair{maxy - 1, y}]
		counts[Pair{maxy - 1, y}]++
		for x := maxx - 2; x >= 0; x-- {
			next := m[Pair{x, y}]
			if next > view {
				counts[Pair{x, y}]++
				view = next
				continue
			}
		}
	}
}

func lookup(pos Pair) {
	cx := pos.x
	cy := pos.y
	view := m[Pair{cx, cy}]
	for y := cy - 1; y >= 0; y-- {
		next := m[Pair{cx, y}]
		if next < view {
			upcounts[Pair{cx, cy}]++
			continue
		}
		upcounts[Pair{cx, cy}]++
		break
	}

}

func lookdown(pos Pair) {
	cx := pos.x
	cy := pos.y
	view := m[Pair{cx, cy}]

	for y := cy + 1; y < maxy; y++ {
		next := m[Pair{cx, y}]
		if next < view {
			downcounts[Pair{cx, cy}]++
			continue
		}
		downcounts[Pair{cx, cy}]++
		break
	}

}

func lookright(pos Pair) {
	cx := pos.x
	cy := pos.y
	view := m[Pair{cx, cy}]

	for x := cx + 1; x < maxx; x++ {
		next := m[Pair{x, cy}]
		if next < view {
			rightcounts[Pair{cx, cy}]++
			continue
		}
		rightcounts[Pair{cx, cy}]++
		break
	}

}

func lookleft(pos Pair) {
	cx := pos.x
	cy := pos.y
	view := m[Pair{cx, cy}]

	for x := cx - 1; x >= 0; x-- {
		next := m[Pair{x, cy}]
		if next < view {
			leftcounts[Pair{cx, cy}]++
			continue
		}
		leftcounts[Pair{cx, cy}]++
		break
	}

}

func display(z map[Pair]int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			next := z[Pair{x, y}]
			fmt.Print(next)
		}
		fmt.Println()
	}

}

func lookaround(z map[Pair]int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			lookup(Pair{x, y})
			lookdown(Pair{x, y})
			lookright(Pair{x, y})
			lookleft(Pair{x, y})
		}
	}

}

func calc(z map[Pair]int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			a := upcounts[Pair{x, y}]
			b := downcounts[Pair{x, y}]
			c := leftcounts[Pair{x, y}]
			d := rightcounts[Pair{x, y}]
			sceniccounts[Pair{x, y}] = a * b * c * d
		}
	}

}

func max(z map[Pair]int) int {
	v := -1
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			next := z[Pair{x, y}]
			if next > v {
				v = next
			}
		}
	}
	return v
}

func visible(z map[Pair]int) int {
	count := 0
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			next := z[Pair{x, y}]
			if next > 0 {
				count++
			}
		}

	}
	return count
}
