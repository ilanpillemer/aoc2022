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

var drops = map[Point]bool{}

func main() {
	f, _ := os.Open("INPUT")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		xyz := strings.Split(line, ",")
		x, y, z := atoi(xyz[0]), atoi(xyz[1]), atoi(xyz[2])
		drops[Point{x, y, z}] = true
	}
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
