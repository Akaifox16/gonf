package branchinput

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type BranchInputModel struct {
	input textinput.Model
}

func InitialModel() BranchInputModel {
	ti := textinput.New()
	ti.Placeholder = "Enter branch name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 20

	return BranchInputModel{
		input: ti,
	}
}

func (m BranchInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m BranchInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m BranchInputModel) View() string {
	return "Enter branch name:\n" + m.input.View()
}

func OpenBranchTextInput() string {
	p := tea.NewProgram(InitialModel())
	model, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	return model.(BranchInputModel).input.Value()
}
