package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pair struct {
	x int
	y int
}

var maxx int
var maxy int
var minx int
var miny int

var hmap = map[Pair]int{}
var tmap = map[Pair]int{}

var currHead = Pair{0, 0}
var curr1 = Pair{0, 0}
var curr2 = Pair{0, 0}
var curr3 = Pair{0, 0}
var curr4 = Pair{0, 0}
var curr5 = Pair{0, 0}
var curr6 = Pair{0, 0}
var curr7 = Pair{0, 0}
var curr8 = Pair{0, 0}
var currTail = Pair{0, 0}

func main() {
	tmap[Pair{currTail.x, currTail.y}]++
	fmt.Println("d9")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		process2(line)
	}
	fmt.Println(solve(tmap))
}

func solve(tmap map[Pair]int) int {
	count := 0
	for _, e := range tmap {
		if e > 0 {
			count++
		}
	}
	return count
}

func process2(line string) {
	var d string
	var steps int
	fmt.Sscanf(line, "%s %d", &d, &steps)
	for i := 0; i < steps; i++ {
		moveHead(d)
		curr1 = moveTail(currHead, curr1)
		curr2 = moveTail(curr1, curr2)
		curr3 = moveTail(curr2, curr3)
		curr4 = moveTail(curr3, curr4)
		curr5 = moveTail(curr4, curr5)
		curr6 = moveTail(curr5, curr6)
		curr7 = moveTail(curr6, curr7)
		curr8 = moveTail(curr7, curr8)
		currTail = moveTail(curr8, currTail)
		tmap[currTail]++
	}

}

func process(line string) {
	var d string
	var steps int
	fmt.Sscanf(line, "%s %d", &d, &steps)
	for i := 0; i < steps; i++ {
		moveHead(d)
		currTail = moveTail(currHead, currTail)
		newT := Pair{currTail.x, currTail.y}
		tmap[newT]++
	}
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func moveTail(knot Pair, tail Pair) Pair {
	dx := (knot.x - tail.x)
	dy := (knot.y - tail.y)

	switch {
	case dx == 0 && dy == 2:
		tail.y++
		return tail
	case dx == 0 && dy == -2:
		tail.y--
		return tail
	}

	switch {
	case dy == 0 && dx == 2:
		tail.x++
		return tail
	case dy == 0 && dx == -2:
		tail.x--
		return tail
	}

	switch {
	case dx == 2:
		tail.x++
		tail.y = tail.y + sign(dy)
	case dy == 2:
		tail.y++
		tail.x = tail.x + sign(dx)
	case dx == -2:
		tail.x--
		tail.y = tail.y + sign(dy)
	case dy == -2:
		tail.y--
		tail.x = tail.x + sign(dx)
	default:
	}
	return tail
}

func moveHead(d string) {
	switch d {
	case "U":
		currHead.y = currHead.y + 1
	case "D":
		currHead.y = currHead.y - 1
	case "R":
		currHead.x = currHead.x + 1
	case "L":
		currHead.x = currHead.x - 1
	default:
		panic("oh no")
	}
	newH := Pair{currHead.x, currHead.y}
	hmap[newH]++
	if newH.x > maxx {
		maxx = newH.x
	}
	if newH.y > maxy {
		maxy = newH.y
	}
	if newH.x < minx {
		minx = newH.x
	}
	if newH.y < miny {
		miny = newH.y
	}
}

func displayTH(x map[Pair]int) {
	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			if (Pair{x, y} == currHead) {
				fmt.Print("H")
				continue
			}
			if (Pair{x, y} == currTail) {
				fmt.Print("T")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func displayTHLong(x map[Pair]int) {
	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			if (Pair{x, y} == currHead) {
				fmt.Print("H")
				continue
			}
			if (Pair{x, y} == curr1) {
				fmt.Print("1")
				continue
			}
			if (Pair{x, y} == curr2) {
				fmt.Print("2")
				continue
			}
			if (Pair{x, y} == curr3) {
				fmt.Print("3")
				continue
			}
			if (Pair{x, y} == curr4) {
				fmt.Print("4")
				continue
			}
			if (Pair{x, y} == curr5) {
				fmt.Print("5")
				continue
			}
			if (Pair{x, y} == curr6) {
				fmt.Print("6")
				continue
			}
			if (Pair{x, y} == curr7) {
				fmt.Print("7")
				continue
			}
			if (Pair{x, y} == curr8) {
				fmt.Print("8")
				continue
			}
			if (Pair{x, y} == currTail) {
				fmt.Print("9")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func display(xs map[Pair]int) {
	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			if xs[Pair{x, y}] != 0 {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}
