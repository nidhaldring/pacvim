package main

import (
	"github.com/gdamore/tcell"
)

type ElementType int

const (
	DEADLY ElementType = iota
	EATABLE
	EATEN
	SPACE
	BLOCKING
)

type Element struct {
	x, y   int
	value  rune
	elType ElementType
	style  tcell.Style
	screen tcell.Screen
}

func newElement(screen tcell.Screen, ch rune, x, y int) *Element {
	elType := getEelementType(ch)
	style := getEelementStyle(elType)
	return &Element{
		x:      x,
		y:      y,
		elType: elType,
		style:  style,
		screen: screen,
		value:  ch,
	}
}

func (e *Element) Intersect(x, y int) bool {
	return x == e.x && y == e.y
}

func (e *Element) CanBeEaten() bool {
	return e.elType == EATABLE
}

func (e *Element) MarkAsEaten() {
	e.elType = EATEN
	e.style = EatenTheme
}

func (e *Element) Draw() {
	e.screen.SetContent(e.x, e.y, e.value, nil, e.style)
}

func getEelementType(ch rune) ElementType {
	switch ch {
	case '~':
		return DEADLY
	case '#':
		return BLOCKING
	case ' ':
		return SPACE
	default:
		return EATABLE
	}
}

func getEelementStyle(elType ElementType) tcell.Style {
	switch elType {
	case DEADLY:
		return DeadlyTheme
	case EATEN:
		return EatenTheme
	default:
		return DefTheme
	}
}
