package cmd

import (
	"fmt"
	"github.com/neptship/ignore/internal"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Available templates for .ignore files",
	Run: func(cmd *cobra.Command, args []string) {
		tr := internal.NewTemplateRegistry()
		templates := tr.List()
		for _, template := range templates {
			fmt.Printf("%s  ", template)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
