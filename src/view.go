package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/muesli/reflow/wordwrap"
)

// here we declare our lipgloss styles
var (
	headerStyle = lipgloss.NewStyle().Bold(true)
	docStyle    = lipgloss.NewStyle().Padding(2, 4).Align(lipgloss.Center)
	footerStyle = lipgloss.NewStyle().Faint(true)
)

// View() handles text dispalyed in the term
func (m Model) View() tea.View {
	appStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Padding(1, 2).
		Width(m.width - 2).
		Height(m.height - 2)

	header := headerStyle.
		Width(m.width).
		Align(lipgloss.Center)

	docWidth := 80
	if m.width < docWidth {
		docWidth = m.width
	}
	wrappedContent := wordwrap.String(m.editor.View(), docWidth-10)

	document := docStyle.Width(docWidth).Render(wrappedContent)
	document = lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		document,
	)
	content := strings.Join([]string{
		header.Render("DraftTerm | untitled.md"),
		document,
		footerStyle.Render("Ctrl+S Save | Ctrl+Q Quit"),
	}, "\n")

	return tea.NewView(appStyle.Render(content))
}
