package main

import (
	"github.com/gdamore/tcell"
)

type Player struct {
	x, y   int
	screen tcell.Screen
	gMap   *Map
}

func NewPlayer(s tcell.Screen, gMap *Map) *Player {
	return &Player{
		x:      1,
		y:      1,
		screen: s,
		gMap:   gMap,
	}
}

func (p *Player) Move(key rune) {
	switch key {
	case UP_KEY:
		if !p.gMap.IsTraversable(p.x, p.y-1) {
			p.beep()
			return
		}
		p.y--
	case DOWN_KEY:
		if !p.gMap.IsTraversable(p.x, p.y+1) {
			p.beep()
			return
		}
		p.y++
	case LEFT_KEY:
		if !p.gMap.IsTraversable(p.x-1, p.y) {
			p.beep()
			return
		}
		p.x--
	case RIGHT_KEY:
		if !p.gMap.IsTraversable(p.x+1, p.y) {
			p.beep()
			return
		}
		p.x++
	}
}

func (p *Player) Draw() {
	p.screen.ShowCursor(p.x, p.y)
}

func (p *Player) Hide() {
	p.screen.HideCursor()
}

func (p *Player) GetCurrentPos() (int, int) {
	return p.x, p.y
}

func (p *Player) beep() {
	p.screen.Beep()
}
