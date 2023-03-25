package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Wpm int
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
		return Config{Wpm: 200}
	}

	config := Config{}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{Wpm: 200}
	}

	fmt.Println("a", config)

	return config
}
