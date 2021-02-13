package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"

	"hinter/model"
)

func main() {
	p := tea.NewProgram(model.InitialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}
}
