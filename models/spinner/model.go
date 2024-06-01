package spinner

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type spinnerModel struct {
	spinner spinner.Model
	task    *exec.Cmd
	err     error
	done    bool
}

func NewSpinnerModel(task *exec.Cmd) spinnerModel {

	s := spinner.New()
	s.Spinner = spinner.Line

	return spinnerModel{spinner: s, task: task}
}

func (m spinnerModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, runTask(m.task))
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		if m.done {
			return m, tea.Quit
		}

		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case taskDoneMsg:
		m.done = true
		m.err = msg.err
		return m, tea.Quit
	}

	return m, nil
}

func (m spinnerModel) View() string {
	if m.done {
		if m.err != nil {
			return fmt.Sprintf("Error: %v\n", m.err)
		}

		return "Task completed successfully!\n"
	}

	return fmt.Sprintf("%s %s\n", m.spinner.View(), m.task)
}

type taskDoneMsg struct {
	err error
}

func runTask(task *exec.Cmd) tea.Cmd {
	return func() tea.Msg {
		err := task.Run()
		time.Sleep(2 * time.Second)
		return taskDoneMsg{err: err}
	}
}

func RunSpinner(task *exec.Cmd) error {
	p := tea.NewProgram(NewSpinnerModel(task))
	_, err := p.Run()

	return err
}
