package cli

import (
	"ptask/cmd/works"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func Select(choices works.Works) map[string]*works.Work {
	model, err := tea.NewProgram(createModelSelect(choices)).Run()
	if err != nil {
		log.Fatal(err)
	}
	w := model.(modelSelect).selected
	return w
}
