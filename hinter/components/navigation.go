package components

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	. "hinter/hinter/common"
)

type NavigationModel struct {
	MenuTabs  []Tab
	ActiveTab Tab
}

var AmountOfTabs = len(AllTabs())

func InitialNavigation() NavigationModel {
	return NavigationModel{
		MenuTabs:  AllTabs(),
		ActiveTab: AddTab(),
	}
}

func (m NavigationModel) Update(msg tea.Msg) (NavigationModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "left":
			m.ActiveTab = FromIndex((int(m.ActiveTab) - 1 + AmountOfTabs) % AmountOfTabs)
		case "right":
			m.ActiveTab = FromIndex((int(m.ActiveTab) + 1 + AmountOfTabs) % AmountOfTabs)
		case "ctrl+c", "ctrl+d":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m NavigationModel) View() string {
	s := ""
	for tab := 0; tab < len(m.MenuTabs); tab++ {
		if tab == int(m.ActiveTab) {
			s += fmt.Sprintf("[%s]", m.MenuTabs[tab])
		} else {
			s += fmt.Sprintf(" %s ", m.MenuTabs[tab])
		}
		s += " "
	}
	return s
}
