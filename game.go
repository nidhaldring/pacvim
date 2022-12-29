package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

type GameStatus int

const (
	PLAYING = iota
	LOST
	WON
)

type Game struct {
	status                GameStatus
	cursor                *Cursor
	gMap                  *Map
	screen                tcell.Screen
	eatenElementsNumber   int
	eatableElementsNumber int
}

func NewGame() *Game {
	screen := NewScreen()
	gMap, eatableElementsNumber := NewMap(screen)

	return &Game{
		status:                PLAYING,
		screen:                screen,
		gMap:                  gMap,
		cursor:                NewCursor(screen),
		eatableElementsNumber: eatableElementsNumber,
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
	switch g.status {
	case PLAYING:
		g.gMap.Draw()
		g.cursor.Draw()
		g.drawScore()
	case WON:
		g.drawWinningScreen()
	case LOST:
		g.drawLosingScreen()
	}
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
			g.handleMovementEvents(ev)
		}
	}

}

func (g *Game) handleMovementEvents(ev *tcell.EventKey) {
	curX, curY := g.cursor.GetCurrentPos()
	switch ev.Rune() {
	case UP_KEY:
		g.handleCursorMovements(curX, curY-1)
	case DOWN_KEY:
		g.handleCursorMovements(curX, curY+1)
	case LEFT_KEY:
		g.handleCursorMovements(curX-1, curY)
	case RIGHT_KEY:
		g.handleCursorMovements(curX+1, curY)
	}

}

func (g *Game) handleCursorMovements(x, y int) {
	cursor := g.cursor
	elm := g.gMap.GetElementAt(x, y)
	switch elm.elType {
	case DEADLY:
		g.lose()
	case BLOCKING:
		cursor.Beep()
	default:
		cursor.SetPos(x, y)
		if elm.elType == EATABLE {
			elm.MarkAsEaten()
			g.eatenElementsNumber++
			if g.eatableElementsNumber == g.eatenElementsNumber {
				g.win()
			}
		}
	}
}

func (g *Game) drawScore() {
	screen := g.screen
	score := fmt.Sprintf("%d/%d", g.eatenElementsNumber, g.eatableElementsNumber)
	_, screenY := screen.Size()
	for i, c := range score {
		screen.SetContent(i, screenY-10, c, nil, ScoreTheme)
	}
}

func (g *Game) win() {
	g.status = WON
}

func (g *Game) lose() {
	g.status = LOST
}

func (g *Game) drawMessageToScreen(msg string) {
	for i, c := range msg {
		g.screen.SetContent(i, 0, c, nil, DefTheme)
	}
}

func (g *Game) drawWinningScreen() {
	g.cursor.Hide()
	g.drawMessageToScreen("CONGRATS, YOU HAVE WONE !")
}

func (g *Game) drawLosingScreen() {
	g.cursor.Hide()
	g.drawMessageToScreen("YOU HAVE LOST!")
}
