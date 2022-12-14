package main

import (
	"bufio"
	"fmt"
	. "image"
	"log"
	"os"
	"strconv"
	"strings"
)

var grid = map[Point]string{}
var maxx = 0
var maxy = 0
var minx = 1000000
var entry = Pt(500, 0)

func main() {
	fmt.Println("d14")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		process(line)
	}
	display(grid)
	fmt.Println()
	i := 0

	// part 2

	for {
		i++
		flow := pour(grid, entry)
		fmt.Printf("after %d units of sand..\n", i)
		//display(grid)
		fmt.Println()
		//		if !flow {
		//			fmt.Println("solution to part 1:", i-1)
		//			break
		//		}

		if !flow {
			fmt.Println("solution to part 2:", i)
			display(grid)
			break
		}
	}
}

func pour(xs map[Point]string, pt Point) bool {

	// part 1
	if pt.X < minx || pt.X > maxx || pt.Y > maxy {
		//	return false
	}

	//fmt.Println("trying ", pt)
	down := pt.Add(Pt(0, 1))
	left := pt.Add(Pt(-1, 1))
	right := pt.Add(Pt(1, 1))

	switch {
	case grid[down] == "" && pt.Y < maxy+1:
		return pour(xs, down)
	case grid[left] == "" && pt.Y < maxy+1:
		return pour(xs, left)
	case grid[right] == "" && pt.Y < maxy+1:
		return pour(xs, right)
	default:
		grid[pt] = "o"
		if pt == entry {
			return false
		}
		return true
	}

}

func display(xs map[Point]string) {
	for y := 0; y < maxy+3; y++ {
		for x := minx - 20; x < maxx+1+20; x++ {
			tile := xs[Pt(x, y)]
			if tile == "" {
				tile = "."
			}
			fmt.Print(tile)
		}
		fmt.Println()
	}

}

func process(line string) {
	xs := strings.Split(line, "->")
	s := Pt(-1, -1)
	for _, ys := range xs {
		zs := strings.Split(ys, ",")
		x, y := atoi(zs[0]), atoi(zs[1])
		if x > maxx {
			maxx = x
		}
		if minx > x {
			minx = x
		}
		if y > maxy {
			maxy = y
		}
		fmt.Println(x, y)
		if s == Pt(-1, -1) {
			s = Pt(x, y)
			continue
		}
		draw(s, Pt(x, y))
		s = Pt(x, y)
	}

}

func atoi(x string) int {
	x = strings.TrimSpace(x)
	y, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return y
}

func draw(a, b Point) {
	switch {
	case a.X == b.X:
		x := a.X
		dy := (b.Y - a.Y)
		for y := a.Y; y != b.Y+sign(dy); y = y + sign(dy) {
			fmt.Print(x, y, " ")
			grid[Pt(x, y)] = "#"
		}
	case a.Y == b.Y:
		y := a.Y
		dx := (b.X - a.X)
		for x := a.X; x != b.X+sign(dx); x = x + sign(dx) {
			fmt.Print(x, y, " ")
			grid[Pt(x, y)] = "#"
		}
	default:
		msg := fmt.Sprintf("not a perpendicular line [%v] -> [%v]", a, b)
		panic(msg)
	}
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
