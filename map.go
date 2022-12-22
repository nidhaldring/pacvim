package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

type ElementType int

const (
	DEADLY ElementType = iota
	EATABLE
	BLOCKING
)

type Element struct {
	x, y   int
	value  rune
	elType ElementType
}

func newElement(ch rune, x, y int) *Element {
	var elType ElementType
	switch ch {
	case '~':
		elType = DEADLY
	case '*':
		elType = BLOCKING
	default:
		elType = EATABLE
	}

	return &Element{
		x:      x,
		y:      y,
		elType: elType,
	}
}

type Map struct {
	elements [][]Element
}

func NewMap() *Map {
	mapLines := readMap("map.txt")
	elements := make([][]Element, 0)

	for i := 0; i < len(mapLines); i++ {
		row := make([]Element, 0)
		for j := 0; j < len(mapLines[i]); j++ {
			ch := rune(mapLines[i][j])
			row = append(row, *newElement(ch, i, j))
		}
		elements = append(elements, row)
	}

	return &Map{
		elements: elements,
	}
}

func readMap(name string) []string {
	filePath := path.Join("maps", name)
	res, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("map %s was not found", filePath))
	}

	return strings.Split(string(res), "\n")
}
