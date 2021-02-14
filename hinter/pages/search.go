package pages

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"hinter/hinter/common"
)

type SearchModel struct {
	query   textinput.Model
	results []common.Entry
}

func InitialSearch() SearchModel {
	query := textinput.NewModel()
	query.Placeholder = "start searching"
	query.Focus()
	query.Prompt = common.FocusedPrompt
	query.TextColor = common.FocusedTextColor
	return SearchModel{query: query}
}

func (m SearchModel) Update(msg tea.Msg, repo common.Repository) (SearchModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.results = repo.Search(m.query.Value())
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.query, cmd = m.query.Update(msg)
	return m, cmd
}

func (m SearchModel) View() string {
	s := "\n"
	s += m.query.View()
	s += "\n"
	for i := 0; i < len(m.results); i++ {
		s += m.results[i].Key
		s += " | "
		s += m.results[i].Value
		s += "\n"
	}
	return s
}
