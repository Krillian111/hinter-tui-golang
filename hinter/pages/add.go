package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"hinter/hinter/common"

	"github.com/charmbracelet/bubbles/textinput"
)

type AddModel struct {
	key         textinput.Model
	value       textinput.Model
	activeInput int
}

const amountOfInputs = 2

func InitialAdd() AddModel {
	key := textinput.NewModel()
	key.Placeholder = "key"
	value := textinput.NewModel()
	value.Placeholder = "value"
	return AddModel{key, value, 0}
}

func (m AddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddModel) Update(msg tea.Msg, repo common.Repository) (AddModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "tab", "up", "down":

			if msg.String() == "tab" || msg.String() == "down" {
				m.activeInput = (m.activeInput + 1 + amountOfInputs) % amountOfInputs
			}

			if msg.String() == "up" {
				m.activeInput = (m.activeInput - 1 + amountOfInputs) % amountOfInputs
			}

			inputs := []textinput.Model{
				m.key,
				m.value,
			}

			for i := 0; i <= len(inputs)-1; i++ {
				if i == m.activeInput {
					// Set focused state
					inputs[i].Focus()
					inputs[i].Prompt = common.FocusedPrompt
					inputs[i].TextColor = common.FocusedTextColor
					continue
				}
				// Remove focused state
				inputs[i].Blur()
				inputs[i].Prompt = common.BlurredPrompt
				inputs[i].TextColor = ""
			}

			m.key = inputs[0]
			m.value = inputs[1]

			return m, nil
		case "enter":
			key := m.key.Value()
			value := m.value.Value()
			repo.Add(common.Entry{Key: key, Value: value})
			m.key.Reset()
			m.value.Reset()
			return m, nil
		}
	}

	var keyCmds, valueCmds tea.Cmd
	m.key, keyCmds = m.key.Update(msg)
	m.value, valueCmds = m.value.Update(msg)
	return m, tea.Batch(keyCmds, valueCmds)
}

func (m AddModel) View() string {
	s := "\n"

	inputs := []string{
		m.key.View(),
		m.value.View(),
	}

	for i := 0; i < len(inputs); i++ {
		s += inputs[i]
		if i < len(inputs)-1 {
			s += "\n"
		}
	}
	return s
}
