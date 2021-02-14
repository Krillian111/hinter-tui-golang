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
	Repo       common.Repository
}

var InitialModel = Model{
	Navigation: components.InitialNavigation(),
	Add:        pages.InitialAdd(),
	Search:     pages.InitialSearch(),
	Repo:       common.Repo(),
}

func (m Model) Init() tea.Cmd {
	addCmds := m.Add.Init()
	return tea.Batch(addCmds)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	navModel, navCmds := m.Navigation.Update(msg)
	m.Navigation = navModel
	cmd = tea.Batch(cmd, navCmds)

	switch m.Navigation.ActiveTab {
	case common.AddTab():
		addModel, addCmds := m.Add.Update(msg, m.Repo)
		m.Add = addModel
		cmd = tea.Batch(cmd, addCmds)
	case common.SearchTab():
		searchModel, searchCmds := m.Search.Update(msg, m.Repo)
		m.Search = searchModel
		cmd = tea.Batch(cmd, searchCmds)
	}

	return m, cmd
}

func (m Model) View() string {

	s := ""
	s += m.Navigation.View()

	switch m.Navigation.ActiveTab {
	case common.AddTab():
		s += m.Add.View()
	case common.SearchTab():
		s += m.Search.View()
	}

	return s
}
