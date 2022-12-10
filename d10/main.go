package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var x = 1
var y = 0
var cycle = 1
var offset = 0
var store = map[int]int{}

func main() {
	fmt.Println("d10")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		process(line)
	}
	fmt.Printf("There are %d cycles\n", cycle)
	fmt.Println("p1 :", solve1(store))
	solve2()

}

func solve1(reg map[int]int) int {
	sum := (20 * reg[20]) +
		(60 * reg[60]) +
		(100 * reg[100]) +
		(140 * reg[140]) +
		(180 * reg[180]) +
		(220 * reg[220])
	return sum
}

func solve2() {
	for i := 0; i < 240; i++ {
		if newline(i) {
			fmt.Println(" ", offset)
		}
		if pixel(i) {
			fmt.Print("#")
			continue
		}
		fmt.Printf("%v", " ")
	}
	fmt.Println(" ", offset)
}

func pixel(x int) bool {
	mid := store[x+1]
	x = x - offset
	if x == mid-1 || x == mid || x == mid+1 {
		return true
	}
	return false
}

func newline(x int) bool {
	switch x {
	case 40, 80, 120, 160, 200:
		offset = offset + 40
		return true
	}
	return false
}

func process(line string) {
	var instr string
	var par int
	fmt.Sscanf(line, "%s %d", &instr, &par)

	switch instr {
	case "noop":
		store[cycle] = x
		cycle++
	case "addx":
		store[cycle] = x
		cycle++
		store[cycle] = x
		cycle++
		x = x + par
		store[cycle] = x
	}
}
