package hinter

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"hinter/hinter/common"
	"hinter/hinter/pages"
)

type Tab int

const (
	Add = iota
	Search
)

const amountOfTabs = 2

func (t Tab) String() string {
	return [...]string{"Add", "Search"}[t]
}

type Model struct {
	menuTabs  []Tab
	activeTab Tab
	Entries   []common.Entry
	Add       pages.AddModel
	Search    pages.SearchModel
}

var InitialModel = Model{
	menuTabs:  []Tab{Add, Search},
	activeTab: Add,
	Entries:   []common.Entry{},
	Add:       pages.InitialModel(),
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
	cmds, add, entries := m.Add.Update(msg, m.Entries)
	m.Add = add
	m.Entries = entries

	return m, cmds
}

func (m Model) View() string {

	s := ""
	for tab := 0; tab < len(m.menuTabs); tab++ {
		if tab == int(m.activeTab) {
			s += fmt.Sprintf("[%s]", m.menuTabs[tab])
		} else {
			s += fmt.Sprintf(" %s ", m.menuTabs[tab])
		}
		s += " "
	}
	if m.activeTab == Add {
		s += m.Add.View()
	}
	if m.activeTab == Search {
		s += m.Search.View(m.Entries)
	}

	return s
}
