package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [filename]",
	Short: "Create .ignore file",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("./." + args[0] + "ignore")
		if err != nil {
			if os.IsNotExist(err) {
			} else {
				log.Fatal("Error:", err)
			}
		} else {
			log.Fatal("This file already exists")
		}
		_, err = os.Create("." + args[0] + "ignore")
		if err != nil {
			log.Fatal("An error has occurred")
		}

		fmt.Println("." + args[0] + "ignore was created successfully")
	},
}
