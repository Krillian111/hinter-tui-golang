package pages

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"hinter/hinter/common"
	"hinter/hinter/views"
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
	s += views.FromEntries([]int{20, 60}, m.results).View()
	return s
}
