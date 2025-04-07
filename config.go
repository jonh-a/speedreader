package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Wpm          int  `yaml:"wpm"`
	HighlightORP bool `yaml:"highlight_orp"`
}

func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	dir := usr.HomeDir
	return filepath.Join(dir, ".speedreaderrc")
}

func readConfig() Config {
	c := getConfigPath()
	file, err := ioutil.ReadFile(c)
	if err != nil {
		return Config{Wpm: DEFAULT_WPM, HighlightORP: DEFAULT_HIGHLIGHT_ORP}
	}

	config := Config{}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{Wpm: DEFAULT_WPM, HighlightORP: DEFAULT_HIGHLIGHT_ORP}
	}

	return config
}
