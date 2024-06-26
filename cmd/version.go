package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/neptship/ignore/version"
	"github.com/spf13/cobra"
	"html/template"
	"runtime"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the ignore",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := version.Latest()
		if err != nil {
			fmt.Println("Error version check")
		}

		versionInfo := struct {
			Version  string
			OS       string
			Arch     string
			App      string
			Compiler string
		}{
			Version:  version.Version,
			App:      "ignore",
			OS:       runtime.GOOS,
			Arch:     runtime.GOARCH,
			Compiler: runtime.Compiler,
		}

		t, err := template.New("version").Funcs(map[string]any{
			"faint":   lipgloss.NewStyle().Faint(true).Render,
			"bold":    lipgloss.NewStyle().Bold(true).Render,
			"magenta": lipgloss.NewStyle().Foreground(lipgloss.Color("#5a6368")).Render,
		}).Parse(`{{ magenta "▇▇▇" }} {{ magenta .App }} 

  {{ faint "Version" }}  {{ bold .Version }}
  {{ faint "Platform" }} {{ bold .OS }}/{{ bold .Arch }}
  {{ faint "Compiler" }} {{ bold .Compiler }}
`)
		handleErr(err)
		handleErr(t.Execute(cmd.OutOrStdout(), versionInfo))
	},
}
