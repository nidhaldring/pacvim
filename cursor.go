package main

import "github.com/gdamore/tcell"

type Cursor struct {
	x, y   int
	screen tcell.Screen
}

func NewCursor(s tcell.Screen) *Cursor {
	return &Cursor{
		screen: s,
	}
}

func (c *Cursor) Draw() {
	c.screen.ShowCursor(c.x, c.y)
}

func (c *Cursor) HandleEvents(ev *tcell.EventKey) {
	switch ev.Rune() {
	case 'h':
		c.moveRight()
	case 'j':
		c.moveDown()
	case 'k':
		c.moveUp()
	case 'l':
		c.moveLeft()
	}
}

func (c *Cursor) moveUp() {
	if c.y == 0 {
		c.screen.Beep()
	} else {
		c.y--
	}
}

func (c *Cursor) moveDown() {
	c.y++
}

func (c *Cursor) moveLeft() {
	c.x--
}

func (c *Cursor) moveRight() {
	c.x++
}
