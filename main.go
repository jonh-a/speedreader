package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var wpm = DEFAULT_WPM

func main() {
	wpmFlag := flag.Int("w", 0, "words per minute (default 200)")
	filepathFlag := flag.String("f", "", "read path/to/file")
	pausedFlag := flag.Bool("p", false, "start paused")
	highlightORPFlag := flag.Bool("o", DEFAULT_HIGHLIGHT_ORP, "highlight ORP")
	printVrsionFlag := flag.Bool("v", false, "print version")
	flag.Parse()

	wpmArg := *wpmFlag
	filepath := *filepathFlag
	paused := *pausedFlag
	highlightORP := *highlightORPFlag
	printVersion := *printVrsionFlag
	input := ""
	source := ""

	if printVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	c := readConfig()

	if wpmArg != 0 {
		wpm = wpmArg
	} else if c.Wpm != 0 {
		wpm = c.Wpm
	} else {
		wpm = DEFAULT_WPM
	}

	fmt.Print(wpm)

	if !highlightORP {
		highlightORP = c.HighlightORP
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

	model := createModel(input, source, paused, highlightORP)

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
	highlightORP  bool
	endOfSentence int
}

func createModel(inp string, source string, paused bool, highlightORP bool) model {
	return model{
		words:         splitInput(inp),
		cursor:        0,
		paused:        paused,
		source:        source,
		highlightORP:  highlightORP,
		endOfSentence: 0,
	}
}
