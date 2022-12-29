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
	elType := getElementType(ch)
	style := getElementStyle(elType)
	return &Element{
		x:      x,
		y:      y,
		elType: elType,
		style:  style,
		screen: screen,
		value:  ch,
	}
}

func (e *Element) IsTraversable() bool {
	return e.elType != BLOCKING
}

func (e *Element) IsDeadly() bool {
	return e.elType == DEADLY
}

func (e *Element) IsEatable() bool {
	return e.elType == EATABLE
}

func (e *Element) MarkAsEaten() {
	e.elType = EATEN
	e.style = EatenTheme
}

func (e *Element) Draw() {
	e.screen.SetContent(e.x, e.y, e.value, nil, e.style)
}

func getElementType(ch rune) ElementType {
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

func getElementStyle(elType ElementType) tcell.Style {
	switch elType {
	case DEADLY:
		return DeadlyTheme
	case EATEN:
		return EatenTheme
	default:
		return DefTheme
	}
}
