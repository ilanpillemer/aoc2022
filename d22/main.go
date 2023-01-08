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
	xoffset  int
	yoffset  int
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

func printme(side int, posfrom image.Point, dir int, posto image.Point, tiles map[image.Point]*tile) {
	tile := tiles[posfrom]
	if posfrom.X == tile.xoffset || posfrom.X == tile.xoffset+49 {
		if posfrom.Y == tile.yoffset || posfrom.Y == tile.yoffset+49 {
			//	fmt.Println(side, posfrom, dir, "-->", posto)
		}
	}

}

func wrap(side int, xy image.Point, direction int) image.Point {
	z := side*10 + direction
	switch z {
	case 13: // side 1 to side 6
		vx, vy := offsets(6)
		wx, _ := offsets(1)
		x := vx
		y := vy + (xy.X - wx)
		return image.Pt(x, y)
	case 23: // side 2 to 6
		vx, vy := offsets(6)
		wx, _ := offsets(2)
		y := vy + 49
		x := xy.X - wx + vx
		return image.Pt(x, y)
	case 43: //side 4 to side 3
		vx, vy := offsets(3)
		wx, _ := offsets(4)
		x := vx
		y := vy + (xy.X - wx)
		return image.Pt(x, y)
	case 61: // side 6 to side 2
		vx, vy := offsets(2)
		wx, _ := offsets(6)
		y := vy
		x := xy.X - wx + vx
		return image.Pt(x, y)
	case 21: // side 2 to side 3
		vx, vy := offsets(3)
		wx, _ := offsets(2)
		x := vy + 49
		y := vx + (xy.X - wx)
		return image.Pt(x, y)
	case 51: // side 5 to side 6
		vx, vy := offsets(6)
		wx, _ := offsets(5)
		x := vx + 49
		y := vy + (xy.X - wx)
		return image.Pt(x, y)
	case 62:
		vx, vy := offsets(1)
		_, wy := offsets(6)
		x := vx + (xy.Y - wy)
		y := vy
		return image.Pt(x, y)
	case 12: //side 1 -> side 4
		vx, vy := offsets(4)
		_, wy := offsets(1)
		x := vx
		y := vy + (49 - (xy.Y - wy))
		return image.Pt(x, y)
	case 32: //side 3 -> side 4
		vx, vy := offsets(4)
		_, wy := offsets(3)
		x := vx + (xy.Y - wy)
		y := vy
		return image.Pt(x, y)
	case 42: // side 4->side1
		vx, vy := offsets(1)
		_, wy := offsets(4)
		//	x := vx + (49 - (xy.Y - wy))
		x := vx
		y := vy + (49 - (xy.Y - wy))
		return image.Pt(x, y)
	case 20:
		vx, vy := offsets(5)
		_, wy := offsets(2)
		//	x := vx + (49 - (xy.Y - wy))
		x := vx + 49
		y := vy + (49 - (xy.Y - wy))
		return image.Pt(x, y)
	case 30:
		vx, vy := offsets(2)
		_, wy := offsets(3)

		x := vx + (xy.Y - wy)
		y := vy + 49
		return image.Pt(x, y)
	case 60:
		vx, vy := offsets(5)
		_, wy := offsets(6)

		x := vx + (xy.Y - wy)
		y := vy + 49
		return image.Pt(x, y)
	case 50:
		vx, vy := offsets(2)
		_, wy := offsets(5)
		//	x := vx + (49 - (xy.Y - wy))
		x := vx + 49
		y := vy + (49 - (xy.Y - wy))
		return image.Pt(x, y)
	default:
		panic(fmt.Sprintf("missing wrap %d", z))
	}

	return image.Pt(0, 0)
}

func offsets(side int) (int, int) {
	switch side {
	case 1:
		return 50, 0
	case 2:
		return 100, 0
	case 3:
		return 50, 50
	case 4:
		return 0, 100
	case 5:
		return 50, 100
	case 6:
		return 0, 150
	default:
		panic("illegal side")
	}
	return 0, 0
}

var xs = map[int]int{
	26: UP,
	12: RIGHT,
	21: LEFT,
	16: RIGHT,
	62: DOWN,
	23: LEFT,
	31: UP,
	64: UP,
	41: RIGHT,
	14: RIGHT,
	46: DOWN,
	32: UP,
	25: LEFT,
	53: UP,
	13: DOWN,
	35: DOWN,
	34: DOWN,
	65: UP,
	56: LEFT,
	45: RIGHT,
	54: LEFT,
	61: DOWN,
	43: RIGHT,
	52: LEFT,
}

func newfacing(f, from, to int) int {
	if from == to {
		return f
	}
	f, ok := xs[from*10+to]
	if !ok {
		panic(fmt.Sprintf("%d::%d -> %d", from*10+to, from, to))
	}
	return f
}
func part2() {
	maxx := 0
	maxy := 0
	facing := 0 // facing to the right
	tiles := map[image.Point]*tile{}
	fmt.Println("d22")
	b, _ := os.ReadFile("INPUT")
	xs := strings.Split(string(b), "\n\n")
	board, path := xs[0], xs[1]
	grid := strings.Split(board, "\n")
	//fmt.Println(path)
	//fmt.Println(board)

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
	for _, v := range tiles {
		row, col := 0, 0
		x, y := v.position.X, v.position.Y
		if x < 50 {
			col = 1
		} else if x >= 50 && x < 100 {
			col = 2
		} else if x >= 100 && x < 150 {
			col = 3
		} else {
			panic("invalid x " + fmt.Sprintf("%d", x))
		}
		if y < 50 {
			row = 10
		} else if y >= 50 && y < 100 {
			row = 20
		} else if y >= 100 && y < 150 {
			row = 30
		} else if y >= 150 && y < 200 {
			row = 40
		} else {
			panic("invalid y")
		}
		side := 0
		switch row + col {
		case 12:
			side = 1
		case 13:
			side = 2
		case 22:
			side = 3
		case 31:
			side = 4
		case 32:
			side = 5
		case 41:
			side = 6
		default:
			panic("unknown side: " + fmt.Sprintf("%d", row+col))
		}
		v.side = side
		v.xoffset, v.yoffset = offsets(side)
		tiles[v.position] = v
	}

	//parse to link tiles
	fmt.Println(maxx, maxy)
	for k, v := range tiles {
		// l + r
		left := v.position.X - 1
		for {
			if left < 0 {

				pos := wrap(v.side, v.position, LEFT)
				printme(v.side, v.position, LEFT, pos, tiles)
				leftTile, ok := tiles[pos]
				if !ok {
					panic("messed up wrap left")
				}
				v.left = leftTile
				tiles[k] = v
				break
			}
			if leftTile, ok := tiles[image.Pt(left, v.position.Y)]; ok &&
				left != v.position.X {
				v.left = leftTile
				tiles[k] = v
				break
			}
			left--
		}
		right := v.position.X + 1
		for {
			if right > maxx {

				pos := wrap(v.side, v.position, RIGHT)
				printme(v.side, v.position, RIGHT, pos, tiles)
				rightTile, ok := tiles[pos]
				if !ok {
					panic("messed up wrap right")
				}
				v.right = rightTile
				tiles[k] = v
				break
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

				pos := wrap(v.side, v.position, UP)
				printme(v.side, v.position, UP, pos, tiles)
				upTile, ok := tiles[pos]
				if !ok {
					panic("messed up wrap up")
				}
				v.up = upTile
				tiles[k] = v
				break
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

				pos := wrap(v.side, v.position, DOWN)
				printme(v.side, v.position, DOWN, pos, tiles)
				downTile, ok := tiles[pos]
				if !ok {
					panic("messed up wrap down")
				}
				v.down = downTile
				tiles[k] = v
				break
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
		//	fmt.Printf("facing %d, position side %d %v\n", facing, current.position, current.side)
		switch facing {
		case 0:
			if current.right.value != '#' {
				if current.side != current.right.side {
					//fmt.Printf("side [%d] -> side [%d]\n", current.side, current.right.side)

				}
				facing = newfacing(facing, current.side, current.right.side)
				current = current.right
			}
			move(count - 1)
		case 1:
			if current.down.value != '#' {
				if current.side != current.down.side {
					//	fmt.Printf("side [%d] -> side [%d]\n", current.side, current.down.side)
				}
				facing = newfacing(facing, current.side, current.down.side)
				current = current.down
			}
			move(count - 1)
		case 2:
			if current.left.value != '#' {
				facing = newfacing(facing, current.side, current.left.side)
				current = current.left
			}
			move(count - 1)
		case 3:
			if current.up.value != '#' {
				facing = newfacing(facing, current.side, current.up.side)
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
					//	println(num)
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
			//	fmt.Printf("%c\n", path[p])
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
