package cmd

import (
	"fmt"
	"github.com/neptunsk1y/ignore/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var addCommand = &cobra.Command{
	Use:   "add [filename] [template]",
	Short: "Add a template to .ignore file",
	Run: func(cmd *cobra.Command, args []string) {
		pathFile := "./." + args[0] + "ignore"
		_, err := os.Stat(pathFile)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatal("The file does not exist")
			} else {
				log.Fatal("Error:", err)
			}
		}

		tr := internal.NewTemplateRegistry()
		template := args[1]
		if !tr.HasTemplate(template) {
			log.Fatal("template does not exist")
		}

		file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		err = tr.CopyTemplate(template, file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s template has been added to .%signore\n", template, args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
