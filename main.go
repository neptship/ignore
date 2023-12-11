package main

import (
	"github.com/charmbracelet/log"
	"github.com/neptunsk1y/ignore/cmd"
	"os"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Error("crashed", "err", err)
		os.Exit(1)
	}
}

func main() {
	defer handlePanic()
	cmd.Execute()
}
