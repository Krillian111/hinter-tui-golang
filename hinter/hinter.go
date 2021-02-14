package hinter

import (
	tea "github.com/charmbracelet/bubbletea"
	"hinter/hinter/common"
	"hinter/hinter/components"
	"hinter/hinter/pages"
)

type Model struct {
	Navigation components.NavigationModel
	Add        pages.AddModel
	Search     pages.SearchModel
	Entries    []common.Entry
}

var InitialModel = Model{
	Navigation: components.InitialNavigation(),
	Add:        pages.InitialAdd(),
	Search:     pages.InitialSearch(),
	Entries:    []common.Entry{},
}

func (m Model) Init() tea.Cmd {
	addCmds := m.Add.Init()
	return tea.Batch(addCmds)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	navModel, navCmds := m.Navigation.Update(msg)
	m.Navigation = navModel

	addModel, entries, addCmds := m.Add.Update(msg, m.Entries)
	m.Add = addModel
	m.Entries = entries

	cmds := tea.Batch(navCmds, addCmds)
	return m, cmds
}

func (m Model) View() string {

	s := ""
	s += m.Navigation.View()

	switch m.Navigation.ActiveTab {
	case common.AddTab():
		s += m.Add.View()
	case common.SearchTab():
		s += m.Search.View(m.Entries)
	}

	return s
}
