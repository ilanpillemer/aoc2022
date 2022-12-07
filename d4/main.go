package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sections struct {
	lhs int
	rhs int
}

func main() {
	fmt.Println("d4")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	count2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		lhs, rhs := pair(line)
		fullyContained := contains(lhs, rhs) || contains(rhs, lhs)
		overlapped := overlap(lhs, rhs) || overlap(rhs, lhs)
		if fullyContained {
			count++
		}
		if overlapped {
			count2++
		}
	}
	fmt.Println("part1: ", count)
	fmt.Println("part2: ", count2)
}

func pair(x string) (lhs, rhs sections) {
	xs := strings.Split(x, ",")
	return intPair(xs[0]), intPair(xs[1])
}

func intPair(xs string) sections {
	s := sections{}
	ys := strings.Split(xs, "-")
	s.lhs, _ = strconv.Atoi(string(ys[0]))
	s.rhs, _ = strconv.Atoi(string(ys[1]))
	return s
}

func contains(lhs, rhs sections) bool {
	return lhs.lhs <= rhs.lhs && lhs.rhs >= rhs.rhs
}

func overlap(lhs, rhs sections) bool {
	return lhs.lhs <= rhs.rhs && lhs.rhs >= rhs.lhs
}
