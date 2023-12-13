package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "available templates for .ignore files",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("templates/")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Printf("%s  ", file.Name()[:len(file.Name())-10])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
