package tui

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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

type DepositModel struct {
	currentTransaction Transaction
	focusedInputIndex  int
	focusedInput       textinput.Model
	inputs             map[int]textinput.Model // contain their own values
	currentUser        UserModel
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

func (dm DepositModel) Init() tea.Cmd {
	return nil
}

func (dm DepositModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return dm, nil
}

func (dm DepositModel) View() string {
	return "Deposit Model"
}

func GetInitialDepositModel() DepositModel {
	log.Default().Println("Called getInitialDepositModel")
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
