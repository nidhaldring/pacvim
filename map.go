package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gdamore/tcell"
)

type Map struct {
	elements              [][]*Element
	eatableElementsNumber int
	cursor                *Cursor
	screen                tcell.Screen
}

func NewMap(screen tcell.Screen) *Map {
	mapLines := readMap("map.txt")
	elements := make([][]*Element, 0)
	eatableElementsNumber := 0

	for i := 0; i < len(mapLines); i++ {
		row := make([]*Element, 0)
		for j := 0; j < len(mapLines[i]); j++ {
			ch := rune(mapLines[i][j])
			elm := newElement(screen, ch, j, i)
			if elm.CanBeEaten() {
				eatableElementsNumber++
			}
			row = append(row, elm)
		}
		elements = append(elements, row)
	}

	return &Map{
		elements:              elements,
		eatableElementsNumber: eatableElementsNumber,
		screen:                screen,
	}
}

func (m *Map) GetElementAt(x, y int) *Element {
	return m.elements[y][x]
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
