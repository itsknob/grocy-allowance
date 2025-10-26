package tui

import (
	"fmt"
	"sort"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modelModalUi struct {
	modals        map[string]tea.Model
	selectedModal string
	curosr        int
}

func (m *modelModalUi) Init() tea.Cmd {
	return nil
}

func (m *modelModalUi) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case tea.KeyCtrlOpenBracket.String():
			m.switchPreviousModal()
			return m, nil
		case tea.KeyCtrlCloseBracket.String():
			m.switchNextModal()
			return m, nil
		}
	}

	_, cmd := m.Update(msg)

	return m, tea.Batch(doTick(), cmd)
}

func (m *modelModalUi) View() string {
	current, exists := m.modals[m.selectedModal]
	if !exists {
		return (&defaultModel{}).View()
	}

	styleModal := lipgloss.NewStyle().
		BorderBottom(true).
		BorderTop(true).
		BorderStyle(lipgloss.RoundedBorder())

	view := m.headerLine() + "\r\n"
	view += styleModal.Render(current.View()) + "\r\n"
	view += m.footerLine()

	return view
}

func (m *modelModalUi) headerLine() string {
	tabs := make([]string, 0)

	for _, name := range m.GetModalNames() {
		if name == m.selectedModal {
			style := lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#7D56F4"))
			name = style.Render(fmt.Sprintf(" %s ", name))
		}

		tabs = append(tabs, lipgloss.NewStyle().Padding(0, 1).Render(name))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (m *modelModalUi) footerLine() string {
	var footer string
	hotkeys := []string{
		"Previous::ctrl+[",
		"Next::ctrl+]",
	}

	for _, entry := range hotkeys {
		footer += entry
	}
	return footer
}

func (m *modelModalUi) CurrentModalName() string {
	return m.selectedModal
}

func (m *modelModalUi) GetModalNames() (keys []string) {
	return m.getSortedModalNameList()
}

func (m *modelModalUi) getSortedModalNameList() (keys []string) {
	for key := range m.modals {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return
}

func (m *modelModalUi) SwitchModal(name string) {
	if _, exists := m.modals[name]; exists {
		m.selectedModal = name
	}
}

func (m *modelModalUi) switchPreviousModal() {
	m.SwitchModal(m.previous(m.CurrentModalName()))
}

func (m *modelModalUi) previous(current string) string {
	keys := m.GetModalNames()

	for i, item := range keys {
		// Don't process current item in loop
		if item != current {
			continue
		}

		// Wrap if already at first
		if i == 0 {
			return keys[len(keys)-1]
		}

		// Switch to Previous
		return keys[i-1]
	}

	// return first if current not found in keys
	return keys[0]
}

func (m *modelModalUi) next(current string) string {
	keys := m.GetModalNames()

	for i, item := range keys {
		if item == current {
			if i == len(keys)-1 {
				// wrap if current is this item and this item is last in list
				return keys[0]
			}
			// Next item
			return keys[i+1]
		}
	}
	// default to first in list if current not found in lsit
	return keys[0]
}

func (m *modelModalUi) switchNextModal() {
	m.SwitchModal(m.next(m.CurrentModalName()))
}

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second/4, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
