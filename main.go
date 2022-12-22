package main

// func quit(s tcell.Screen) {
// 	s.Fini()
// 	os.Exit(0)
// }

// func drawToScreen(s tcell.Screen, x, y int, char rune) {
// 	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
// 	s.SetContent(x, y, char, nil, defStyle)
// }

// func readMap(name string) []string {
// 	filePath := path.Join("maps", name)
// 	data, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		panic(fmt.Sprintf("file %s not found !", name))
// 	}

// 	return strings.Split(string(data), "\n")
// }

// func drawMap(s tcell.Screen, name string) {
// 	mapText := readMap("map.txt")
// 	for x, line := range mapText {
// 		y := 0
// 		for _, c := range line {
// 			drawToScreen(s, y, x, c)
// 			y++
// 		}
// 	}

// }

func main() {
	g := NewGame()

	g.Start()
}
