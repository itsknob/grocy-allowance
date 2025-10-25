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

// TUI Styles
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

type Transaction struct {
	amount      float32
	date        string
	description string
}

type DepositModel struct { // Deposit Page State
	currentTransaction Transaction
	focusedInputIndex  int
	focusedInput       textinput.Model
	inputs             map[int]textinput.Model // contain their own values
}

type WithdrawlModel struct {
	currentTransaction Transaction
	focusedInputIndex  int
	focusedInput       textinput.Model
	inputs             map[int]textinput.Model // contain their own values
}

type BalanceModel struct {
	focusedInputIndex int
	focusedInput      textinput.Model
	inputs            map[int]textinput.Model // contain their own values
	transactions      []Transaction
}

type UserModel struct {
	name         string
	balance      float32
	transactions []Transaction
}

func getInitialDepositModel() DepositModel {
	initialInputs := make(map[int]textinput.Model, 4)

	m := DepositModel{
		inputs:             initialInputs,
		focusedInput:       initialInputs[0],
		focusedInputIndex:  0,
		currentTransaction: Transaction{},
	}

	var t textinput.Model
	for i := range m.inputs {
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
		case 3:
			t.Prompt = "Description:"
			t.Width = 32
			t.Placeholder = "Reason"
		}
		m.inputs[i] = t
	}

	return m
}

type ApplicationModel struct { // Application State
	focusedInput      textinput.Model         // Current Input
	focusedInputIndex int                     // Current Position
	inputs            map[int]textinput.Model // Menu Options
	grocyClient       grocy.GrocyClient       // Client
}

/**
* Sets all fields on the Model
* Set's Prompts for textinputs to be menu options
 */
func GetInitialApplicationModel() ApplicationModel {
	initialInputs := make(map[int]textinput.Model, 4)

	// Setup Inputs with Labels
	// TOOD: These should not be textinputs.
	for idx, current := range initialInputs {
		switch idx {
		case 0:
			current.Prompt = "Home"
		case 1:
			current.Prompt = "Deposit"
		case 2:
			current.Prompt = "Withdrawl"
		case 3:
			current.Prompt = "Balance"
		}
	}

	// Set up Grocy Client
	config := grocy.GrocyConfig{
		GROCY_URL: os.Getenv("GROCY_URL"),
	}
	if config.GROCY_URL == "" {
		log.Fatal("GROCY_URL env variable is not set.")
	}

	grocyClient := grocy.NewGrocyClient(config.GROCY_URL)

	// Ensure Allowance Entity Exists in Grocy
	if !grocyClient.HasAllowance() {
		log.Default().Println("Allowance not found. Setting up new userentity")
		grocyClient.InitAllowance()
	}

	return ApplicationModel{
		focusedInputIndex: 0,
		inputs:            initialInputs,
		focusedInput:      initialInputs[0],
		grocyClient:       *grocyClient,
	}
}

func (m ApplicationModel) Init() tea.Cmd {
	if m.focusedInput.Value() == PAGE_HOME || m.focusedInput.Value() == PAGE_BALANCE {
		GetInitialApplicationModel()
		return cursor.Blink
	} else {
		return textinput.Blink
	}
}

// // Initialize the Application State, Return Model
// // Deprecated
// func InitialModel() ApplicationModel {
// 	m := ApplicationModel{
// 		inputs: make(map[int]textinput.Model, 3),
// 	}
//
// 	config := grocy.GrocyConfig{
// 		GROCY_URL: os.Getenv("GROCY_URL"),
// 	}
// 	if config.GROCY_URL == "" {
// 		log.Fatal("GROCY_URL env variable is not set.")
// 	}
//
// 	grocyClient := grocy.NewGrocyClient(config.GROCY_URL)
//
// 	// Ensure Allowance Exists
// 	if !grocyClient.HasAllowance() {
// 		log.Default().Println("Allowance not found. Setting up new userentity")
// 		grocyClient.InitAllowance()
// 	}
//
// 	// Initial Menu Options
// 	m.inputs = []string{PAGE_HOME, PAGE_DEPOSIT, PAGE_WITHDRAWL, PAGE_BALANCE}
// 	m.focusedInput = PAGE_HOME
// 	m.grocyClient = *grocyClient
//
// 	return m
// }

func (m DepositModel) initDepositModel() tea.Cmd {
	// User, Amount, Description, Date
	m.inputs = make(map[int]textinput.Model, 4)

	// Generate Text Input for each input on the page
	var t textinput.Model
	for i := range m.inputs {
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
		case 3:
			t.Prompt = "Description:"
			t.Width = 32
			t.Placeholder = "Reason"
		}

		m.inputs[i] = t

	}

	return nil
}

/**
* Called when things happen
* It's a Reducer
 */
func (m ApplicationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Switch on Type of Message Received
	switch msg := msg.(type) {
	// is it a kepress?
	case tea.KeyMsg:
		switch msg.String() {
		// Quit
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			m.focusedInputIndex--
			if m.focusedInputIndex < 0 {
				m.focusedInputIndex = len(m.inputs) - 1
			}
		case "down":
			m.focusedInputIndex++
			if m.focusedInputIndex > len(m.inputs)-1 {
				m.focusedInputIndex = 0
			}

		// Update model with selection when enter pressed
		case "enter":
			switch m.focusedInputIndex {
			case 0:
				return updateHomePage(msg, m)
			case 1:
				return updateDepositPage(msg, m)
			case 2:
				return updateWithdrawlPage(msg, m)
			case 3:
				return updateBalancePage(msg, m)
			}
			// Update the Selected Page
			if m.focusedInputIndex != 0 { // not empty or Home
				m.focusedInput = m.inputs[m.focusedInputIndex]
				return updateSelectedPage(msg, m)
			}
			return updateHomePage(msg, m)
		}
	}
	return m, nil
}

func updateBalancePage(msg tea.KeyMsg, m ApplicationModel) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

func updateWithdrawlPage(msg tea.KeyMsg, m ApplicationModel) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

func updateDepositPage(msg tea.KeyMsg, m ApplicationModel) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

func updateSelectedPage(msg tea.KeyMsg, m ApplicationModel) (tea.Model, tea.Cmd) {
	log.Default().Println("Called updateSelectedPage")
	panic("unimplemented")
}

func updateHomePage(msg tea.KeyMsg, m ApplicationModel) (tea.Model, tea.Cmd) {
	log.Default().Println("Called updateHomePage")
	return GetInitialApplicationModel(), nil
}

// func (m ApplicationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	// Switch on Type of Message Received
// 	switch msg := msg.(type) {
// 	// is it a kepress?
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		case "tab", "shift+tab", "enter", "up", "down":
// 			s := msg.String()
// 			// Did user press enter while submit button was focused? Exists
// 			if s == "enter" && m.focusIndex == len(m.pageInputs) && m.selectedMenuOption != PAGE_HOME {
// 				log.Default().Println("Submitted")
//
// 				return m, nil
// 			}
//
// 			// 0. Home
// 			// 1. Deposit
// 			if s == "enter" && m.focusIndex == 1 {
// 				newState := allowanceModel{
// 					selectedMenuOption: PAGE_DEPOSIT,
// 					pageInputs:         make([]textinput.Model, 4),
// 				}
// 				return newState, nil
// 				// m.pageInputs = make([]textinput.Model, 4)
//
// 				// var t textinput.Model
// 				// for i := range m.pageInputs {
// 				// 	t = textinput.New()
// 				// 	t.Cursor.Style = cursorStyle
// 				// 	t.CharLimit = 32
// 				//
// 				// 	switch i {
// 				// 	// Home
// 				// 	case 0:
// 				// 		t.Prompt = "Home"
// 				// 		// t.Focus()
// 				// 		// t.PromptStyle = focusedStyle
// 				// 		// t.TextStyle = focusedStyle
// 				//
// 				// 		// Deposit
// 				// 	case 1:
// 				// 		t.Prompt = "Deposit"
// 				//
// 				// 		// Withdrawl
// 				// 	case 2:
// 				// 		t.Prompt = "Withdrawl"
// 				//
// 				// 	// Balance
// 				// 	case 3:
// 				// 		t.Prompt = "Balance"
// 				//
// 				// 	}
// 				// 	m.pageInputs[i] = t
// 				// }
// 				return m, nil
// 			}
//
// 			if s == "up" || s == "shift+tab" {
// 				m.focusIndex--
// 			} else {
// 				m.focusIndex++
// 			}
//
// 			if m.focusIndex > len(m.pageInputs) {
// 				m.focusIndex = 0
// 			} else if m.focusIndex < 0 {
// 				m.focusIndex = len(m.pageInputs)
// 			}
//
// 			cmds := make([]tea.Cmd, len(m.pageInputs))
// 			for i := 0; i <= len(m.pageInputs)-1; i++ {
// 				if i == m.focusIndex {
// 					// set focused state
//
// 					// Debug
// 					fmt.Printf("Page Inputs: %v", m.pageInputs[i].View())
// 					// Debug
//
// 					cmds[i] = m.pageInputs[i].Focus()
// 					m.pageInputs[i].PromptStyle = focusedStyle
// 					m.pageInputs[i].TextStyle = focusedStyle
// 					continue
// 				}
// 				// remove focused state
// 				m.pageInputs[i].Blur()
// 				m.pageInputs[i].PromptStyle = noStyle
// 				m.pageInputs[i].TextStyle = noStyle
// 			}
// 			return m, tea.Batch(cmds...)
// 		}
// 	}
//
// 	// Handle character input and blinking
// 	cmd := m.updateInputs(msg)
//
// 	// Return the updated model to Tea for processing
// 	// Note: Not returning command
// 	return m, cmd
// }

func (m *ApplicationModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will response, so it's safe to simply
	// update all of them here without further logic
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// View implements tea.Model.
func (m ApplicationModel) View() string {
	var b strings.Builder

	switch m.focusedInput.Value() {
	case PAGE_DEPOSIT:
		b.WriteString("Deposit\n")
		b.WriteString("-------\n")
		for i := range m.inputs {
			b.WriteString(m.inputs[i].View())
			if i < len(m.inputs)-1 {
				b.WriteRune('\n')
			}
		}
		button := &blurredButton
		if m.focusedInputIndex == len(m.inputs) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	// PAGE_HOME
	default:
		// header
		b.WriteString("Select a menu option:\n")

		for i, option := range m.inputs {
			// Is the cursor here?
			cursor := " " // no cursor
			if m.focusedInputIndex == i {
				cursor = ">" // cursor!
			}

			// render the row
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, option.Value()))

		}
	}

	// footer
	b.WriteString("\n Press q to quit.\n")
	b.WriteString(fmt.Sprintf("\nCurrent Index: %d", m.focusedInputIndex))

	// send to UI for rendering
	return b.String()
}
