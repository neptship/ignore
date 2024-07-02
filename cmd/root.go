package cmd

import (
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/neptship/ignore/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Create files .ignore quickly and simply",
	Run: func(cmd *cobra.Command, args []string) {
		ignoreFile, err := chooseIgnoreFile()
		if err != nil {
			internal.CallClear()
			os.Exit(1)
		}
		ignoreTemplate, err := chooseIgnoreTemplate()
		if err != nil {
			internal.CallClear()
			os.Exit(1)
		}
		internal.AddIgnoreTemplate(ignoreFile, ignoreTemplate)
	},
}

func chooseIgnoreFile() (string, error) {
	tr := internal.NewTemplateRegistry()
	templatesIgnore := tr.List()

	templates := &promptui.SelectTemplates{
		Active:   `{{">" | blue | bold }} {{ . | blue | bold }}`,
		Inactive: `{{.}}`,
		Selected: `{{ "√ " | green | bold }} {{ "Choose .ignore file" | bold }} {{"»" | black}} {{ . | blue }}`,
		Label:    `{{ . | bold }}`,
	}

	prompt := promptui.Select{
		Label:     promptui.IconInitial + " Choose .ignore file",
		Templates: templates,
		Items:     internal.IgnoreFiles,
		Size:      8,
		Searcher: func(input string, index int) bool {
			pepper := templatesIgnore[index]
			name := strings.Replace(strings.ToLower(pepper), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
	}

	_, ignoreFile, err := prompt.Run()
	return ignoreFile, err
}

func chooseIgnoreTemplate() (string, error) {
	tr := internal.NewTemplateRegistry()
	templatesIgnore := tr.List()

	templates := &promptui.SelectTemplates{
		Active:   `{{">" | blue | bold }} {{ . | blue | bold }}`,
		Inactive: `{{ . | underline}}`,
		Selected: `{{ "√ " | green | bold }} {{ "Choose .ignore template" | bold }} {{"»" | black}} {{ . | blue }}`,
		Label:    `{{ . | bold }}`,
	}

	prompt := promptui.Select{
		Label:     promptui.IconInitial + " Choose .ignore template",
		Templates: templates,
		Items:     templatesIgnore,
		Size:      8,
		Searcher: func(input string, index int) bool {
			pepper := templatesIgnore[index]
			name := strings.Replace(strings.ToLower(pepper), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
	}

	_, ignoreTemplate, err := prompt.Run()
	return ignoreTemplate, err
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
