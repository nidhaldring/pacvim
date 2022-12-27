package main

import "github.com/gdamore/tcell"

type Cursor struct {
	x, y   int
	screen tcell.Screen
	gMap   *Map
}

func NewCursor(s tcell.Screen) *Cursor {
	return &Cursor{
		x:      1,
		y:      1,
		screen: s,
	}
}

func (c *Cursor) SetMap(gMap *Map) {
	c.gMap = gMap
}

func (c *Cursor) Draw() {
	c.screen.ShowCursor(c.x, c.y)
}

func (c *Cursor) Hide() {
	c.screen.HideCursor()
}

func (c *Cursor) HandleEvents(ev *tcell.EventKey) {
	switch ev.Rune() {
	case 'h':
		c.moveLeft()
	case 'j':
		c.moveDown()
	case 'k':
		c.moveUp()
	case 'l':
		c.moveRight()

	}
}

func (c *Cursor) moveUp() {
	if !(c.gMap.IsTraversable(c.x, c.y-1)) {
		c.screen.Beep()
	} else {
		c.y--
	}

}

func (c *Cursor) moveDown() {
	if !(c.gMap.IsTraversable(c.x, c.y+1)) {
		c.screen.Beep()
	} else {
		c.y++
	}
}

func (c *Cursor) moveLeft() {
	if !(c.gMap.IsTraversable(c.x-1, c.y)) {
		c.screen.Beep()
	} else {
		c.x--
	}
}

func (c *Cursor) moveRight() {
	if !(c.gMap.IsTraversable(c.x+1, c.y)) {
		c.screen.Beep()
	} else {
		c.x++
	}
}
