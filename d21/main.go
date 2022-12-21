package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var part1 bool

func main() {
	//part1 = true
	file, err := os.Open("INPUT2")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		process(line)
	}
	//fmt.Println("unsorted")
	//	for k, v := range deps {
	//	fmt.Println(k, v)
	//	}

	sorted := topoSort(deps)
	//	fmt.Println("sorted")
	//	for k, v := range sorted {
	//		fmt.Println(k, v)
	//	}

	if part1 {
		for _, v := range sorted {
			results[v] = solve(actions[v])
		}
		fmt.Println("part1", results["root"])
	}

	high := 9223372036854775807
	//high = 1000
	low := 0
	i := 0

	for {
		//fmt.Println("low", low)
		//fmt.Println("high", high)
		actions["humn"] = []string{fmt.Sprintf("%d", i)}
		//fmt.Println(actions["humn"])
		for _, v := range sorted {
			results[v] = solve(actions[v])
		}

		//fmt.Println("temp part2", i, results["root"], sign(results["root"]))
		tmp := results["root"]
		if results["root"] == 0 {
			fmt.Println("part2:", i, results["root"])
			os.Exit(0)
		}
		if tmp > 0 { // sample needs to be "<"
			low = i
			i = low + ((high - low) / 2)
		} else {
			high = i
			i = low + ((high - low) / 2)
		}

	}

}

func sign(x float64) int {
	if x > 0 {
		return 1
	}

	if x < 1 {
		return -1
	}

	return 0
}

func solve(action []string) float64 {
	if len(action) == 1 {
		return atoi(action[0])
	}
	//	fmt.Println("action,", action)
	lhs, rhs := results[action[0]], results[action[2]]
	switch action[1] {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	case "=":
	//	fmt.Println(lhs, "-", rhs, "=", lhs-rhs)
		return lhs - rhs
	default:
		panic("unknown op")

	}
	return 0

}

func atoi(str string) float64 {
	x, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return float64(x)
}

var deps = map[string][]string{}
var results = map[string]float64{}
var actions = map[string][]string{}

func process(line string) {
	sides := strings.Split(line, ":")
	lhs, rhs := strings.TrimSpace(sides[0]), strings.TrimSpace(sides[1])
	//fmt.Println(lhs, rhs)
	// dependencies
	fields := strings.Fields(rhs)
	actions[lhs] = fields
	if len(fields) > 1 {
		// has dependencies
		d1 := fields[0]
		d2 := fields[2]
		deps[lhs] = []string{d1, d2}
	}

}

//https://github.com/adonovan/gopl.io/blob/master/ch5/toposort/main.go

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}
