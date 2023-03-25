package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nsf/termbox-go"
)

var wpm = 200

func main() {
	wpmFlag := flag.Int("w", 0, "words per minute (default 200)")
	filepathFlag := flag.String("f", "", "read path/to/file")
	pausedFlag := flag.Bool("p", false, "start paused")
	flag.Parse()

	wpmArg := *wpmFlag
	filepath := *filepathFlag
	paused := *pausedFlag
	input := ""
	source := ""

	c := readConfig()

	if wpmArg != 0 {
		wpm = wpmArg
	} else {
		wpm = c.Wpm
	}

	if isPiped() {
		text := readPipedInput(os.Stdin, os.Stdout)
		input = text
		source = "STDIN"
	} else {
		text := readFileInput(filepath)
		input = text
		source = filepath
	}

	if source == "" {
		flag.Usage()
		os.Exit(1)
	}

	model := createModel(input, source, paused)

	err := termbox.Init()
	if err != nil {
		fmt.Println("Error initializing termbox:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred initiating the app: %v", err)
		os.Exit(1)
	}
}

type model struct {
	words         []string
	cursor        int
	paused        bool
	source        string
	endOfSentence int
}

func createModel(inp string, source string, paused bool) model {
	return model{
		words:         splitInput(inp),
		cursor:        0,
		paused:        paused,
		source:        source,
		endOfSentence: 0,
	}
}
