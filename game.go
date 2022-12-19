package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

type Game struct {
	cusor  Cursor
	theme  tcell.Style
	screen tcell.Screen
}

func NewGame() Game {
	return Game{
		theme: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
	}
}

func (g *Game) Start() {

	g.initScreen()

	for {
		g.screen.Clear()

		g.screen.Show()
	}

}

func (g *Game) quit() {
	g.screen.Fini()
	os.Exit(0)
}

func (g *Game) initScreen() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	if err := s.Init(); s != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	s.SetStyle(g.theme)

	g.screen = s
}

func (g *Game) handleMovementKeys(key rune) {
	switch key {
	case 'h':
		g.cusor.MoveLeft()
	case 'l':
		g.cusor.MoveRight()
	case 'j':
		g.cusor.MoveDown()
	case 'k':
		g.cusor.MoveUp()
	}
}

func (g *Game) handleEvents() {
	ev := g.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			g.quit()
		}
		ev.Rune()
	}

}
