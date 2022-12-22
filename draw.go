package main

import "github.com/gdamore/tcell"

func DrawToScreen(s tcell.Screen, x, y int, ch rune) {
	s.SetContent(x, y, ch, nil, DefTheme)
}
