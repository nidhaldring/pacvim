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

func (c *Cursor) MoveUp() {
	if c.y == 0 {
		c.screen.Beep()
	} else {
		c.y--
	}
}

func (c *Cursor) MoveDown() {
	c.y++
}

func (c *Cursor) MoveLeft() {
	c.x--
}

func (c *Cursor) MoveRight() {
	c.x++
}

func (c *Cursor) GetPos() (int, int) {
	return c.x, c.y
}
