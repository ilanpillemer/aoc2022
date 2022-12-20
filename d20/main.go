package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.ReadFile("INPUT")
	lines := strings.Fields(string(file))
	//fmt.Println(lines)
	values := []int{}
	pointers := []*list.Element{}
	for i := 0; i < len(lines); i++ {
		values = append(values, atoi(lines[i]))
	}
	r := list.New()

	for i := 0; i < len(values); i++ {
		e := r.PushBack(values[i])
		pointers = append(pointers, e)
	}

	// Iterate through list and print its contents.
	for e := r.Front(); e != nil; e = e.Next() {
		//	fmt.Print(e.Value, ", ")
	}
	//fmt.Println()
	for i := 0; i < len(lines); i++ {
		v := pointers[i].Value.(int)
		e := pointers[i]
		//fmt.Printf("About to move %d %d spaces\n", v, v)
		if v < 0 {
			next := e
			for j := 0; j < abs(v); j++ {

				next = next.Prev()
				if next == e {
					next = next.Prev()
				}
				if next == nil {
					next = r.Back()
				}
			}
			r.MoveBefore(e, next)

		} else {
			next := e
			for j := 0; j < abs(v); j++ {

				next = next.Next()
				if next == e {
					next = next.Next()
				}
				if next == nil {
					next = r.Front()
				}
			}
			r.MoveAfter(e, next)

		}

	}

	findZero := func() *list.Element {
		//find zero
		e := r.Front()
		for {
			if e.Value.(int) == 0 {
				ret := e
				return ret
			}
			e = e.Next()
		}
	}

	getFromZero := func(dist int) int {
		z := findZero()
		for i := 0; i < dist; i++ {
			//	fmt.Println(i)
			z = z.Next()
			if z == nil {
				z = r.Front()
			}
		}
		return z.Value.(int)
	}

	a := getFromZero(1000)
	b := getFromZero(2000)
	c := getFromZero(3000)
	fmt.Println(a, b, c)
	fmt.Println(a + b + c)
}

func atoi(str string) int {
	c, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
