package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Create files .ignore quickly and simply",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiBlue + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleErr(err error) {
	if err == nil {
		return
	}

	log.Error(err)
	_, _ = fmt.Fprintf(
		os.Stderr,
		"%s\n",
		strings.Trim(err.Error(), " \n"),
	)
	os.Exit(1)
}
