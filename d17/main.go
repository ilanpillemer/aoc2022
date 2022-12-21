package main

import (
	"fmt"
	"image"
	. "image"
	"os"
)

var pieces = []map[Point]bool{}
var stopped = true
var chamber = map[Point]bool{}
var top = 1
var x = 0
var turn = 0
var blow int

func main() {
	b, _ := os.ReadFile("INPUT")
	moves := string(b)
	modular := len(moves)
	//	fmt.Println(moves)
	t := 0
	p := 0
	flr := floor()
	var a map[image.Point]bool

	addChamber(flr)
	for {
		if stopped {
			stopped = false
			a = getPiece(p)
			//	fmt.Println("got piece", a)
			p++
		}

		if blow%2 != 0 {
			nextmove := moves[turn%modular]
			turn++
			blow++
			a = dropPiece(a, nextmove)
		} else {
			a = movedown(a)
			blow++
		}
		if p == 2023 {
			fmt.Println("part1", top)
			break

		}
		t = t + 1
		//	fmt.Println(p, top)
	}
	// fmt.Println(chamber)
}

func addChamber(piece map[image.Point]bool) {
	for k := range piece {
		chamber[k] = true
	}
}

func dropPiece(piece map[image.Point]bool, move byte) map[image.Point]bool {
	//	fmt.Println("dropping piece")
	switch move {
	case '<':
		piece = moveleft(piece)
	case '>':
		piece = moveright(piece)
	default:
		panic("oy")

	}

	return piece
}

func canMoveLeft(piece map[image.Point]bool) bool {

	for k := range piece {
		if chamber[Pt(k.X-1, k.Y)] || k.X == 1 {
			return false
		}
	}
	return true
}

func canMoveRight(piece map[image.Point]bool) bool {
	for k := range piece {
		if chamber[Pt(k.X+1, k.Y)] || k.X == 7 {
			return false
		}
	}
	return true
}

func canMoveDown(piece map[image.Point]bool) bool {
	for k := range piece {
		if chamber[Pt(k.X, k.Y-1)] {
			return false
		}
	}
	return true
}

func moveleft(piece map[image.Point]bool) map[image.Point]bool {
	//	fmt.Println("moving left")
	newPiece := map[image.Point]bool{}

	if canMoveLeft(piece) {
		for k := range piece {
			newPiece[Pt(k.X-1, k.Y)] = true
		}
		//	fmt.Println(newPiece)
		return newPiece
	}
	return piece
}

func movedown(piece map[image.Point]bool) map[image.Point]bool {
	//	fmt.Println("moving down")
	//	fmt.Println(piece)
	newPiece := map[image.Point]bool{}
	if canMoveDown(piece) {
		for k := range piece {
			newPiece[Pt(k.X, k.Y-1)] = true
		}
		//	fmt.Println(newPiece)
		return newPiece
	} else {
		newTop := 0
		for k := range piece {
			newTop = k.Y
			if newTop > top {
				top = newTop
			}
		}
		//	fmt.Println("stopped")
		addChamber(piece)
		stopped = true
	}
	return piece
}

func moveright(piece map[image.Point]bool) map[image.Point]bool {
	//	fmt.Println("moving right")
	//	fmt.Println(piece)
	newPiece := map[image.Point]bool{}
	if canMoveRight(piece) {
		for k := range piece {
			newPiece[Pt(k.X+1, k.Y)] = true
		}
		//	fmt.Println(newPiece)
		return newPiece
	}
	//	fmt.Println(piece)
	return piece
}

// ####
//
// .#.
// ###
// .#.
//
// ..#
// ..#
// ###
//
// #
// #
// #
// #
//
// ##
// ##
// The tall, vertical chamber is exactly seven units wide. Each rock appears so that its left edge is two units away from the left wall and its bottom edge is three units above the highest rock in the room (or the floor, if there isn't one).

func getPiece(t int) map[Point]bool {
	p := t % 5
	//	fmt.Println("getting piece", p)
	switch p {
	case 0:
		return newPiece1(top)
	case 1:
		return newPiece2(top)
	case 2:
		return newPiece3(top)
	case 3:
		return newPiece4(top)
	case 4:
		return newPiece5(top)
	}

	panic("oy")

}

func floor() map[Point]bool {
	piece := map[Point]bool{}
	piece[Pt(0, 0)] = true
	piece[Pt(1, 0)] = true
	piece[Pt(2, 0)] = true
	piece[Pt(3, 0)] = true
	piece[Pt(4, 0)] = true
	piece[Pt(5, 0)] = true
	piece[Pt(6, 0)] = true
	piece[Pt(7, 0)] = true
	piece[Pt(8, 0)] = true
	return piece
}

func newPiece1(top int) map[Point]bool {
	piece := map[Point]bool{}
	x := 3
	y := top + 4
	piece[Pt(x, y)] = true
	piece[Pt(x+1, y)] = true
	piece[Pt(x+2, y)] = true
	piece[Pt(x+3, y)] = true
	return piece
}

func newPiece2(top int) map[Point]bool {
	piece := map[Point]bool{}
	x := 3
	y := top + 4
	piece[Pt(x+1, y+2)] = true
	piece[Pt(x, y+1)] = true
	piece[Pt(x+1, y+1)] = true
	piece[Pt(x+2, y+1)] = true
	piece[Pt(x+1, y)] = true
	return piece
}

func newPiece3(top int) map[Point]bool {
	piece := map[Point]bool{}
	x := 3
	y := top + 4
	piece[Pt(x, y)] = true
	piece[Pt(x+1, y)] = true
	piece[Pt(x+2, y)] = true
	piece[Pt(x+2, y+1)] = true
	piece[Pt(x+2, y+2)] = true

	return piece
}

func newPiece4(top int) map[Point]bool {
	piece := map[Point]bool{}
	x := 3
	y := top + 4
	piece[Pt(x, y+3)] = true
	piece[Pt(x, y+2)] = true
	piece[Pt(x, y+1)] = true
	piece[Pt(x, y)] = true
	return piece
}

func newPiece5(top int) map[Point]bool {
	//	fmt.Println("PIECE5")
	piece := map[Point]bool{}
	x := 3
	y := top + 4
	piece[Pt(x, y+1)] = true
	piece[Pt(x, y)] = true
	piece[Pt(x+1, y+1)] = true
	piece[Pt(x+1, y)] = true
	return piece
}
