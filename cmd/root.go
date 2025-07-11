package cmd

import (
	"os"

	"github.com/neptship/ignore/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Create files .ignore quickly and simply", Run: func(cmd *cobra.Command, args []string) {
		ignoreFile, template, err := chooseBothViaTUI()
		if err != nil {
			internal.CallClear()
			os.Exit(1)
		}

		if ignoreFile == "" || template == "" {
			internal.CallClear()
			return
		}

		internal.AddIgnoreTemplate(ignoreFile, template)
	},
}

func chooseBothViaTUI() (string, string, error) {
	tr := internal.NewTemplateRegistry()
	templates := tr.List()

	return internal.RunBubbleTUI(internal.IgnoreFiles, templates)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
		internal.CallClear()
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
