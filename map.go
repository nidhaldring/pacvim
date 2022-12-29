package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gdamore/tcell"
)

type Map struct {
	elements [][]*Element
	screen   tcell.Screen
}

func NewMap(screen tcell.Screen) (*Map, int) {
	mapLines := readMap("map.txt")
	elements := make([][]*Element, 0)
	eatableElementsNumber := 0

	for i := 0; i < len(mapLines); i++ {
		row := make([]*Element, 0)
		for j := 0; j < len(mapLines[i]); j++ {
			ch := rune(mapLines[i][j])
			elm := newElement(screen, ch, j, i)
			if elm.IsEatable() {
				eatableElementsNumber++
			}
			row = append(row, elm)
		}
		elements = append(elements, row)
	}

	return &Map{
		elements: elements,
		screen:   screen,
	}, eatableElementsNumber
}

func (m *Map) IsTraversable(x, y int) bool {
	elm := m.getElementAt(x, y)
	return elm != nil && elm.IsTraversable()
}

func (m *Map) Draw() {
	for _, elmRow := range m.elements {
		for _, elm := range elmRow {
			elm.Draw()
		}
	}
}

func (m *Map) getElementAt(x, y int) *Element {
	if y > -1 && y < len(m.elements) && x > -1 && x < len(m.elements[y]) {
		return m.elements[y][x]
	}
	return nil
}

func readMap(name string) []string {
	filePath := path.Join("maps", name)
	res, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("map %s was not found", filePath))
	}

	return strings.Split(string(res), "\n")
}
