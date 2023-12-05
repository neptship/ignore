package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create .ignore file",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("./." + args[0] + "ignore")
		if err != nil {
			if os.IsNotExist(err) {
			} else {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("This file already exists")
			os.Exit(1)
		}
		_, err = os.Create("." + args[0] + "ignore")
		if err != nil {
			fmt.Println("An error has occurred")
			os.Exit(1)
		}

		fmt.Println("The file was created successfully")
	},
}
