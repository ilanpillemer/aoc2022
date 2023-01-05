package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type tile struct {
	position image.Point
	value    rune
	up       *tile
	down     *tile
	left     *tile
	right    *tile
	side     int
}

const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

func main() {
	//	part1()
	part2()
}

func part2() {
	maxx := 0
	maxy := 0
	facing := 0 // facing to the right
	tiles := map[image.Point]*tile{}
	fmt.Println("d22")
	b, _ := os.ReadFile("SAMPLE")
	xs := strings.Split(string(b), "\n\n")
	board, path := xs[0], xs[1]
	grid := strings.Split(board, "\n")
	fmt.Println(board)
	fmt.Println(path)
	path = path + "$"
	//parse to get all tiles

	for i, t := range grid {
		for j, v := range t {
			if v != ' ' {
				maxx = max(maxx, j)
				maxy = max(maxy, i)
				e := &tile{
					position: image.Pt(j, i),
					value:    v,
				}
				tiles[image.Pt(j, i)] = e
			}
		}
	}

	// add wrapping links
	//top
	// this will be where y is 0
	fmt.Println("SIDE1")
	//sample rules rules
	for _, v := range tiles {
		if v.position.Y >= 0 && v.position.Y < 4 {
			v.side = 1
			tiles[v.position] = v
			continue
		}
		if v.position.Y >= 8 && v.position.Y < 12 {
			v.side = 5
			if v.position.X >= 12 {
				v.side = 6
			}
			tiles[v.position] = v
			continue
		}
		switch {
		case v.position.X >= 4 && v.position.X < 8:
			v.side = 3
		case v.position.X >= 8:
			v.side = 4
		default:
			v.side = 2
		}
		tiles[v.position] = v
		continue
	}

	//parse to link tiles
	fmt.Println(maxx, maxy)
	for k, v := range tiles {
		// l + r
		left := v.position.X - 1
		for {
			if left < 0 {
				break
			}
			if leftTile, ok := tiles[image.Pt(left, v.position.Y)]; ok &&
				left != v.position.X &&
				leftTile.side == v.side {
				v.left = leftTile
				tiles[k] = v
				break
			}
			left--
		}
		right := v.position.X + 1
		for {
			if right > maxx {
				break
			}
			if rightTile, ok := tiles[image.Pt(right, v.position.Y)]; ok && right != v.position.X &&
				rightTile.side == v.side {
				v.right = rightTile
				tiles[k] = v
				break
			}
			right++
		}
		// u + d
		up := v.position.Y - 1
		for {
			if up < 0 {
				break
			}
			if upTile, ok := tiles[image.Pt(v.position.X, up)]; ok && up != v.position.Y &&
				upTile.side == v.side {
				v.up = upTile
				tiles[k] = v
				break
			}
			up--
		}
		down := v.position.Y + 1
		for {
			if down > maxy {
				break
			}
			if downTile, ok := tiles[image.Pt(v.position.X, down)]; ok && down != v.position.Y &&
				downTile.side == v.side {
				v.down = downTile
				tiles[k] = v
				break
			}
			down++
		}
	}

	//	for _, v := range tiles {
	//		fmt.Println(v.side, v.position)
	//
	//	}

	//os.Exit(0)
	//	for _, v := range tiles {
	//		fmt.Printf("up: %v -> %v\n", v.position, v.up)
	//		fmt.Printf("down: %v -> %v\n", v.position, v.down)
	//		fmt.Printf("left: %v -> %v\n", v.position, v.left)
	//		fmt.Printf("right: %v -> %v\n", v.position, v.right)
	//
	//	}

	leftmost := 2147483647
	for k := range tiles {
		if k.Y == 0 {
			leftmost = min(leftmost, k.X)
		}
	}
	current := tiles[image.Pt(leftmost, 0)]
	var move func(int)
	move = func(count int) {
		if count == 0 {
			return
		}
		fmt.Printf("facing %d, position side %d %v\n", facing, current.position, current.side)
		newfacing := facing
		switch facing {
		case 0:
			if current.right == nil {
				switch current.side {
				case 4:
					current.right = tiles[image.Pt(current.position.X+3, current.position.Y+3)]
					newfacing = DOWN
				case 2:
					current.right = tiles[image.Pt(current.position.X+1, current.position.Y)]
					newfacing = RIGHT
				default:
					panic("oops")
				}
			}
			if current.right.value != '#' {
				if current.side != current.down.side {
					fmt.Printf("side [%d] -> side [%d]\n", current.side, current.down.side)

				}
				facing = newfacing
				current = current.right
			}

			move(count - 1)
		case 1:
			if current.down == nil {
				switch current.side {
				case 1:
					current.down = tiles[image.Pt(current.position.X, current.position.Y+1)]
					newfacing = DOWN
				case 5:
					current.down = tiles[image.Pt(current.position.X-9, current.position.Y-4)]
					newfacing = UP
				default:
					panic("oops")
				}
			}
			if current.down.value != '#' {
				if current.side != current.down.side {
					fmt.Printf("side [%d] -> side [%d]\n", current.side, current.down.side)
				}
				facing = newfacing
				current = current.down
			}

			move(count - 1)
		case 2:
			if current.left == nil {
				switch current.side {
				case 6:
					current.left = tiles[image.Pt(current.position.X-1, current.position.Y)]
					newfacing = LEFT
				default:
					panic("oops")
				}
			}
			if current.left.value != '#' {

				current = current.left
			}

			move(count - 1)
		case 3:
			if current.up == nil {
				switch current.side {
				case 3:
					fmt.Println(image.Pt(current.position.X+8, current.position.Y+4))
					current.up = tiles[image.Pt(current.position.X+8, current.position.Y+7)]
					newfacing = UP
				default:
					panic("oops")
				}
			}
			if current.up.value != '#' {
				facing = newfacing
				current = current.up
			}
			facing = newfacing

			move(count - 1)
		default:
			panic("oops")

		}
	}

	process := func() {
		p := 0
		for {
			if p >= len(path) {
				break
			}
			num := ""
			for {
				if unicode.IsDigit(rune(path[p])) && p < len(path) {
					num = num + string(path[p])
					p++
				} else {
					println(num)
					move(atoi(num))
					break
				}
			}
			switch path[p] {
			case 'R':
				switch facing {
				case 0:
					facing = 1
				case 1:
					facing = 2
				case 2:
					facing = 3
				case 3:
					facing = 0
				}
			case 'L':
				switch facing {
				case 0:
					facing = 3
				case 1:
					facing = 0
				case 2:
					facing = 1
				case 3:
					facing = 2
				}
			case '$':
				fmt.Println("finished")
				fmt.Printf("final position index 1: facing %d, position %v %v\n", facing, current.position.X+1, current.position.Y+1)
				fmt.Println("password", facing+(4*(current.position.X+1))+((current.position.Y+1)*1000))
			default:
				panic("oops")
			}
			fmt.Printf("%c\n", path[p])
			p++
		}

	}
	process()
}

func part1() {
	maxx := 0
	maxy := 0
	facing := 0 // facing to the right
	tiles := map[image.Point]*tile{}
	fmt.Println("d22")
	b, _ := os.ReadFile("SAMPLE")
	xs := strings.Split(string(b), "\n\n")
	board, path := xs[0], xs[1]
	grid := strings.Split(board, "\n")
	fmt.Println(board)
	fmt.Println(path)
	path = path + "$"
	//parse to get all tiles
	for i, t := range grid {
		for j, v := range t {
			if v != ' ' {
				maxx = max(maxx, j)
				maxy = max(maxy, i)
				e := &tile{
					position: image.Pt(j, i),
					value:    v,
				}
				tiles[image.Pt(j, i)] = e
			}
		}
	}
	//parse to link tiles
	fmt.Println(maxx, maxy)
	for k, v := range tiles {
		// l + r
		left := v.position.X - 1
		for {
			if left < 0 {
				left = maxx
			}
			if leftTile, ok := tiles[image.Pt(left, v.position.Y)]; ok && left != v.position.X {
				v.left = leftTile
				tiles[k] = v
				break
			}
			left--
		}
		right := v.position.X + 1
		for {
			if right > maxx {
				right = 0
			}
			if rightTile, ok := tiles[image.Pt(right, v.position.Y)]; ok && right != v.position.X {
				v.right = rightTile
				tiles[k] = v
				break
			}
			right++
		}
		// u + d
		up := v.position.Y - 1
		for {
			if up < 0 {
				up = maxy
			}
			if upTile, ok := tiles[image.Pt(v.position.X, up)]; ok && up != v.position.Y {
				v.up = upTile
				tiles[k] = v
				break
			}
			up--
		}
		down := v.position.Y + 1
		for {
			if down > maxy {
				down = 0
			}
			if downTile, ok := tiles[image.Pt(v.position.X, down)]; ok && down != v.position.Y {
				v.down = downTile
				tiles[k] = v
				break
			}
			down++
		}
	}

	//	for _, v := range tiles {
	//		fmt.Printf("up: %v -> %c\n", v.position, v.up.value)
	//		fmt.Printf("down: %v -> %c\n", v.position, v.down.value)
	//		fmt.Printf("left: %v -> %c\n", v.position, v.left.value)
	//		fmt.Printf("right: %v -> %c\n", v.position, v.right.value)
	//
	//}

	leftmost := 2147483647
	for k := range tiles {
		if k.Y == 0 {
			leftmost = min(leftmost, k.X)
		}
	}
	current := tiles[image.Pt(leftmost, 0)]
	var move func(int)
	move = func(count int) {
		if count == 0 {
			return
		}
		fmt.Printf("facing %d, position %v\n", facing, current.position)
		switch facing {
		case 0:
			if current.right.value != '#' {
				current = current.right
			}
			move(count - 1)
		case 1:
			if current.down.value != '#' {
				current = current.down
			}
			move(count - 1)
		case 2:

			if current.left.value != '#' {
				current = current.left
			}
			move(count - 1)
		case 3:
			if current.up.value != '#' {
				current = current.up
			}
			move(count - 1)
		default:
			panic("oops")

		}
	}

	process := func() {
		p := 0
		for {
			if p >= len(path) {
				break
			}
			num := ""
			for {
				if unicode.IsDigit(rune(path[p])) && p < len(path) {
					num = num + string(path[p])
					p++
				} else {
					println(num)
					move(atoi(num))
					break
				}
			}
			switch path[p] {
			case 'R':
				switch facing {
				case 0:
					facing = 1
				case 1:
					facing = 2
				case 2:
					facing = 3
				case 3:
					facing = 0
				}
			case 'L':
				switch facing {
				case 0:
					facing = 3
				case 1:
					facing = 0
				case 2:
					facing = 1
				case 3:
					facing = 2
				}
			case '$':
				fmt.Println("finished")
				fmt.Printf("final position index 1: facing %d, position %v %v\n", facing, current.position.X+1, current.position.Y+1)
				fmt.Println("password", facing+(4*(current.position.X+1))+((current.position.Y+1)*1000))
			default:
				panic("oops")
			}
			fmt.Printf("%c\n", path[p])
			p++
		}

	}
	process()
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

func atoi(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return x
}
