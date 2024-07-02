package main

import (
	"log"
	"os"

	"github.com/neptship/ignore/cmd"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Fatal("crashed", "err", err)
		os.Exit(1)
	}
}

func main() {
	defer handlePanic()
	cmd.Execute()
}
