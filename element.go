package main

import "github.com/gdamore/tcell"

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
