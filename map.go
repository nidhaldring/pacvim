package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gdamore/tcell"
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
	screen tcell.Screen
}

func newElement(screen tcell.Screen, ch rune, x, y int) Element {
	var elType ElementType
	switch ch {
	case '~':
		elType = DEADLY
	case '#':
		elType = BLOCKING
	default:
		elType = EATABLE
	}

	return Element{
		x:      x,
		y:      y,
		elType: elType,
		screen: screen,
		value:  ch,
	}
}

func (e *Element) Draw() {
	e.screen.SetContent(e.y, e.x, e.value, nil, DefTheme)
}

type Map struct {
	elements [][]Element
}

func NewMap(screen tcell.Screen) *Map {
	mapLines := readMap("map.txt")
	elements := make([][]Element, 0)

	for i := 0; i < len(mapLines); i++ {
		row := make([]Element, 0)
		for j := 0; j < len(mapLines[i]); j++ {
			ch := rune(mapLines[i][j])
			row = append(row, newElement(screen, ch, i, j))
		}
		elements = append(elements, row)
	}

	return &Map{
		elements: elements,
	}
}

func (m *Map) Draw() {
	for i := 0; i < len(m.elements); i++ {
		for j := 0; j < len(m.elements[i]); j++ {
			m.elements[i][j].Draw()
		}
	}
}

func readMap(name string) []string {
	filePath := path.Join("maps", name)
	res, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("map %s was not found", filePath))
	}

	return strings.Split(string(res), "\n")
}
