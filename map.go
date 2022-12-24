package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gdamore/tcell"
)

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
			row = append(row, newElement(screen, ch, j, i))
		}
		elements = append(elements, row)
	}

	return &Map{
		elements: elements,
	}
}

func (m *Map) IsTraversable(x, y int) bool {
	if x < 0 || y < 0 || x > len(m.elements) || y > len(m.elements[x]) {
		return false
	}

	elm := m.elements[x][y]
	return elm.elType == BLOCKING
}

func (m *Map) Draw() {
	for _, elmRow := range m.elements {
		for _, elm := range elmRow {
			elm.Draw()
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
	e.screen.SetContent(e.x, e.y, e.value, nil, DefTheme)
}
