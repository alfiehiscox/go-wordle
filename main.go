package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	target   []rune
	progress [][]rune

	index   int
	guess   int
	guesses [][]rune
}

func NewModel(target string) Model {

	progress := make([][]rune, len(target))
	for i := range progress {
		p := make([]rune, len(target))
		for j := range p {
			p[j] = '_'
		}
		progress[i] = p
	}

	guesses := make([][]rune, 5)
	for i := range guesses {
		guesses[i] = make([]rune, len(target))
	}

	return Model{
		target:   []rune(target),
		progress: progress,

		index:   0,
		guess:   0,
		guesses: guesses,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.guess > 0 {
				copy(m.progress[m.guess], m.progress[m.guess-1])
			}

			for i, letter := range m.guesses[m.guess] {
				if letter == m.target[i] {
					m.progress[m.guess][i] = letter
				}
			}
			m.index = 0
			m.guess++
		case "back":
			if m.index > 0 {
				m.guesses[m.guess][m.index] = '_'
				m.index -= 1
			}
		default:
			s := msg.String()
			if len(s) == 1 {
				r := rune(s[0])
				if unicode.IsLetter(r) {
					m.guesses[m.guess][m.index] = r
					m.index += 1
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {

	b := strings.Builder{}
	for i, guess := range m.guesses {
		progress := m.progress[i]
		if string(progress) == string(m.target) {
			return "Congrats you won"
		}
		b.WriteString(string(guess) + "->" + string(progress) + "\n")
	}
	return b.String()
}

func main() {
	p := tea.NewProgram(NewModel("alfie"))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Could not run program: %v\n", err)
		os.Exit(1)
	}
}
