package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Tab int

const (
	Add = iota
	Remove
)

const amountOfTabs = 2

func (t Tab) String() string {
	return [...]string{"Add", "Remove"}[t]
}

type Model struct {
	menuTabs  []Tab
	activeTab Tab
}

var InitialModel = Model{
	menuTabs:  []Tab{Add, Remove},
	activeTab: Add,
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "left":
			m.activeTab = (m.activeTab - 1 + amountOfTabs) % amountOfTabs
		case "right":
			m.activeTab = (m.activeTab + 1 + amountOfTabs) % amountOfTabs
		case "ctrl+c", "ctrl+d":
			return m, tea.Quit
		}

	}

	return m, nil
}
