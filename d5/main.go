package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type stack []string

func main() {
	fmt.Println("d5")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stacks := fixture1()
	fmt.Println(stacks)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "move") {
			fmt.Println(line)
			move2(stacks, line)
		}
	}
	fmt.Println(stacks)

	for i := range stacks {
		x := pop(stacks, i)
		fmt.Print(x)
	}
	fmt.Println()

}

//    [D]
//[N] [C]
//[Z] [M] [P]
// 1   2   3

func fixture0() []stack {
	stacks := []stack{}
	stackx := []string{}
	stacks = append(stacks, stackx)

	stackx = []string{"Z", "N"}
	stacks = append(stacks, stackx)

	stackx = []string{"M", "C", "D"}
	stacks = append(stacks, stackx)

	stackx = []string{"P"}
	stacks = append(stacks, stackx)
	return stacks
}

func fixture1() []stack {
	stacks := []stack{}
	stackx := []string{}
	stacks = append(stacks, stackx)

	stackx = []string{"F", "D", "B", "Z", "T", "J", "R", "N"}
	stacks = append(stacks, stackx)

	stackx = []string{"R", "S", "N", "J", "H"}
	stacks = append(stacks, stackx)

	stackx = []string{"C", "R", "N", "J", "G", "Z", "F", "Q"}
	stacks = append(stacks, stackx)

	stackx = []string{"F", "V", "N", "G", "R", "T", "Q"}
	stacks = append(stacks, stackx)

	stackx = []string{"L", "T", "Q", "F"}
	stacks = append(stacks, stackx)

	stackx = []string{"Q", "C", "W", "Z", "B", "R", "G", "N"}
	stacks = append(stacks, stackx)

	stackx = []string{"F", "C", "L", "S", "N", "H", "M"}
	stacks = append(stacks, stackx)

	stackx = []string{"D", "N", "Q", "M", "T", "J"}
	stacks = append(stacks, stackx)

	stackx = []string{"P", "G", "S"}
	stacks = append(stacks, stackx)
	return stacks
}

func move(xs []stack, mv string) []stack {
	//move 1 from 2 to 1
	var amount int
	var from int
	var dest int
	fmt.Sscanf(mv, "move %d from %d to %d", &amount, &from, &dest)
	for i := 0; i < amount; i++ {
		x := pop(xs, from)
		push(xs, dest, x)
	}
	return xs
}

func move2(xs []stack, mv string) []stack {
	//move 1 from 2 to 1
	var amount int
	var from int
	var dest int
	fmt.Sscanf(mv, "move %d from %d to %d", &amount, &from, &dest)
	items := []string{}
	for i := 0; i < amount; i++ {
		x := pop(xs, from)
		items = append(items, x)

	}
	l := len(items) - 1

	for x := l; x > -1; x-- {
		push(xs, dest, items[x])
	}

	return xs
}

func pop(stack []stack, pos int) string {
	n := len(stack[pos]) - 1
	if n < 0 {
		return ""
	}
	x := stack[pos][n]
	stack[pos] = stack[pos][:n]
	return x
}

func push(stack []stack, pos int, x string) {
	stack[pos] = append(stack[pos], x)
}
