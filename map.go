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

func (m *Map) GetNextMoveToGoal(startX, startY, goalX, goalY int) (int, int) {
	goalElm := m.getElementAt(goalX, goalY)

	stack := NewStack()
	visited := NewStack()

	visited.Push(&NodePath{m.getElementAt(startX, startY), nil})
	stack.Push(&NodePath{m.getElementAt(startX, startY), nil})
	for stack.IsNotEmpty() {
		nodePath := stack.Pop()
		visited.Push(nodePath)
		if nodePath.n == goalElm {
			return GetFirstMoveInPath(nodePath)
		}

		neighbors := m.getTraversableNeighbors(nodePath.n)
		for _, n := range neighbors {
			if !visited.Contains(n) {
				stack.Push(&NodePath{n, nodePath})
			}
		}

	}

	return -1, -1
}

func (m *Map) getElementAt(x, y int) *Element {
	if y > -1 && y < len(m.elements) && x > -1 && x < len(m.elements[y]) {
		return m.elements[y][x]
	}
	return nil
}

func (m *Map) getTraversableNeighbors(e *Element) []*Element {
	neighbors := make([]*Element, 0)

	up := m.getElementAt(e.x-1, e.y)
	if up != nil && up.IsTraversable() {
		neighbors = append(neighbors, up)
	}

	down := m.getElementAt(e.x+1, e.y)
	if down != nil && down.IsTraversable() {
		neighbors = append(neighbors, down)
	}

	left := m.getElementAt(e.x, e.y-1)
	if left != nil && left.IsTraversable() {
		neighbors = append(neighbors, left)
	}

	right := m.getElementAt(e.x, e.y+1)
	if right != nil && right.IsTraversable() {
		neighbors = append(neighbors, right)
	}

	return neighbors
}

func readMap(name string) []string {
	filePath := path.Join("maps", name)
	res, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("map %s was not found", filePath))
	}

	return strings.Split(string(res), "\n")
}
