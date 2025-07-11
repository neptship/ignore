package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func AddIgnoreTemplate(fileName string, templateName string) {
	pathFile := "./" + fileName
	_, err := os.Stat(pathFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(fileName)
			if err != nil {
				log.Fatal("An error has occurred")
			}
		} else {
			log.Fatal("Error:", err)
		}
	}

	tr := NewTemplateRegistry()
	if !tr.HasTemplate(templateName) {
		log.Fatal("template does not exist")
	}

	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = tr.CopyTemplate(templateName, file)
	if err != nil {
		log.Fatal(err)
	}

	successStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00D26A")).
		MarginTop(1).
		MarginBottom(1)

	checkMark := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00D26A"))

	fmt.Println(successStyle.Render("🎉 Success!"))
	fmt.Printf("%s %s created successfully with %s template\n",
		checkMark.Render("✓"),
		fileName,
		templateName)
}
