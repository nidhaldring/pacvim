package main

import "github.com/gdamore/tcell"

var (
	DefTheme    = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	EatenTheme  = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)
	DeadlyTheme = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	EnemyTheme  = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	ScoreTheme  = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	UP_KEY      = 'k'
	DOWN_KEY    = 'j'
	LEFT_KEY    = 'h'
	RIGHT_KEY   = 'l'
)
