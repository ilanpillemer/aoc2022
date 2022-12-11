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

type Monkey struct {
	id          int
	items       []int64
	fn          func(old int64) int64
	div         int
	iftrue      int
	iffalse     int
	inspections int
}

var monkeys = map[int]Monkey{}
var count int

func main() {
	fmt.Println("d11")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			count++
			var i int
			fmt.Sscanf(line, "Monkey %d:", &i)
			m := Monkey{}
			m.items = []int64{}
			m.id = i
			for scanner.Scan() {
				monkeys[i] = m
				line := scanner.Text()
				if line == "" {
					break
				}
				line = trim(line)
				xs := strings.Split(line, ":")
				rhs := xs[1]
				switch {
				case strings.HasPrefix(line, "Starting"):
					items := strings.Split(rhs, ",")
					for _, e := range items {
						m.items = append(m.items, atoi(e))
					}
				case strings.HasPrefix(line, "Operation"):
					ys := strings.Split(rhs, "=")
					ys[1] = trim(ys[1])
					xs := strings.Fields(ys[1])
					op := trim(xs[1])
					rhs := trim(xs[2])

					switch op {
					case "+":
						if rhs == "old" {
							m.fn = func(old int64) int64 { return old + old }
							continue
						}
						m.fn = func(old int64) int64 { return old + atoi(rhs) }
					case "*":
						if rhs == "old" {
							m.fn = func(old int64) int64 { return old * old }
							continue
						}
						m.fn = func(old int64) int64 { return old * atoi(rhs) }
					default:
						panic(op)
					}
				case strings.HasPrefix(line, "Test"):
					var d int
					rhs = trim(rhs)
					fmt.Sscanf(rhs, "divisible by %d", &d)
					m.div = d

				case strings.HasPrefix(line, "If true"):
					var d int
					rhs = trim(rhs)
					fmt.Sscanf(rhs, "throw to monkey %d", &d)
					m.iftrue = d

				case strings.HasPrefix(line, "If false"):
					var d int
					rhs = trim(rhs)
					fmt.Sscanf(rhs, "throw to monkey %d", &d)
					m.iffalse = d
				}

			}
		}
	}
	fmt.Printf("%d %v\n", count, monkeys)
	display(monkeys)
	for i := 1; i < 20+1; i++ {
		fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", i)
		round(monkeys)
		display(monkeys)
		fmt.Println()
	}
	fmt.Println("part1", part1(monkeys))
}

func part1(xs map[int]Monkey) int {
	total := 0
	list := []int{}
	for k := 0; k < count; k++ {
		m := monkeys[k]
		list = append(list, m.inspections)
	}
	sort.Ints(list)
	fmt.Println(list)
	l := len(list)
	total = list[l-1] * list[l-2]
	return total
}

func round(xs map[int]Monkey) {
	for k := 0; k < count; k++ {
		m := xs[k]
		//fmt.Printf("Monkey :%d\n", k)
		for _, item := range m.items {
			m.inspections++
			//fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
			item := m.fn(item)
			//fmt.Printf("    Worry level becomes %d.\n", item)
			item = item / 3
			//fmt.Printf("    Boredom makes worry level %d.\n", item)
			if item%int64(m.div) == 0 {
				//fmt.Printf("    Current worry level is divisible by %d.\n", m.div)
				nextm := xs[m.iftrue]
				nextm.items = append(nextm.items, item)
				xs[m.iftrue] = nextm
				//fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", item, m.iftrue)
			} else {
				//fmt.Printf("    Current worry level is not divisible by %d.\n", m.div)
				nextm := xs[m.iffalse]
				nextm.items = append(nextm.items, item)
				xs[m.iffalse] = nextm
				//fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", item, m.iffalse)
			}
		}
		m.items = []int64{}
		xs[k] = m
	}
}

func atoi(x string) int64 {
	x = strings.TrimSpace(x)
	i, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func trim(x string) string {
	return strings.TrimSpace(x)
}

func display(xs map[int]Monkey) {
	for k := 0; k < count; k++ {
		m := monkeys[k]
		fmt.Printf("Monkey %d: %v  [inspections: %d]\n", m.id, m.items, m.inspections)
	}
}
