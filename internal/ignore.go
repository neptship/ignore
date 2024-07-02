package internal

import (
	"fmt"
	"log"
	"os"
)

func AddIgnoreTemplate(fileName string, templateName string) {
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
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
	fmt.Println(string(colorGreen))
	fmt.Print("√  ", string(colorReset))
	fmt.Println(fileName + " created successfully")
}
