package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type tickMsg struct{ time.Time }

func tick() tea.Cmd {
	return tea.Tick(calcWordDuration(wpm), func(t time.Time) tea.Msg {
		return tickMsg{t}
	})
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tick(),
		tea.EnterAltScreen,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case " ", "p":
			m.paused = !m.paused

		case "right":
			wpm += 10

		case "left":
			if wpm-10 > 0 {
				wpm -= 10
			}
		}

	case tickMsg:
		if m.cursor+1 >= len(m.words) {
			return m, tea.Quit
		}

		if !m.paused {
			if wordEndsWithPunctuation(m.words[m.cursor]) {
				if m.endOfSentence == 1 {
					m.cursor += 1
					m.endOfSentence = 0
				} else {
					m.endOfSentence += 1
				}
			} else {
				m.cursor += 1
			}

		}

		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Failed to get terminal size:", err)
		return ""
	}

	status := fmt.Sprint(wpm)

	if m.paused {
		status = "PAUSED"
	}

	padding := "\n\n\n\n"
	t := "\n" + m.source + " - " + status + padding + m.words[m.cursor]

	text := lipgloss.NewStyle().
		Width(w).
		Height(h).
		Align(lipgloss.Center).
		Render(t)

	return text
}
