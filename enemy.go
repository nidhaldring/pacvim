package main

import (
	"github.com/adam-lavrik/go-imath/ix"
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

}

func (e *Enemy) GetCurrentPos() (int, int) {
	return e.x, e.y
}

func (e *Enemy) Draw() {
	e.screen.SetContent(e.x, e.y, 'E', nil, EnemyTheme)
}

func calcMovementCost(fromX, fromY, toX, toY int) int {
	return ix.Abs(fromX-toX) + ix.Abs(fromY-toY)
}
