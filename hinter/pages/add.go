package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"
	"hinter/hinter/common"

	"github.com/charmbracelet/bubbles/textinput"
)

type AddModel struct {
	key         textinput.Model
	value       textinput.Model
	activeInput int
}

const amountOfInputs = 2

const focusedTextColor = "205"

var (
	color         = te.ColorProfile().Color
	focusedPrompt = te.String("> ").Foreground(color("205")).String()
	blurredPrompt = "> "
)

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

func (m AddModel) Update(msg tea.Msg, entries []common.Entry) (AddModel, []common.Entry, tea.Cmd) {
	var cmd tea.Cmd

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
					inputs[i].Prompt = focusedPrompt
					inputs[i].TextColor = focusedTextColor
					continue
				}
				// Remove focused state
				inputs[i].Blur()
				inputs[i].Prompt = blurredPrompt
				inputs[i].TextColor = ""
			}

			m.key = inputs[0]
			m.value = inputs[1]

			return m, entries, nil
		case "enter":
			key := m.key.Value()
			value := m.value.Value()
			entries = append(entries, common.Entry{Key: key, Value: value})
			m.key.Reset()
			m.value.Reset()
			return m, entries, nil
		}
	}

	// Handle character input and blinks
	m, cmd = updateInputs(msg, m)
	return m, entries, cmd
}

// Pass messages and models through to text input components. Only text inputs
// with Focus() set will respond, so it's safe to simply update all of them
// here without any further logic.
func updateInputs(msg tea.Msg, m AddModel) (AddModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.key, cmd = m.key.Update(msg)
	cmds = append(cmds, cmd)

	m.value, cmd = m.value.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
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
