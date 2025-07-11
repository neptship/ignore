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
	templates      fs.FS
	templatesList  []string
	templatesCache map[string]string
	maxCacheSize   int
	cacheKeys      []string
}

func NewTemplateRegistry() *TemplateRegistry {
	return &TemplateRegistry{
		templates:      templates,
		templatesCache: make(map[string]string),
		maxCacheSize:   100,
	}
}

func (tr *TemplateRegistry) HasTemplate(name string) bool {
	_, err := fs.Stat(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	return err == nil
}

func (tr *TemplateRegistry) List() []string {
	if tr.templatesList != nil {
		return tr.templatesList
	}

	tr.templatesList = make([]string, 0, 600)
	fs.WalkDir(tr.templates, ".", func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}
		template := strings.TrimSuffix(filepath.Base(path), ".gitignore")
		tr.templatesList = append(tr.templatesList, template)
		return nil
	})
	return tr.templatesList
}

func (tr *TemplateRegistry) CopyTemplate(name string, dst io.Writer) error {
	b, err := fs.ReadFile(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	if err != nil {
		return err
	}
	io.Copy(dst, bytes.NewReader(b))
	return nil
}

func (tr *TemplateRegistry) GetTemplateContent(name string) (string, error) {
	if content, exists := tr.templatesCache[name]; exists {
		return content, nil
	}

	b, err := fs.ReadFile(tr.templates, fmt.Sprintf("templates/%s.gitignore", name))
	if err != nil {
		return "", err
	}

	content := string(b)
	tr.templatesCache[name] = content

	if len(tr.templatesCache) > tr.maxCacheSize {
		oldestKey := tr.cacheKeys[0]
		tr.cacheKeys = tr.cacheKeys[1:]
		delete(tr.templatesCache, oldestKey)
	}

	tr.cacheKeys = append(tr.cacheKeys, name)

	return content, nil
}
