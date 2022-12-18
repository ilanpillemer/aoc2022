package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

var maxx, maxy, maxz int
var minx, miny, minz int

var drops = map[Point]bool{}
var part1 = false

func main() {
	minx = 1000
	miny = 1000
	minz = 1000
	f, _ := os.Open("INPUT")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		xyz := strings.Split(line, ",")
		x, y, z := atoi(xyz[0]), atoi(xyz[1]), atoi(xyz[2])
		drops[Point{x, y, z}] = true
	}

	//part1 = true
	if part1 {
		solvePart1()
		os.Exit(0)
	}
	for v, _ := range drops {
		x, y, z := v.x, v.y, v.z
		maxx = max(maxx, x)
		maxy = max(maxx, y)
		maxz = max(maxx, z)
		minx = min(minx, x)
		miny = min(minx, y)
		minz = min(minx, z)
	}
	total := 0
	for v, _ := range drops {
		x, y, z := v.x, v.y, v.z

		if exposed(Point{x + 1, y, z}) {
			total++
		}
		if exposed(Point{x, y + 1, z}) {
			total++
		}
		if exposed(Point{x, y, z + 1}) {
			total++
		}
		if exposed(Point{x - 1, y, z}) {
			total++
		}
		if exposed(Point{x, y - 1, z}) {
			total++
		}
		if exposed(Point{x, y, z - 1}) {
			total++
		}
	}
	fmt.Println("part1: ", total)

}

func display() {

}

func isoutside(pt Point) bool {
	return pt.x > maxx ||
		pt.y > maxy ||
		pt.z > maxz ||
		pt.x < minx ||
		pt.y < miny ||
		pt.z < minz
}

var memo = map[Point]bool{}

func exposed(pt Point) bool {
	if result, ok := memo[pt]; ok {
		return result
	}
	Q := []Point{pt}
	S := map[Point]bool{}
	for len(Q) > 0 {
		next := Q[0]
		Q = Q[1:]
		if drops[next] {
			continue
		}
		if S[next] {
			continue
		}
		S[next] = true
		if isoutside(next) {
			for pt := range S {
				memo[pt] = true
			}
			return true
		}
		Q = append(Q, Point{next.x + 1, next.y, next.z})
		Q = append(Q, Point{next.x - 1, next.y, next.z})
		Q = append(Q, Point{next.x, next.y + 1, next.z})
		Q = append(Q, Point{next.x, next.y - 1, next.z})
		Q = append(Q, Point{next.x, next.y, next.z + 1})
		Q = append(Q, Point{next.x, next.y, next.z - 1})
	}
	for pt := range S {
		memo[pt] = false
	}
	memo[pt] = false
	return false
}

func solvePart1() {
	total := 0
	for v, _ := range drops {
		x, y, z := v.x, v.y, v.z

		if drops[Point{x + 1, y, z}] == false {
			total++
		}
		if drops[Point{x, y + 1, z}] == false {
			total++
		}
		if drops[Point{x, y, z + 1}] == false {
			total++
		}
		if drops[Point{x - 1, y, z}] == false {
			total++
		}
		if drops[Point{x, y - 1, z}] == false {
			total++
		}
		if drops[Point{x, y, z - 1}] == false {
			total++
		}
	}
	fmt.Println("part1: ", total)
}

func atoi(str string) int {
	str = strings.TrimSpace(str)
	i, _ := strconv.Atoi(str)
	return i
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
