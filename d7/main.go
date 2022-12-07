package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name   string
	file   []*File
	dirs   []*Dir
	parent *Dir
}

func main() {
	fmt.Println("d7")
	file, err := os.Open("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tree := &Dir{}
	tree.parent = tree
	tree.name = "/"
	for scanner.Scan() {
		line := scanner.Text()
		tree = exec(tree, line)
	}
	tree = tree.goRoot()
	sizes := tree.dirSizes([]int{})

	fmt.Println(part1(sizes))
	fmt.Println(part2(sizes))

}

func part1(xs []int) int {
	total := 0
	for _, x := range xs {
		if x <= 100000 {
			total = total + x
		}
	}
	return total
}

func part2(xs []int) int {
	// /: 48381165 SAMPLE
	// /: 48748071 INPUT
	//
	//level := 30000000 - (70000000 - 48381165)
	level := 30000000 - (70000000 - 48748071)
	_ = level
	options := []int{}
	for _, x := range xs {
		if x >= level {
			options = append(options, x)
		}
	}
	min := 70000000
	for _, x := range options {
		if x < min {
			min = x
		}
	}
	return min
}

func exec(tree *Dir, cmd string) *Dir {
	switch {
	case strings.HasPrefix(cmd, "$ cd /"):
		tree = tree.goRoot()
		return tree
	case strings.HasPrefix(cmd, "$ cd .."):
		tree = tree.up()
		return tree
	case strings.HasPrefix(cmd, "$ cd"):
		var name string
		fmt.Sscanf(cmd, "$ cd %s", &name)
		tree = tree.cd(name)
	case strings.HasPrefix(cmd, "$ ls"):
	case strings.HasPrefix(cmd, "dir"):
		var name string
		fmt.Sscanf(cmd, "dir %s", &name)
		tree.addDir(name)
	default:
		var name string
		var size int
		fmt.Sscanf(cmd, "%d %s", &size, &name)
		tree.addFile(name, size)
	}
	return tree
}

func (dir *Dir) goRoot() *Dir {
	current := dir
	for {
		if current.name == "/" {
			return current
		}
		current = current.parent
	}
}

func (dir *Dir) size() int {
	total := 0
	for _, x := range dir.file {
		total = total + x.size
	}
	for _, x := range dir.dirs {
		total = total + x.size()
	}

	return total

}

func (dir *Dir) dirSizes(acc []int) []int {
	acc = append(acc, dir.size())
	for _, x := range dir.dirs {
		acc = x.dirSizes(acc)
	}
	return acc
}

func (dir *Dir) up() *Dir {
	current := dir
	current = current.parent
	return current
}

func (dir *Dir) cd(name string) *Dir {
	current := dir
	for _, x := range dir.dirs {
		if x.name == name {
			current = x
			return current
		}
	}
	newDir := &Dir{}
	newDir.name = name
	newDir.parent = current
	dir.dirs = append(dir.dirs, newDir)
	return newDir
}

func (dir *Dir) addDir(name string) {
	for _, x := range dir.dirs {
		if x.name == name {
			return
		}
	}
	newDir := &Dir{}
	newDir.name = name
	newDir.parent = dir
	dir.dirs = append(dir.dirs, newDir)
}

func (dir *Dir) addFile(name string, size int) {
	for _, x := range dir.file {
		if x.name == name {
			return
		}
	}
	newFile := &File{}
	newFile.name = name
	newFile.size = size
	dir.file = append(dir.file, newFile)
}
