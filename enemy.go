package main

import "github.com/gdamore/tcell"

type Enemy struct {
	x, y   int
	screen tcell.Screen
	gMap   *Map
	player *Player
}

func (e *Enemy) Move() {
	cursorX, cursorY := e.player.GetCurrentPos()
	if e.x != cursorX {
		if e.x > cursorX && e.gMap.IsTraversable(e.x-1, e.y) {
			e.x--
		} else if e.gMap.IsTraversable(e.x+1, e.y) {
			e.x++
		}
	} else if e.y != cursorY {
		if e.y > cursorX && e.gMap.IsTraversable(e.x, e.y-1) {
			e.y--
		} else if e.gMap.IsTraversable(e.x, e.y+1) {
			e.y++
		}
	}
}

func (e *Enemy) Draw() {
	e.screen.SetContent(e.x, e.y, 'E', nil, EnemyTheme)
}
