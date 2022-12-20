package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var key = int64(1)
var key = int64(811589153)

func main() {

	file, _ := os.ReadFile("INPUT")
	lines := strings.Fields(string(file))
	values := []int64{}
	pointers := []*list.Element{}
	for i := 0; i < len(lines); i++ {
		values = append(values, key*atoi(lines[i]))
	}
	r := list.New()

	for i := 0; i < len(values); i++ {
		e := r.PushBack(values[i])
		pointers = append(pointers, e)
	}

	reducer := int64(len(lines) - 1)
	for k := 0; k < 10; k++ {
		for i := 0; i < len(lines); i++ {
			v := pointers[i].Value.(int64)
			e := pointers[i]
			if v < 0 {
				next := e
				for j := int64(0); j < abs(v)%reducer; j++ {
					next = next.Prev()
					if next == e {
						next = next.Prev()
					}
					if next == nil {
						next = r.Back()
						if next == e {
							next = next.Prev()
						}
					}
				}
				r.MoveBefore(e, next)

			} else {
				next := e
				for j := int64(0); j < abs(v)%reducer; j++ {
					next = next.Next()
					if next == e {
						next = next.Next()
					}
					if next == nil {
						next = r.Front()
						if next == e {
							next = next.Next()
						}
					}
				}
				r.MoveAfter(e, next)

			}
		}
	}

	findZero := func() *list.Element {
		//find zero
		e := r.Front()
		for {
			if e.Value.(int64) == 0 {
				ret := e
				return ret
			}
			e = e.Next()
		}
	}

	getFromZero := func(dist int) int64 {
		z := findZero()
		for i := 0; i < dist; i++ {
			//	fmt.Println(i)
			z = z.Next()
			if z == nil {
				z = r.Front()
			}
		}
		return z.Value.(int64)
	}

	a := getFromZero(1000)
	b := getFromZero(2000)
	c := getFromZero(3000)
	fmt.Println(a, b, c)
	fmt.Println(a + b + c)
}

func atoi(str string) int64 {
	c, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return c

}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
