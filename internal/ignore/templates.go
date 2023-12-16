package ignore

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed templates/*
var templates embed.FS

// TemplateRegistry is a collection of all ignore templates available for use.
type TemplateRegistry struct {
	templates fs.FS
}

// NewTemplateRegistry creates a new instance of TemplateRegistry.
func NewTemplateRegistry() *TemplateRegistry {
	return &TemplateRegistry{templates: templates}
}

// HasTemplate indicates if the registry contains a spcfic template.
func (tr *TemplateRegistry) HasTemplate(name string) bool {
	_, err := fs.Stat(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	return err == nil
}

// List returns a list of all templates contained in the registry.
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

// CopyTemplate copys the contents of a template corresponding to the
// provided name to the provided writer.
func (tr *TemplateRegistry) CopyTemplate(name string, dst io.Writer) error {
	b, err := fs.ReadFile(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	if err != nil {
		return err
	}
	io.Copy(dst, bytes.NewReader(b))
	return nil
}
