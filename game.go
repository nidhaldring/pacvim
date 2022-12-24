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
	gMap := NewMap(screen)
	return &Game{
		screen: screen,
		gMap:   gMap,
		cursor: NewCursor(screen, gMap),
	}
}

func (g *Game) Start() {
	for {
		g.handleEvents()
		g.Draw()
	}
}

func (g *Game) Draw() {
	g.screen.Clear()
	g.gMap.Draw()
	g.cursor.Draw()
	g.screen.Show()
}

func (g *Game) quit() {
	g.screen.Fini()
	os.Exit(0)
}

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
