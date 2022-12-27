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
	eatenElementsNumber   int
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

func (m *Map) setCursor(cursor *Cursor) {
	m.cursor = cursor
}

func (m *Map) IsTraversable(x, y int) bool {
	if x < 0 || y < 0 || y > len(m.elements) || x > len(m.elements[y]) {
		return false
	}

	elm := m.elements[y][x]
	return elm.elType != BLOCKING
}

func (m *Map) Draw() {
	if m.eatenElementsNumber == m.eatableElementsNumber {
		m.DrawWinningMap()
		m.cursor.Hide()
		return
	}

	m.DrawMapElements()
	m.DrawScore()
}

func (m *Map) DrawMapElements() {
	for _, elmRow := range m.elements {
		for _, elm := range elmRow {
			if elm.CanBeEaten() && elm.Intersect(m.cursor.x, m.cursor.y) {
				m.eatenElementsNumber++
				elm.MarkAsEaten()
			}
			elm.Draw()
		}
	}
}

func (m *Map) DrawScore() {
	score := fmt.Sprintf("%d/%d", m.eatenElementsNumber, m.eatableElementsNumber)
	_, screenY := m.screen.Size()
	for i, c := range score {
		m.screen.SetContent(i, screenY-10, c, nil, ScoreTheme)
	}
}

func (m *Map) DrawWinningMap() {
	msg := "CONGRATS, YOU'VE WON !"
	for i, c := range msg {
		m.screen.SetContent(i, 0, c, nil, DefTheme)
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
