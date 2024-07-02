package internal

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

var IgnoreFiles = []string{".gitignore", ".bzrignore", ".chefignore", ".cfignore", ".cvsignore", ".boringignore", ".deployignore", ".dockerignore", ".ebignore", ".eleventyignore", ".eslintignore", ".flooignore", ".gcloudignore", ".helmignore", ".jpmignore", ".jshintignore", ".hgignore", ".mtn-ignore", ".nodemonignore", ".npmignore", ".nuxtignore", ".openapi-generator-ignore", ".p4ignore", ".prettierignore", ".stylelintignore", ".stylintignore", ".swagger-codegen-ignore", ".terraformignore", ".tfignore", ".tokeignore", ".upignore", ".vercelignore", ".yarnignore"}

//go:embed templates/*
var templates embed.FS

type TemplateRegistry struct {
	templates fs.FS
}

func NewTemplateRegistry() *TemplateRegistry {
	return &TemplateRegistry{templates: templates}
}

func (tr *TemplateRegistry) HasTemplate(name string) bool {
	_, err := fs.Stat(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	return err == nil
}

func (tr *TemplateRegistry) List() []string {
	var templates []string
	fs.WalkDir(tr.templates, ".", func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}
		template := strings.TrimSuffix(filepath.Base(path), ".gitignore")
		templates = append(templates, template)
		return nil
	})
	return templates
}

func (tr *TemplateRegistry) CopyTemplate(name string, dst io.Writer) error {
	b, err := fs.ReadFile(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	if err != nil {
		return err
	}
	io.Copy(dst, bytes.NewReader(b))
	return nil
}
