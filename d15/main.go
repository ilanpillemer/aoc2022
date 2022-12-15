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
var minx = 0 //-2147483648
var miny = 0 //-2147483648 1050107
var seen = map[Point]bool{}
var sensors = map[Point]Pair{}
var beacons = map[Point]bool{}
var part1 bool
var limit = 4000000 + 1 //part1

type Pair struct {
	S Point
	B Point
}

func (p Pair) radius() int {
	diff := p.S.Sub(p.B)
	return abs(diff.X) + abs(diff.Y)
}

func (p Pair) visible(pt Point) bool {

	if part1 {
		if _, ok := sensors[pt]; ok {
			return false
		}

		if beacons[pt] {
			return false
		}
	}

	diff := p.S.Sub(pt)
	dist := abs(diff.X) + abs(diff.Y)
	return dist <= p.radius()
}

func main() {
	fmt.Println("d15")
	//part1 = true
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		s, p := process(line)
		sensors[s] = p
	}
	for s, p := range sensors {
		//fmt.Println(s, p, p.radius())
		maxx = max(maxx, s.X+p.radius())
		maxy = max(maxy, s.Y+p.radius())
		minx = min(minx, s.X-p.radius())
		miny = min(miny, s.Y-p.radius())
	}
	fmt.Printf("%d,%d --> %d,%d\n", minx, miny, maxx, maxy)
	if !part1 {
		scanBorders(grid)
	}
	// part1
	if part1 {
		scan2(grid, 10)
		scan2(grid, 2000000)
		fmt.Println("counting")
		fmt.Println("sample count", count(grid, 10))
		fmt.Println("input count", count(grid, 2000000))
	}

}

func count(xs map[Point]string, y int) int {
	c := 0
	for x := minx; x < maxx; x++ {
		tile := xs[Pt(x, y)]
		//fmt.Print(tile)
		if tile != "" {
			c++
		}
	}
	fmt.Println()
	return c
}

func scanBorders(xs map[Point]string) {
	borders := getBorders()
	fmt.Println("spots to scan:", len(borders))
	for pt := range borders {

		//fmt.Println(pt)
		found := false
		for _, sensor := range sensors {
			if sensor.visible(pt) {
				//xs[Pt(x, y)] = "#"
				//seen[pt] = true
				found = true
				break
			}
		}
		if !found {
			distress := pt
			fmt.Println(distress)
			fmt.Println("tuning = ", int64(distress.X)*4000000+int64(distress.Y))
			panic("distress")
		}

	}

}

func getBorders() map[Point]bool {
	borders := map[Point]bool{}
	add := func(xs map[Point]bool, pt Point) {
		if pt.X >= 0 && pt.X < limit && pt.Y >= 0 && pt.Y < limit {
			borders[pt] = true
		}
	}
	for i, sensor := range sensors {
		fmt.Printf("sensor %d -> %v\n", i, sensor)

		dist := sensor.radius() + 1
		for x := 0; x <= dist; x++ {
			a := sensor.S.Add(Pt(x, dist-x))
			b := sensor.S.Add(Pt(x, -(dist - x)))
			c := sensor.S.Add(Pt(-x, dist-x))
			d := sensor.S.Add(Pt(-x, -(dist - x)))
			add(borders,a)
			add(borders,b)
			add(borders,c)
			add(borders,d)
		}
	}
	return borders
}

// still too slow as millions is millions
func scanReduce(xs map[Point]string, atmost int) {
	for y := 0; y < atmost; y++ {
		for x := 0; x < atmost; x++ {
			pt := Pt(x, y)
			// is this point already been found?
			worker := func(x, y int) {
				found := false
				//if !seen[pt] {
				for _, sensor := range sensors {
					if sensor.visible(pt) {
						//xs[Pt(x, y)] = "#"
						//seen[pt] = true
						found = true
						break
					}
				}
				if !found {
					distress := Pt(x, y)
					fmt.Println(distress)
					fmt.Println("tuning = ", int64(distress.X)*4000000+int64(distress.Y))
					panic("distress")
				}
				//	seen[pt] = true
				//}

			}
			go worker(x, y)
		}
		println("scanned row:", y)
	}
}

func scan2(xs map[Point]string, y int) {
	for x := minx; x < maxx; x++ {
		pt := Pt(x, y)
		// is this point already been found?
		if !seen[pt] {
			for _, sensor := range sensors {
				if sensor.visible(pt) {
					xs[Pt(x, y)] = "#"
					seen[pt] = true
					break
				}
			}
			seen[pt] = true
		}
	}

}

func scan(xs map[Point]string) {
	for y := miny; y < maxy; y++ {
		for x := minx; x < maxx; x++ {
			pt := Pt(x, y)
			// is this point already been found?
			if !seen[pt] {
				for _, sensor := range sensors {
					if sensor.visible(pt) {
						xs[Pt(x, y)] = "#"
						seen[pt] = true
						break
					}
				}
				seen[pt] = true
			}
		}
	}
}

func find(xs map[Point]string, atmost int) {
	var distress Point
	for y := 0; y < atmost; y++ {
		for x := 0; x < atmost; x++ {
			tile := xs[Pt(x, y)]
			if tile == "" {
				tile = "."
				distress = Pt(x, y)
				fmt.Println(distress)
				fmt.Println("tuning = ", int64(distress.X)*4000000+int64(distress.Y))
				return
			}
		}
	}
}

func display(xs map[Point]string) {
	for y := miny; y < maxy; y++ {
		for x := minx; x < maxx; x++ {
			tile := xs[Pt(x, y)]
			if tile == "" {
				tile = "."
			}
			fmt.Print(tile)
		}
		fmt.Println()
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

func process(line string) (Point, Pair) {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	var x1, x2, y1, y2 int
	S := Point{}
	B := Point{}
	fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)
	S, B = Pt(x1, y1), Pt(x2, y2)
	P := Pair{S: S, B: B}

	maxx = max(max(maxx, x1), x2)
	maxy = max(max(maxy, y1), y2)
	minx = min(min(minx, x1), x2)
	miny = min(min(miny, y1), y2)
	beacons[Pt(x2, y2)] = true
	return S, P
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
