package main

import (
	"fmt"
	"os"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	target   []rune
	progress []rune
}

func NewModel(target string) Model {

	progress := make([]rune, len(target))
	for i := range progress {
		progress[i] = '_'
	}

	return Model{
		target:   []rune(target),
		progress: progress,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			s := msg.String()
			if len(s) == 1 {
				r := rune(s[0])
				if unicode.IsLetter(r) {
					for i, letter := range m.target {
						if letter == r {
							m.progress[i] = letter
						}
					}
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if string(m.progress) == string(m.target) {
		return "Congrats you won"
	}
	return string(m.progress)
}

func main() {
	p := tea.NewProgram(NewModel("alfie"))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Could not run program: %v\n", err)
		os.Exit(1)
	}
}
