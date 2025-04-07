package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func isPiped() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readPipedInput(r io.Reader, w io.Writer) string {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}

	return text
}

func readFileInput(fp string) string {
	file, err := ioutil.ReadFile(fp)
	if err != nil {
		return ""
	}
	return string(file)
}

func calcWordDuration(wpm int) time.Duration {
	wordsPerMinute := float64(wpm)
	secondsPerMinute := 60.0
	wordDuration := time.Duration((secondsPerMinute/wordsPerMinute)*1000) * time.Millisecond
	return wordDuration
}

func splitInput(input string) []string {
	words := strings.Fields(input)
	return words
}

func wordEndsWithPunctuation(w string) bool {
	punctuation := []string{".", ",", "!", "?", ":", ";", "-", "\""}

	for _, p := range punctuation {
		if strings.HasSuffix(w, p) {
			return true
		}
	}
	return false
}

func findORP(word string) int {
	if len(word) > 13 {
		return 4
	}
	return []int{0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3}[len(word)]
}

func styleMiddleChar(w string, highlight bool) string {
	if len(w) == 0 {
		return ""
	}

	if len(w) == 1 {
		padding := 6
		if highlight {
			return strings.Repeat(" ", padding) + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(w) + strings.Repeat(" ", padding)
		}
		return strings.Repeat(" ", padding) + w + strings.Repeat(" ", padding)
	}

	orp := findORP(w)
	firstChunk := w[:orp]
	orpChar := string(w[orp])
	secondChunk := w[orp+1:]

	var styledORP string
	if highlight {
		styledORP = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(orpChar)
	} else {
		styledORP = orpChar
	}

	centerPos := 6
	beforeORP := len(firstChunk)
	afterORP := len(secondChunk)

	leftPadding := centerPos - beforeORP
	rightPadding := 13 - (centerPos + 1 + afterORP)

	if leftPadding < 0 {
		leftPadding = 0
	}
	if rightPadding < 0 {
		rightPadding = 0
	}

	paddedWord := strings.Repeat(" ", leftPadding) + firstChunk + styledORP + secondChunk + strings.Repeat(" ", rightPadding)

	return paddedWord
}
