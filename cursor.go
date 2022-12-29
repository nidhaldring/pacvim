package main

import "github.com/gdamore/tcell"

type Cursor struct {
	x, y   int
	screen tcell.Screen
}

func NewCursor(s tcell.Screen) *Cursor {
	return &Cursor{
		x:      1,
		y:      1,
		screen: s,
	}
}

func (c *Cursor) Draw() {
	c.screen.ShowCursor(c.x, c.y)
}

func (c *Cursor) Hide() {
	c.screen.HideCursor()
}

func (c *Cursor) Beep() {
	c.screen.Beep()
}

func (c *Cursor) GetCurrentPos() (int, int) {
	return c.x, c.y
}

func (c *Cursor) SetPos(x, y int) {
	c.x, c.y = x, y
}
