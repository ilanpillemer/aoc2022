package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("INPUT")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		process(line)
	}
	fmt.Println("unsorted")
	for k, v := range deps {
		fmt.Println(k, v)
	}
	sorted := topoSort(deps)
	fmt.Println("sorted")
	for k, v := range sorted {
		fmt.Println(k, v)
	}

	for _, v := range sorted {
		results[v] = solve(actions[v])
	}

	fmt.Println("part1", results["root"])

}

func solve(action []string) int {
	if len(action) == 1 {
		return atoi(action[0])
	}
	fmt.Println("action,", action)
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
	default:
		panic("unknown op")

	}
	return 0

}

func atoi(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}

var deps = map[string][]string{}
var results = map[string]int{}
var actions = map[string][]string{}

func process(line string) {
	sides := strings.Split(line, ":")
	lhs, rhs := strings.TrimSpace(sides[0]), strings.TrimSpace(sides[1])
	fmt.Println(lhs, rhs)
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

// https://en.wikipedia.org/wiki/Topological_sorting
// L ← Empty list that will contain the sorted elements
// S ← Set of all nodes with no incoming edge
//
// while S is not empty do
//
//	remove a node n from S
//	add n to L
//	for each node m with an edge e from n to m do
//	    remove edge e from the graph
//	    if m has no other incoming edges then
//	        insert m into S
//
// if graph has edges then
//
//	return error   (graph has at least one cycle)
//
// else
//
//	return L   (a topologically sorted order)
//
//https://github.com/adonovan/gopl.io/blob/master/ch4/graph/main.go
//https://github.com/adonovan/gopl.io/blob/master/ch5/toposort/main.go

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

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
