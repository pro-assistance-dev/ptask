package cli

import (
	"fmt"
	"ptask/cmd/works"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type modelSelect struct {
	focused  int
	err      error
	choices  works.Works
	selected map[string]*works.Work
	cursor   int
}

func (m modelSelect) Init() tea.Cmd {
	return textinput.Blink
}

func createModelSelect(choices works.Works) modelSelect {
	return modelSelect{
		focused:  0,
		err:      nil,
		choices:  choices,
		selected: make(map[string]*works.Work),
	}
}

func (m modelSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k", "л":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j", "о":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.choices[m.cursor].GetName()]
			if ok {
				delete(m.selected, m.choices[m.cursor].GetName())
			} else {
				m.selected[m.choices[m.cursor].GetName()] = &m.choices[m.cursor]
			}
		}
	}
	return m, nil
}

func (m modelSelect) View() string {
	s := strings.Builder{}
	s.WriteString("Select work\n\n")

	for i := 0; i < len(m.choices); i++ {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[m.choices[i].GetName()]; ok {
			checked = "x"
		}

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, m.choices[i].GetName()))
	}

	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
