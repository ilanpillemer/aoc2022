package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Blue struct {
	id           int
	oreCost      int
	clayCost     int
	obsidianCost [2]int
	geodeCost    [2]int
	robots       [4]int
	collection   [4]int
}

var part1 bool

func main() {
	//part1 = true
	factories := map[int]Blue{}
	file, _ := os.Open("INPUT2")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//	fmt.Println(line)
		blue := process(line)
		factories[blue.id] = blue
		//	fmt.Printf("%#v\n", blue)
	}
	for k, v := range factories {
		fmt.Printf("%d -> %#v\n", k, v)
	}
	levels := map[int]int{}
	if part1 {
		for k, v := range factories {
			fmt.Printf("Blueprint %d:\n", k)
			v.robots[ORE] = 1
			levels[k] = solve(v, 24)
			fmt.Println(k, levels[k])
		}
		total := 0
		for k, v := range levels {
			total += (k * v)
		}
		fmt.Println("part1", total)
		os.Exit(0)
	}

	for k, v := range factories {
		fmt.Printf("Blueprint %d:\n", k)
		v.robots[ORE] = 1
		levels[k] = solve(v, 32)
		fmt.Println(k, levels[k])
	}
	total := 1
	for _, v := range levels {
		total *= v
	}
	fmt.Println("part2", total)

}

const ORE = 0
const CLAY = 1
const OBSIDIAN = 2
const GEODE = 3

const OTHER = 1

var memo = map[string]int{}

func copyB(b Blue) Blue {
	r := Blue{}
	r.oreCost = b.oreCost
	r.clayCost = b.clayCost
	r.obsidianCost = b.obsidianCost
	r.geodeCost = b.geodeCost
	r.robots = [4]int{}
	for i, j := range b.robots {
		r.robots[i] = j
	}
	r.collection = [4]int{}
	for i, j := range b.collection {
		r.collection[i] = j
	}
	return r
}

type State struct {
	b    Blue
	time int
}

func solve(b Blue, time int) int {
	best := 0
	besttime := 0
	S := map[State]bool{}
	Q := []State{State{b, time}}
	i := 0
	for len(Q) > 0 {
		next := Q[0]
		Q = Q[1:]
		if best < next.b.collection[GEODE] {
			best = next.b.collection[GEODE]
			besttime = next.time
			fmt.Printf("BEST %d: %#v\n", next.time, next)
		}

		if next.time >= besttime && next.b.collection[GEODE] < best {
			continue
		}

		if next.time == 7 && next.b.collection[GEODE] == 0 {
			continue
		}

		if next.time == 0 {
			continue
		}
		if S[next] {
			continue
		}
		i++
		S[next] = true
		current := next.b
		if i%100000000 == 0 {
			fmt.Printf("%d: %#v\n", time-next.time+1, current)
		}

		// dont build any robots
		opt1 := acquire(current)

		Q = append(Q, State{copyB(opt1), next.time - 1})

		opt2 := buildOreRobot(current)

		Q = append(Q, State{copyB(opt2), next.time - 1})

		opt3 := buildClayRobot(current)
		Q = append(Q, State{copyB(opt3), next.time - 1})

		opt4 := buildObsRobot(current)
		Q = append(Q, State{copyB(opt4), next.time - 1})

		opt5 := buildGeodeRobot(current)
		Q = append(Q, State{copyB(opt5), next.time - 1})

	}
	return best
}

func buildOreRobot(b Blue) Blue {
	i := 0
	newB := copyB(b)
	if newB.robots[ORE] < 5 {
		if newB.oreCost <= newB.collection[ORE] {
			newB.collection[ORE] -= newB.oreCost
			newB.robots[ORE]++
			i = 1
		}
	}
	newB.collection[ORE] += (newB.robots[ORE] - i)
	newB.collection[CLAY] += newB.robots[CLAY]
	newB.collection[OBSIDIAN] += newB.robots[OBSIDIAN]
	newB.collection[GEODE] += newB.robots[GEODE]
	return newB
}

//var OBSIDIAN = 2
//var GEODE = 3

func buildClayRobot(b Blue) Blue {
	i := 0
	newB := copyB(b)
	if newB.robots[CLAY] < 21 {
		if newB.clayCost <= newB.collection[ORE] {
			newB.collection[ORE] = newB.collection[ORE] - newB.clayCost
			newB.robots[CLAY]++
			i = 1
		}
	}
	newB.collection[ORE] += (newB.robots[ORE])
	newB.collection[CLAY] += (newB.robots[CLAY] - i)
	newB.collection[OBSIDIAN] += newB.robots[OBSIDIAN]
	newB.collection[GEODE] += newB.robots[GEODE]
	return newB
}

func buildObsRobot(b Blue) Blue {
	i := 0
	limit := 13
	if part1 {
		limit = 21
	}
	newB := copyB(b)
	if newB.robots[OBSIDIAN] < limit {
		if newB.obsidianCost[ORE] <= newB.collection[ORE] &&
			newB.obsidianCost[OTHER] <= newB.collection[CLAY] {
			newB.collection[ORE] -= newB.obsidianCost[0]
			newB.collection[CLAY] -= newB.obsidianCost[1]
			newB.robots[OBSIDIAN]++
			i = 1
		}
	}
	newB.collection[ORE] += (newB.robots[ORE])
	newB.collection[CLAY] += (newB.robots[CLAY])
	newB.collection[OBSIDIAN] += (newB.robots[OBSIDIAN] - i)
	newB.collection[GEODE] += newB.robots[GEODE]

	return newB
}

func buildGeodeRobot(b Blue) Blue {
	i := 0
	newB := copyB(b)
	if newB.geodeCost[ORE] <= newB.collection[ORE] &&
		newB.geodeCost[OTHER] <= newB.collection[OBSIDIAN] {
		newB.collection[ORE] -= newB.geodeCost[0]
		newB.collection[OBSIDIAN] -= newB.geodeCost[1]
		newB.robots[GEODE]++
		i = 1
	}
	newB.collection[ORE] += (newB.robots[ORE])
	newB.collection[CLAY] += (newB.robots[CLAY])
	newB.collection[OBSIDIAN] += (newB.robots[OBSIDIAN])
	newB.collection[GEODE] += (newB.robots[GEODE] - i)

	return newB
}

func acquire(b Blue) Blue {
	newB := copyB(b)
	newB.collection[ORE] += newB.robots[ORE]
	newB.collection[CLAY] += newB.robots[CLAY]
	newB.collection[OBSIDIAN] += newB.robots[OBSIDIAN]
	newB.collection[GEODE] += newB.robots[GEODE]
	return newB
}

func process(line string) Blue {
	blue := Blue{}
	first := strings.Split(line, ":")
	blue.id = atoi(strings.Fields(first[0])[1])
	rows := strings.Split(first[1], ".")
	blue.oreCost = atoi(strings.Fields(rows[0])[4])
	blue.clayCost = atoi(strings.Fields(rows[1])[4])
	blue.obsidianCost[0] = atoi(strings.Fields(rows[2])[4])
	blue.obsidianCost[1] = atoi(strings.Fields(rows[2])[7])
	blue.geodeCost[0] = atoi(strings.Fields(rows[3])[4])
	blue.geodeCost[1] = atoi(strings.Fields(rows[3])[7])
	//fmt.Println(rows)
	return blue
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func atoi(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}
