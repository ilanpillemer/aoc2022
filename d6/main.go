package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("d6")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)-4; i++ {
			a, b := group(i, 4)
			if uniq(line[a:b]) {
				fmt.Println("p1: ", i+4)
				break
			}
		}
		for i := 0; i < len(line)-14; i++ {
			a, b := group(i, 14)
			if uniq(line[a:b]) {
				fmt.Println("p2: ", i+14)
				break
			}
		}
	}
}

func group(n, dist int) (a, b int) {
	return n, n + dist
}

func uniq(x string) bool {
	set := map[rune]int{}
	for _, b := range x {
		set[b]++
		if set[b] > 1 {
			return false
		}
	}
	return true
}
