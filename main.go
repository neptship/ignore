package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/neptship/ignore/cmd"
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
