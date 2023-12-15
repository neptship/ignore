package ignore

import (
	"bytes"
	"slices"
	"testing"
	"testing/fstest"
)

func TestTemplateRegistry_HasTemplate(t *testing.T) {
	mockTemplates := fstest.MapFS{
		"templates/bar.gitignore": &fstest.MapFile{},
		"templates/foo.gitignore": &fstest.MapFile{},
	}

	tr := NewTemplateRegistry()
	tr.templates = mockTemplates

	if !tr.HasTemplate("bar") {
		t.Error("HasTemplate(bar) = false, want = true")
	}

	if tr.HasTemplate("dne") {
		t.Error("HasTemplate(bar) = true, want = false")
	}
}

func TestTemplateRegistry_List(t *testing.T) {
	mockTemplates := fstest.MapFS{
		"templates/bar.gitignore": &fstest.MapFile{},
		"templates/foo.gitignore": &fstest.MapFile{},
	}

	tr := NewTemplateRegistry()
	tr.templates = mockTemplates

	want := []string{"bar", "foo"}
	got := tr.List()

	if !slices.Equal(got, want) {
		t.Errorf("List() = %v, want = %v", got, want)
	}
}

func TestTemplateRegistry_CopyTemplate(t *testing.T) {
	mockTemplates := fstest.MapFS{
		"templates/bar.gitignore": &fstest.MapFile{
			Data: []byte("bar"),
		},
		"templates/foo.gitignore": &fstest.MapFile{
			Data: []byte("foo"),
		},
	}

	tr := NewTemplateRegistry()
	tr.templates = mockTemplates

	var buf bytes.Buffer
	tr.CopyTemplate("foo", &buf)

	want := "foo"
	got := buf.String()

	if got != want {
		t.Errorf("CopyTemplate() - %v, want = %v", got, want)
	}
}
