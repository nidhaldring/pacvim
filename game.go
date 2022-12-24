package main

import (
	"os"

	"github.com/gdamore/tcell"
)

type Game struct {
	cursor *Cursor
	gMap   *Map
	screen tcell.Screen
}

func NewGame() *Game {
	screen := NewScreen()
	return &Game{
		screen: screen,
		gMap:   NewMap(screen),
		cursor: NewCursor(screen),
	}
}

func (g *Game) Start() {
	for {
		g.screen.Clear()
		g.Draw()
		g.handleEvents()
		g.screen.Show()
	}
}

func (g *Game) Draw() {
	g.gMap.Draw()
	g.cursor.Draw()
}

func (g *Game) quit() {
	g.screen.Fini()
	os.Exit(0)
}

// func (g *Game) handleMovementKeys(key rune) {
// 	// switch key {
// 	// case 'h':
// 	// 	g.cursor.MoveLeft()
// 	// case 'l':
// 	// 	g.cursor.MoveRight()
// 	// case 'j':
// 	// 	g.cursor.MoveDown()
// 	// case 'k':
// 	// 	g.cursor.MoveUp()
// 	// }
// }

func (g *Game) handleEvents() {
	ev := g.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			g.quit()
		} else {
			g.cursor.HandleEvents(ev)
		}

	}

}
