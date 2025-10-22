package tui

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"example.com/grocy-allowance/grocy"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/**
* @See https://github.com/charmbracelet/bubbletea/tree/main/tutorials/basics
*
* Init - a function that returns initial command to run app
* Update - a function that handles incoming events and updates the model
* View - a function that renders the UI based on data in model
 */

const (
	PAGE_HOME      = "Home"
	PAGE_DEPOSIT   = "Deposit"
	PAGE_WITHDRAWL = "Withdrawl"
	PAGE_BALANCE   = "Balance"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

// Application State
type allowanceModel struct {
	menuOptions        []string          // Menu Options
	pageOptions        []string          // Page Options
	focusIndex         int               // Current Position
	selectedMenuOption string            // Current Page
	selectedPageOption string            // Current Input
	pageInputs         []textinput.Model // Text Inputs on Page
	grocyClient        grocy.GrocyClient // Client
}

// Initialize the Application State, Return Model
func InitialModel() allowanceModel {
	m := allowanceModel{
		pageInputs: make([]textinput.Model, 3),
	}
	var t textinput.Model
	for i := range m.pageInputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		// User
		case 0:
			t.Prompt = "User:"
			t.Width = 32
			t.Placeholder = "Dad"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			// Amount
		case 1:
			t.Prompt = "Amount:"
			t.Width = 32
			t.Placeholder = "0.00"
			t.CharLimit = 7 // 9999.99
			// Date
		case 2:
			t.Prompt = "Date:"
			t.Width = 32
			t.Placeholder = time.Now().Format("YYYY-MM-DD")
		}

		m.pageInputs[i] = t

	}

	config := grocy.GrocyConfig{
		GROCY_URL: os.Getenv("GROCY_URL"),
	}
	if config.GROCY_URL == "" {
		log.Fatal("GROCY_URL env variable is not set.")
	}

	grocyClient := grocy.NewGrocyClient(config.GROCY_URL)

	// Ensure Allowance Exists
	if !grocyClient.HasAllowance() {
		log.Default().Println("Allowance not found. Setting up new userentity")
		grocyClient.InitAllowance()
	}

	// Initial Menu Options
	m.menuOptions = []string{PAGE_HOME, PAGE_DEPOSIT, PAGE_WITHDRAWL, PAGE_BALANCE}
	m.selectedMenuOption = PAGE_HOME
	m.grocyClient = *grocyClient

	return m
}

func (m allowanceModel) Init() tea.Cmd {
	if m.selectedMenuOption == PAGE_HOME || m.selectedMenuOption == PAGE_BALANCE {
		return cursor.Blink
	} else {
		return textinput.Blink
	}
}

func (m allowanceModel) makeHomeMenu() []tea.Cmd {
	// Reset State
	m.selectedMenuOption = PAGE_HOME

	// Return new Inputs for Page
	return make([]tea.Cmd, len(m.pageInputs))
}

/**
* Called when things happen
* It's a Reducer
 */
func (m allowanceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Switch on Type of Message Received
	switch msg := msg.(type) {
	// is it a kepress?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			// Did user press enter while submit button was focused? Exists
			if s == "enter" && m.focusIndex == len(m.pageInputs) && m.selectedMenuOption != PAGE_HOME {
				log.Default().Println("Submitted")

				return m, nil
			}

			// 0. Home
			// 1. Deposit
			if s == "enter" && m.focusIndex == 1 {
				newState := allowanceModel{
					selectedMenuOption: PAGE_DEPOSIT,
					pageInputs:         make([]textinput.Model, 4),
				}
				return newState, nil
				// m.pageInputs = make([]textinput.Model, 4)

				// var t textinput.Model
				// for i := range m.pageInputs {
				// 	t = textinput.New()
				// 	t.Cursor.Style = cursorStyle
				// 	t.CharLimit = 32
				//
				// 	switch i {
				// 	// Home
				// 	case 0:
				// 		t.Prompt = "Home"
				// 		// t.Focus()
				// 		// t.PromptStyle = focusedStyle
				// 		// t.TextStyle = focusedStyle
				//
				// 		// Deposit
				// 	case 1:
				// 		t.Prompt = "Deposit"
				//
				// 		// Withdrawl
				// 	case 2:
				// 		t.Prompt = "Withdrawl"
				//
				// 	// Balance
				// 	case 3:
				// 		t.Prompt = "Balance"
				//
				// 	}
				// 	m.pageInputs[i] = t
				// }
				return m, nil
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.pageInputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.pageInputs)
			}

			cmds := make([]tea.Cmd, len(m.pageInputs))
			for i := 0; i <= len(m.pageInputs)-1; i++ {
				if i == m.focusIndex {
					// set focused state

					// Debug
					fmt.Printf("Page Inputs: %v", m.pageInputs[i].View())
					// Debug

					cmds[i] = m.pageInputs[i].Focus()
					m.pageInputs[i].PromptStyle = focusedStyle
					m.pageInputs[i].TextStyle = focusedStyle
					continue
				}
				// remove focused state
				m.pageInputs[i].Blur()
				m.pageInputs[i].PromptStyle = noStyle
				m.pageInputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}

	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	// Return the updated model to Tea for processing
	// Note: Not returning command
	return m, cmd
}

func (m *allowanceModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.pageInputs))

	// Only text inputs with Focus() set will response, so it's safe to simply
	// update all of them here without further logic
	for i := range m.pageInputs {
		m.pageInputs[i], cmds[i] = m.pageInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// View implements tea.Model.
func (m allowanceModel) View() string {
	var b strings.Builder

	switch m.selectedMenuOption {
	case PAGE_DEPOSIT:
		b.WriteString("Deposit\n")
		b.WriteString("-------\n")
		for i := range m.pageInputs {
			b.WriteString(m.pageInputs[i].View())
			if i < len(m.pageInputs)-1 {
				b.WriteRune('\n')
			}
		}
		button := &blurredButton
		if m.focusIndex == len(m.pageInputs) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	// PAGE_HOME
	default:
		// header
		b.WriteString("Select a menu option:\n")

		for i, option := range m.menuOptions {
			// Is the cursor here?
			cursor := " " // no cursor
			if m.focusIndex == i {
				cursor = ">" // cursor!
			}

			// render the row
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, option))

		}
	}

	// footer
	b.WriteString("\n Press q to quit.\n")

	// send to UI for rendering
	return b.String()
}
