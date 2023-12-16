package cmd

import (
	"fmt"

	"github.com/neptunsk1y/ignore/internal/ignore"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "available templates for .ignore files",
	Run: func(cmd *cobra.Command, args []string) {
		tr := ignore.NewTemplateRegistry()
		templates := tr.List()
		for _, template := range templates {
			fmt.Printf("%s  ", template)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
