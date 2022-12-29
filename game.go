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
	player                *Player
	enemy                 *Enemy
	gMap                  *Map
	screen                tcell.Screen
	eatenElementsNumber   int
	eatableElementsNumber int
}

func NewGame() *Game {
	screen := NewScreen()
	gMap, eatableElementsNumber := NewMap(screen)
	player := NewPlayer(screen, gMap)

	return &Game{
		status:                PLAYING,
		screen:                screen,
		gMap:                  gMap,
		player:                player,
		enemy:                 NewEnemy(screen, gMap, player),
		eatableElementsNumber: eatableElementsNumber,
	}
}

func (g *Game) Start() {
	for {
		g.handleEvents()
		g.draw()
		g.handleCollision()
	}
}

func (g *Game) draw() {
	g.screen.Clear()
	switch g.status {
	case PLAYING:
		g.gMap.Draw()
		g.player.Draw()
		g.enemy.Draw()
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
	g.player.Move(ev.Rune())
	g.enemy.Move()
}

func (g *Game) handleCollision() {
	g.handleCollisionWithMap()
	g.handleCollisionWithEnemies()
}

func (g *Game) handleCollisionWithMap() {
	playerX, playerY := g.player.GetCurrentPos()
	elementUnderPlayer := g.gMap.getElementAt(playerX, playerY)

	if elementUnderPlayer.IsDeadly() {
		g.lose()
	} else if elementUnderPlayer.IsEatable() {
		elementUnderPlayer.MarkAsEaten()
		g.eatenElementsNumber++
	}
}

func (g *Game) handleCollisionWithEnemies() {
	playerX, playerY := g.player.GetCurrentPos()
	enemyX, enemyY := g.enemy.GetCurrentPos()

	if enemyX == playerX && enemyY == playerY {
		g.lose()
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
	g.player.Hide()
	g.drawMessageToScreen("CONGRATS, YOU HAVE WON !")
}

func (g *Game) drawLosingScreen() {
	g.player.Hide()
	g.drawMessageToScreen("YOU HAVE LOST!")
}
