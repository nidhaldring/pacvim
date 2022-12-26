package main

import "github.com/gdamore/tcell"

var (
	DefTheme   = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	EatenTheme = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)
)
