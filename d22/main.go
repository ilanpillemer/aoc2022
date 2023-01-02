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
}

const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

func main() {
	maxx := 0
	maxy := 0
	facing := 0 // facing to the right
	tiles := map[image.Point]*tile{}
	fmt.Println("d22")
	b, _ := os.ReadFile("INPUT")
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
	move := func(count int) {
		fmt.Printf("facing %d, position %v\n", facing, current.position)
		switch facing {
		case 0:
			fmt.Println("move right", count)
			for i := 0; i < count; i++ {
				if current.right.value != '#' {
					current = current.right
				}
			}
		case 1:
			fmt.Println("move down", count)
			for i := 0; i < count; i++ {
				if current.down.value != '#' {
					current = current.down
				}
			}
		case 2:
			fmt.Println("move left", count)
			for i := 0; i < count; i++ {
				if current.left.value != '#' {
					current = current.left
				}
			}
		case 3:
			fmt.Println("move up", count)
			for i := 0; i < count; i++ {
				if current.up.value != '#' {
					current = current.up
				}
			}
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
