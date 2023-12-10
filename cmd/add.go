package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "add a template to .ignore file",
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
		count := 1
		if _, err = os.Stat("templates/other/" + args[1] + ".txt"); err != nil {
			if os.IsNotExist(err) {
				if _, err = os.Stat("templates/languages/" + args[1] + ".txt"); err != nil {
					if os.IsNotExist(err) {
						log.Fatal("The template does not exist")
					}

				}
				count += 1
			} else {
				log.Fatal("Error:", err)
			}
		}
		var pathTemplateFile string
		if count == 1 {
			pathTemplateFile = "templates/other/" + args[1] + ".txt"
		} else {
			pathTemplateFile = "templates/languages/" + args[1] + ".txt"
		}
		templateFile, err := os.ReadFile(pathTemplateFile)
		if err != nil {
			log.Fatal("Error reading the file")
		}
		file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if _, err := file.WriteString("\n" + string(templateFile)); err != nil {
			log.Fatal(err)
		}
		fmt.Println(args[1] + " template has been added to ." + args[0] + "ignore")
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
