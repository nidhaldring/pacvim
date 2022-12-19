package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gdamore/tcell"
)

func quit(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}

func drawToScreen(s tcell.Screen, x, y int, char rune) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetContent(x, y, char, nil, defStyle)
}

func readMap(name string) []string {
	filePath := path.Join("maps", name)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("file %s not found !", name))
	}

	return strings.Split(string(data), "\n")
}

func drawMap(s tcell.Screen, name string) {
	mapText := readMap("map.txt")
	for x, line := range mapText {
		y := 0
		for _, c := range line {
			drawToScreen(s, y, x, c)
			y++
		}
	}

}

func handleCusorMovements(s tcell.Screen, pos *CursorPos, key *tcell.EventKey) {
	switch key.Rune() {
	case 'h':
		pos.x--
	case 'j':
		pos.y++
	case 'k':
		pos.y--
	case 'l':
		pos.x++
	}
}

func handleEvents(s tcell.Screen, pos *CursorPos) {
	ev := s.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			quit(s)
		} else {
			handleCusorMovements(s, pos, ev)
		}

	}
}

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	if err := s.Init(); err != nil {
		fmt.Printf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	s.SetStyle(defStyle)

	cursorPos := CursorPos{0, 0}

	for {
		s.Clear()
		drawMap(s, "map0.txt")
		s.ShowCursor(cursorPos.x, cursorPos.y)
		s.Show()
		handleEvents(s, &cursorPos)
	}

}
