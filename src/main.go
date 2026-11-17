package main

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
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
		case "ctrl+o":
			m.saveAs = true
			m.filenameInput.SetValue("")
			return m, nil
		}

	// handle terminal resizing
	case tea.WindowSizeMsg:
		// this gets the new terminal size and updates our Model struct
		m.width = msg.Width
		m.height = msg.Height
		m.editor.SetWidth(m.width - 14)
		m.editor.SetHeight(m.height - 8)
	}

	var cmd tea.Cmd

	// Checks where keyboard input needs to be sent
	if m.saveAs {
		m.filenameInput, cmd = m.filenameInput.Update(msg)
	} else {
		m.editor, cmd = m.editor.Update(msg)
	}

	if m.editor.Value() != m.savedContent {
		m.ifModified = true
	}

	return m, cmd
}

func main() {

	editor := textarea.New()

	filenameInput := textinput.New()
	filenameInput.Placeholder = "Enter Filename:"

	editor.SetWidth(70)
	editor.ShowLineNumbers = false
	editor.Prompt = ""
	model := Model{
		filename:      "untitled.md",
		editor:        editor,
		filenameInput: filenameInput,
	}
	program := tea.NewProgram(
		model,
	)
	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
