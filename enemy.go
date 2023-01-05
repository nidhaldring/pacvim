package main

import (
	"github.com/gdamore/tcell"
)

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
	goalX, goalY := e.player.GetCurrentPos()
	nextX, nextY := e.gMap.GetNextMoveToGoal(e.x, e.y, goalX, goalY)

	if nextX != -1 && nextY != -1 {
		e.x = nextX
		e.y = nextY
	}
}

func (e *Enemy) GetCurrentPos() (int, int) {
	return e.x, e.y
}

func (e *Enemy) Draw() {
	e.screen.SetContent(e.x, e.y, 'E', nil, EnemyTheme)
}
