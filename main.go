package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/neptunsk1y/ignore/cmd"
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
