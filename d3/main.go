package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("d3")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	threesome := []string{}
	total := 0
	total2 := 0
	count := 0
	for scanner.Scan() {
		count = count + 1
		line := scanner.Text()
		total = total + priority(common(divide(line)))
		threesome = append(threesome, line)
		if count%3 == 0 {
			total2 = total2 + priority(common3(threesome))
			threesome = []string{}
		}
	}
	fmt.Println(total)
	fmt.Println(total2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func divide(x string) (lhs, rhs string) {
	l := len(x)
	lhs, rhs = x[:l/2], x[l/2:]
	return
}

func common3(xs []string) (r rune) {
	for _, r = range xs[0] {
		if strings.ContainsRune(xs[1], r) && strings.ContainsRune(xs[2], r) {
			return
		}
	}
	panic("nothing in common, your code sucks")

}

func common(lhs, rhs string) (r rune) {
	for _, r = range lhs {
		if strings.ContainsRune(rhs, r) {
			return
		}
	}
	panic("nothing in common, your code sucks")
}

func priority(r rune) int {
	switch {
	case r >= 'A' && r <= 'Z':
		return int(r) - 'A' + 27
	case r >= 'a' && r <= 'z':
		return int(r) - 'a' + 1

	default:
		panic("your code really sucks")
	}
}
