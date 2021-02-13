package model

import (
	"fmt"
)

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

	return s
}
