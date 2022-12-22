package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func NewScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}

	if err := s.Init(); err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}

	s.SetStyle(DefTheme)

	return s

}
