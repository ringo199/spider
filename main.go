package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ringo199/spider/ui"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
