package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

const (
	win  = 6
	lose = 0
	draw = 3
)

func main() {
	fmt.Println("d2")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	total := 0
	total2 := 0
	for scanner.Scan() {
		moves := strings.Fields(scanner.Text())
		total = total + score(moves[0], moves[1])

		move := mymove(moves[0], moves[1])
		total2 = total2 + score(moves[0], move)
	}
	println(total)
	println(total2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func score(lhs string, rhs string) int {
	return shape(rhs) + outcome(lhs, rhs)
}

func shape(x string) int {
	shapes := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	return shapes[x]
}

func outcome(lhs string, rhs string) int {
	outcomes := map[string]int{
		"AX": draw, //rock rock
		"AY": win,  //rock paper
		"AZ": lose, //rock scissors
		//
		"BX": lose, //paper rock
		"BY": draw, // paper paper
		"BZ": win,  //paper scissors
		//
		"CX": win,  //scissors rock
		"CY": lose, //scissors paper
		"CZ": draw, //scissors scissors
	}
	return outcomes[lhs+rhs]
}

func mymove(lhs string, rhs string) string { // X to lose, Y to draw, Z to win
	mymoves := map[string]string{
		"AX": "Z", //rock to lose --> scissors
		"AY": "X", //rock to draw --> rock
		"AZ": "Y", //rock to win --> paper
		//
		"BX": "X", //paper to lose --> rock
		"BY": "Y", // paper to draw --> paper
		"BZ": "Z", //paper to win -->
		//
		"CX": "Y", //scissors to lose -->
		"CY": "Z", //scissors to draw --> scissors
		"CZ": "X", //scissors to win --> rock
	}
	return mymoves[lhs+rhs]
}
