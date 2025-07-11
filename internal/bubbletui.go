package internal

import (
	"fmt"
	"strings"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var stringSlicePool = sync.Pool{
	New: func() interface{} {
		s := make([]string, 0, 100)
		return &s
	},
}

type TUIState int

const (
	StateIgnoreFile TUIState = iota
	StateTemplate
	StatePreview
	StateSearch
	StateDone
)

var (
	cursorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Bold(true)
	titleStyle       = lipgloss.NewStyle().Bold(true).Underline(true)
	helpStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Italic(true)
	separatorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444"))
	searchStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D26A")).Bold(true)
	previewStyle     = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#7D56F4")).Padding(1)
	previewTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#E0E0E0"))
	successStyle     = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#00D26A"))
)

type BubbleTUIModel struct {
	State             TUIState
	IgnoreFiles       []string
	Templates         []string
	FilteredFiles     []string
	FilteredTemplates []string
	IgnoreCursor      int
	TemplateCursor    int
	ScrollOffset      int
	MaxVisible        int
	SelectedFile      string
	SelectedTemplate  string
	TemplateContent   string
	PreviewScroll     int
	SearchQuery       string
	SearchMode        bool
	SearchContext     string
	Err               error
	Quitting          bool
	templateRegistry  *TemplateRegistry
}

func NewBubbleTUI(ignoreFiles []string, templates []string) *BubbleTUIModel {
	return &BubbleTUIModel{
		State:             StateIgnoreFile,
		IgnoreFiles:       ignoreFiles,
		Templates:         templates,
		FilteredFiles:     ignoreFiles,
		FilteredTemplates: templates,
		MaxVisible:        10,
		templateRegistry:  NewTemplateRegistry(),
	}
}

func (m *BubbleTUIModel) filterItems() {
	if m.SearchQuery == "" {
		m.FilteredFiles = m.IgnoreFiles
		m.FilteredTemplates = m.Templates
		return
	}

	query := strings.ToLower(m.SearchQuery)

	filteredFiles := stringSlicePool.Get().(*[]string)
	filteredTemplates := stringSlicePool.Get().(*[]string)
	defer func() {
		stringSlicePool.Put(filteredFiles)
		stringSlicePool.Put(filteredTemplates)
	}()

	*filteredFiles = (*filteredFiles)[:0]
	for _, file := range m.IgnoreFiles {
		if strings.Contains(strings.ToLower(file), query) {
			*filteredFiles = append(*filteredFiles, file)
		}
	}

	*filteredTemplates = (*filteredTemplates)[:0]
	for _, template := range m.Templates {
		if strings.Contains(strings.ToLower(template), query) {
			*filteredTemplates = append(*filteredTemplates, template)
		}
	}

	m.FilteredFiles = make([]string, len(*filteredFiles))
	copy(m.FilteredFiles, *filteredFiles)

	m.FilteredTemplates = make([]string, len(*filteredTemplates))
	copy(m.FilteredTemplates, *filteredTemplates)
}

func (m *BubbleTUIModel) resetSearch() {
	m.SearchQuery = ""
	m.SearchMode = false
	m.filterItems()
	m.ScrollOffset = 0
	m.IgnoreCursor = 0
	m.TemplateCursor = 0
}

func (m *BubbleTUIModel) Init() tea.Cmd {
	return nil
}

func (m *BubbleTUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.State {
	case StateIgnoreFile:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if m.SearchMode {
				switch msg.String() {
				case "enter":
					m.SearchMode = false
					m.ScrollOffset = 0
					m.IgnoreCursor = 0
				case "esc":
					m.resetSearch()
				case "backspace":
					if len(m.SearchQuery) > 0 {
						m.SearchQuery = m.SearchQuery[:len(m.SearchQuery)-1]
						m.filterItems()
					} else {
						m.resetSearch()
					}
				default:
					if len(msg.String()) == 1 {
						m.SearchQuery += msg.String()
						m.filterItems()
						m.ScrollOffset = 0
						m.IgnoreCursor = 0
					}
				}
			} else {
				switch msg.String() {
				case "ctrl+c", "q", "й":
					m.Quitting = true
					return m, tea.Quit
				case "/":
					m.SearchMode = true
					m.SearchContext = "ignore"
					m.SearchQuery = ""
				case "up", "k":
					if m.IgnoreCursor > 0 {
						m.IgnoreCursor--
						if m.IgnoreCursor < m.ScrollOffset {
							m.ScrollOffset = m.IgnoreCursor
						}
					}
				case "down", "j":
					if m.IgnoreCursor < len(m.FilteredFiles)-1 {
						m.IgnoreCursor++
						if m.IgnoreCursor >= m.ScrollOffset+m.MaxVisible {
							m.ScrollOffset = m.IgnoreCursor - m.MaxVisible + 1
						}
					}
				case "enter":
					if m.IgnoreCursor < len(m.FilteredFiles) {
						m.SelectedFile = m.FilteredFiles[m.IgnoreCursor]
						m.State = StateTemplate
						m.TemplateCursor = 0
						m.ScrollOffset = 0
						m.resetSearch()
					}
				}
			}
		}
	case StateTemplate:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if m.SearchMode {
				switch msg.String() {
				case "enter":
					m.SearchMode = false
					m.ScrollOffset = 0
					m.TemplateCursor = 0
				case "esc":
					m.resetSearch()
				case "backspace":
					if len(m.SearchQuery) > 0 {
						m.SearchQuery = m.SearchQuery[:len(m.SearchQuery)-1]
						m.filterItems()
					} else {
						m.resetSearch()
					}
				default:
					if len(msg.String()) == 1 {
						m.SearchQuery += msg.String()
						m.filterItems()
						m.ScrollOffset = 0
						m.TemplateCursor = 0
					}
				}
			} else {
				switch msg.String() {
				case "ctrl+c", "q", "й":
					m.Quitting = true
					return m, tea.Quit
				case "esc", "backspace":
					m.State = StateIgnoreFile
					m.ScrollOffset = 0
					m.resetSearch()
					return m, nil
				case "/":
					m.SearchMode = true
					m.SearchContext = "template"
					m.SearchQuery = ""
				case "up", "k":
					if m.TemplateCursor > 0 {
						m.TemplateCursor--
						if m.TemplateCursor < m.ScrollOffset {
							m.ScrollOffset = m.TemplateCursor
						}
					}
				case "down", "j":
					if m.TemplateCursor < len(m.FilteredTemplates)-1 {
						m.TemplateCursor++
						if m.TemplateCursor >= m.ScrollOffset+m.MaxVisible {
							m.ScrollOffset = m.TemplateCursor - m.MaxVisible + 1
						}
					}
				case "enter":
					if m.TemplateCursor < len(m.FilteredTemplates) {
						m.State = StateDone
						return m, tea.Quit
					}
				case "p":
					if m.TemplateCursor < len(m.FilteredTemplates) {
						m.SelectedTemplate = m.FilteredTemplates[m.TemplateCursor]
						content, err := m.templateRegistry.GetTemplateContent(m.SelectedTemplate)
						if err != nil {
							m.Err = err
							return m, nil
						}
						m.TemplateContent = content
						m.State = StatePreview
						m.PreviewScroll = 0
					}
				}
			}
		}
	case StatePreview:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q", "й":
				m.Quitting = true
				return m, tea.Quit
			case "esc", "backspace":
				m.State = StateTemplate
				return m, nil
			case "up", "k":
				if m.PreviewScroll > 0 {
					m.PreviewScroll--
				}
			case "down", "j":
				lines := strings.Split(m.TemplateContent, "\n")
				maxScroll := len(lines) - 15
				if maxScroll < 0 {
					maxScroll = 0
				}
				if m.PreviewScroll < maxScroll {
					m.PreviewScroll++
				}
			case "enter":
				m.State = StateDone
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m *BubbleTUIModel) View() string {
	switch m.State {
	case StateIgnoreFile:
		return m.viewIgnoreFile()
	case StateTemplate:
		return m.viewTemplate()
	case StatePreview:
		return m.viewPreview()
	case StateDone:
		if m.Err != nil {
			return fmt.Sprintf("Error: %v\n", m.Err)
		}
		return ""
	default:
		return ""
	}
}

func (m *BubbleTUIModel) viewIgnoreFile() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("Choose ignore file:") + "\n")

	if m.SearchMode {
		searchPrompt := fmt.Sprintf("Search: %s", m.SearchQuery)
		b.WriteString(searchStyle.Render(searchPrompt) + "\n")
	}
	b.WriteString("\n")

	start := m.ScrollOffset
	end := start + m.MaxVisible
	if end > len(m.FilteredFiles) {
		end = len(m.FilteredFiles)
	}

	for i := start; i < end; i++ {
		cursor := " "
		if m.IgnoreCursor == i {
			cursor = cursorStyle.Render(">")
		}
		line := fmt.Sprintf(" %s %s", cursor, m.FilteredFiles[i])
		b.WriteString(line + "\n")
	}

	if m.SearchQuery != "" {
		b.WriteString(fmt.Sprintf("\n%s %d results", helpStyle.Render("Found:"), len(m.FilteredFiles)))
	}

	var helpText string
	if m.SearchMode {
		helpText = helpStyle.Render(" Type to search   Enter: Finish search   Esc: Cancel   Backspace: Delete ")
	} else {
		helpText = helpStyle.Render(" ↑/↓: Navigate   Enter: Select   /: Search   q: Quit ")
	}
	separator := separatorStyle.Render(strings.Repeat("─", lipgloss.Width(helpText)))
	b.WriteString("\n" + separator + "\n" + helpText + "\n")
	return b.String()
}

func (m *BubbleTUIModel) viewTemplate() string {
	var b strings.Builder
	b.WriteString(successStyle.Render("✓ ") + titleStyle.Render("Selected: ") + titleStyle.Render(m.SelectedFile) + "\n")
	b.WriteString(titleStyle.Render("Choose template:") + "\n")

	if m.SearchMode {
		searchPrompt := fmt.Sprintf("Search: %s", m.SearchQuery)
		b.WriteString(searchStyle.Render(searchPrompt) + "\n")
	}
	b.WriteString("\n")

	start := m.ScrollOffset
	end := start + m.MaxVisible
	if end > len(m.FilteredTemplates) {
		end = len(m.FilteredTemplates)
	}

	for i := start; i < end; i++ {
		cursor := " "
		if m.TemplateCursor == i {
			cursor = cursorStyle.Render(">")
		}
		line := fmt.Sprintf(" %s %s", cursor, m.FilteredTemplates[i])
		b.WriteString(line + "\n")
	}

	if m.SearchQuery != "" {
		b.WriteString(fmt.Sprintf("\n%s %d results", helpStyle.Render("Found:"), len(m.FilteredTemplates)))
	}

	var helpText string
	if m.SearchMode {
		helpText = helpStyle.Render(" Type to search   Enter: Finish search   Esc: Cancel   Backspace: Delete ")
	} else {
		helpText = helpStyle.Render(" ↑/↓: Navigate   Enter: Select   /: Search   p: Preview   esc/backspace: Back   q: Quit ")
	}
	separator := separatorStyle.Render(strings.Repeat("─", lipgloss.Width(helpText)))
	b.WriteString("\n" + separator + "\n" + helpText + "\n")
	return b.String()
}

func (m *BubbleTUIModel) viewPreview() string {
	var b strings.Builder

	b.WriteString(successStyle.Render("✓ ") + titleStyle.Render("Selected: ") + titleStyle.Render(m.SelectedFile) + "\n")
	b.WriteString(titleStyle.Render("Preview: ") + titleStyle.Render(m.SelectedTemplate) + "\n\n")

	normalizedContent := strings.ReplaceAll(m.TemplateContent, "\r\n", "\n")
	normalizedContent = strings.ReplaceAll(normalizedContent, "\r", "\n")
	lines := strings.Split(normalizedContent, "\n")

	totalLines := len(lines)
	if totalLines == 0 {
		lines = []string{""}
		totalLines = 1
	}
	start := m.PreviewScroll
	if start >= totalLines {
		start = 0
	}
	end := start + 15
	if end > totalLines {
		end = totalLines
	}

	var previewLines []string
	maxWidth := 0
	for i := start; i < end; i++ {
		if len(lines[i]) > maxWidth {
			maxWidth = len(lines[i])
		}
	}
	minWidth := 40
	if maxWidth < minWidth {
		maxWidth = minWidth
	}
	for i := start; i < end; i++ {
		line := lines[i]
		if len(line) < maxWidth {
			line += strings.Repeat(" ", maxWidth-len(line))
		}
		previewLines = append(previewLines, line)
	}
	for len(previewLines) < 15 {
		previewLines = append(previewLines, strings.Repeat(" ", maxWidth))
	}
	previewContent := strings.Join(previewLines, "\n")

	styledContent := previewStyle.Render(previewTextStyle.Render(previewContent))
	b.WriteString(styledContent + "\n\n")

	if totalLines > 15 {
		scrollInfo := fmt.Sprintf("Lines %d-%d of %d", start+1, end, totalLines)
		b.WriteString(helpStyle.Render(scrollInfo) + "\n")
	}

	helpText := helpStyle.Render(" ↑/↓: Scroll   Enter: Select   esc/backspace: Back   q: Quit ")
	separator := separatorStyle.Render(strings.Repeat("─", lipgloss.Width(helpText)))
	b.WriteString(separator + "\n" + helpText + "\n")

	return b.String()
}

func RunBubbleTUI(ignoreFiles []string, templates []string) (string, string, error) {
	model := NewBubbleTUI(ignoreFiles, templates)
	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return "", "", err
	}

	m := finalModel.(*BubbleTUIModel)
	if m.Quitting {
		return "", "", nil
	}

	selectedTemplate := ""
	if m.TemplateCursor < len(m.FilteredTemplates) {
		selectedTemplate = m.FilteredTemplates[m.TemplateCursor]
	}

	return m.SelectedFile, selectedTemplate, nil
}
