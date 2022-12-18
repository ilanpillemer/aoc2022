package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var xs = map[string]Valve{}
var seen = map[string]bool{}
var nodes = map[string]bool{}
var zeronodes = []string{}

type destination struct {
	valve    string
	distance int
}

type Valve struct {
	valve  string
	rate   int
	valves []destination
}

// answer for my INPUT to part 1 is 1617

func main() {
	fmt.Println("d16")
	//part1 = true
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		valve, rate, valves := process(line)
		xs[valve] = Valve{valve, rate, valves}

	}
	//	for k, v := range xs {
	//		fmt.Printf("%s -> %v\n", k, v)
	//	}

	fmt.Println("simplify")
	gra := makeGraph(xs)

	for i, g := range gra {
		fmt.Println(i, g)
	}
	dist := floydWarshall(gra)
	//dist[][] will be the output matrix that will finally
	//have the shortest distances between every pair of vertices
	fmt.Println()
	fmt.Println("final distances\n")
	for _, d := range dist {
		fmt.Printf("%4g\n", d)
	}

	// test proposed answer - to get a suitable scoring algo

	//	//	part1 := 0
	//
	pathI := maps.Clone(keyIndexes)
	delete(pathI, "AA")
	for _, j := range zeronodes {
		delete(pathI, j)
	}
	//pathOptions := maps.Keys(pathI)

	best = solve("AA", map[string]bool{}, 30)

	fmt.Println("winning score is:", best)
}

var best = 0
var limit = 30

// what makes a state unique
//node .. nodes_open ... time

var memo = map[string]int{}

func solve(N string, V map[string]bool, time int) int {
	if time == 0 {
		return 0
	}
	// if I am at node [N] and I have opened [V] valves. and I have [T] time
	// left. What can I score fromt this position?
	key := fmt.Sprintf("%v%v%v", N, V, time)
	if i, ok := memo[key]; ok {
		return i
	}
	ans := 0
	//if I am not open, open me.
	//to be replaced with distance matrix if this works later
	if xs[N].rate > 0 && !V[N] {
		newV := maps.Clone(V)
		newV[N] = true
		ans = max(ans, (time-1)*xs[N].rate+solve(N, newV, time-1))
	}
	for _, e := range xs[N].valves {
		//fmt.Println(e)
		ans = max(ans, solve(e.valve, V, time-1))
	}

	memo[key] = ans
	return ans
}

var keyIndexes = map[string]int{}
var rkeyIndexes = map[int]string{}

func makeGraph(xs map[string]Valve) [][]graph {
	nodes := maps.Keys(xs)

	for k, v := range xs {
		if v.rate == 0 {
			zeronodes = append(zeronodes, k)
		}
	}
	sort.Strings(nodes)
	for i, key := range nodes {
		//fmt.Println(i, key)
		keyIndexes[key] = i
		rkeyIndexes[i] = key
	}

	fmt.Println("nodes", nodes)
	fmt.Println(keyIndexes)

	//	gra := [][]graph{
	//		1: {{2, 3}, {3, 8}, {5, -4}},
	//		2: {{4, 1}, {5, 7}},
	//		3: {{2, 4}},
	//		4: {{1, 2}, {3, -5}},
	//		5: {{4, 6}},
	//	}

	gra := [][]graph{}
	for u := range nodes {
		// get edges for node
		node := rkeyIndexes[u]
		//	fmt.Println("node -> ", node, " ", xs[node].valves)
		row := []graph{}
		for _, v := range xs[node].valves {
			e := graph{keyIndexes[v.valve], 1}
			row = append(row, e)
		}
		gra = append(gra, row)

	}
	return gra
}

func process(line string) (string, int, []destination) {
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	xs := strings.Split(line, ";")
	x1 := xs[0]

	ys := strings.Split(xs[1][23:], ",")
	valves := []destination{}
	for _, v := range ys {
		valves = append(valves, destination{strings.TrimSpace(v), 1})
	}
	// Valve AA has flow rate=0
	var valve string
	var rate int
	fmt.Sscanf(x1, "Valve %s has flow rate=%d", &valve, &rate)
	return valve, rate, valves
}

func atoi(x string) int {
	x = strings.TrimSpace(x)
	y, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return y
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type graph struct {
	to int
	wt float64
}

// https://www.golangprograms.com/golang-program-for-implementation-of-floyd-warshall-algorithm.html
func floydWarshall(g [][]graph) [][]float64 {
	dist := make([][]float64, len(g))
	for i := range dist {
		di := make([]float64, len(g))
		for j := range di {
			di[j] = math.Inf(1)
		}
		di[i] = 0
		dist[i] = di
	}

	fmt.Println("initial distances\n")
	for _, d := range dist {
		fmt.Printf("%4g\n", d)
	}

	for u, graphs := range g {
		for _, v := range graphs {
			dist[u][v.to] = v.wt
		}
	}

	fmt.Println("initial weights inserted\n")
	for _, d := range dist {
		fmt.Printf("%4g\n", d)
	}

	for k, dk := range dist {
		for _, di := range dist {
			for j, dij := range di {
				if d := di[k] + dk[j]; dij > d {
					di[j] = d
				}
			}
		}
	}
	return dist
}

//func run() {
//	gra := [][]graph{
//		1: {{2, 3}, {3, 8}, {5, -4}},
//		2: {{4, 1}, {5, 7}},
//		3: {{2, 4}},
//		4: {{1, 2}, {3, -5}},
//		5: {{4, 6}},
//	}
//
//	fmt.Println(gra)
//	dist := floydWarshall(gra)
//	//dist[][] will be the output matrix that will finally
//	//have the shortest distances between every pair of vertices
//	fmt.Println()
//	fmt.Println("final distances\n")
//	for _, d := range dist {
//		fmt.Printf("%4g\n", d)
//	}
//}
