package main

import "github.com/gdamore/tcell"

type Enemy struct {
	x, y   int
	screen tcell.Screen
	gMap   *Map
	player *Player
}

func NewEnemy(screen tcell.Screen, gMap *Map, player *Player) *Enemy {
	return &Enemy{
		x:      8,
		y:      1,
		screen: screen,
		gMap:   gMap,
		player: player,
	}
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

func (e *Enemy) GetCurrentPos() (int, int) {
	return e.x, e.y
}

func (e *Enemy) Draw() {
	e.screen.SetContent(e.x, e.y, 'E', nil, EnemyTheme)
}
