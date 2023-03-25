package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
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
