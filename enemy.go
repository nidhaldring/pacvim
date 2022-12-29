package main

import "github.com/gdamore/tcell"

type Enenmy struct {
	x, y   int
	screen tcell.Screen
}

func (e *Enenmy) Draw() {
	e.screen.SetContent(e.x, e.y, 'E', nil, EnemyTheme)
}
