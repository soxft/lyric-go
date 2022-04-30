package main

import (
	"flag"
	"log"
	"lyric/core"
	"os"
)

var dir string

func main() {
	//use flag to get dir
	var defaultDir string
	if dirNow, err := os.Getwd(); err == nil {
		defaultDir = dirNow
	}
	flag.StringVar(&dir, "dir", defaultDir, "the directory to be processed")
	flag.Parse()

	// no dir specified
	if dir == "" {
		log.Fatal("no dir specified")
	}
	// check if dir exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatal("dir does not exist")
	}

	core.HandlerDir(dir)
}
