package main

import (
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

// Init() operates any code on startup
func (m Model) Init() tea.Cmd {
	return nil
}

// Update() does opeartions that need to be run continuously, keypresses etc..
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// handle user input
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit
		case "q":
			return m, tea.Quit
		}
	// handle terminal resizing
	case tea.WindowSizeMsg:
		// this gets the new terminal size and updates our Model struct
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func main() {

	editor := textarea.New()
	model := Model{
		filename: "untitled.md",
		editor:   editor,
	}
	program := tea.NewProgram(model)

	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
