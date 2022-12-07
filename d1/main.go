package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("INPUT") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	totals := []int{}
	str := string(b) // convert content to a 'string'
	elves := strings.Split(str, "\n\n")
	for _, elf := range elves {

		totals = append(totals, sum(elf))

	}
	sort.Ints(totals)
	leaders := sumInt(totals[len(totals)-3:])
	fmt.Println(leaders)

}

func sumInt(xs []int) int {
	total := 0
	for _, x := range xs {
		total = total + x
	}
	return total

}

func sum(elf string) int {
	xs := strings.Fields(elf)
	total := 0
	for _, x := range xs {
		i, _ := strconv.Atoi(x)
		total = total + i
	}
	return total
}

func main2() {
	fmt.Println("d1")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
