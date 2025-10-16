package tui

import (
	"fmt"
	"log"
	"os"

	"example.com/grocy-allowance/grocy"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* @See https://github.com/charmbracelet/bubbletea/tree/main/tutorials/basics
*
* Init - a function that returns initial command to run app
* Update - a function that handles incoming events and updates the model
* View - a function that renders the UI based on data in model
 */

type model struct {
	choices     []string
	cursor      int
	selected    map[int]struct{}
	grocyClient grocy.GrocyClient
}

func (m model) Init() tea.Cmd {
	if m.grocyClient.HasAllowance() {
		return nil
	} else {
		m.grocyClient.InitAllowance()
		return nil
	}
}

func InitModel(initialChoices []string) model {
	config := grocy.GrocyConfig{
		GROCY_URL: os.Getenv("GROCY_URL"),
	}
	if config.GROCY_URL == "" {
		log.Fatal("GROCY_URL env variable is not set.")
	}

	grocyClient := grocy.NewGrocyClient(config.GROCY_URL)

	return model{
		choices:     initialChoices,
		selected:    make(map[int]struct{}),
		grocyClient: *grocyClient,
	}
}

func (m *model) printStock() {

	var items *[]grocy.StockEntry
	var err error
	items, err = m.grocyClient.GetStock()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("| ------------------------------------ | ------------------------------------ | ---------- |\r\n")
	fmt.Printf("| %-36s | %-36s | %10s |\r\n", "Name", "Quantity", "Units")
	fmt.Printf("| ------------------------------------ | ------------------------------------ | ---------- |\r\n")
	for _, item := range *items {
		// fmt.Println(item.Product.Name)
		units := m.grocyClient.GetUnits()
		if units == "" {
			units = "Units"
		}
		fmt.Printf("| %36s | %36f | %10s |\r\n", item.StockId, item.Amount, "units")

	}
	fmt.Printf("| ------------------------------------ | ------------------------------------ |\r\n")
}

func (m *model) printAllowancePage() {
	log.Fatal("Not implemented")
}

// func (m model) Init() tea.Cmd {
// 	return nil // no I/O right now
// }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// is it a kepress?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				// delete(m.selected, m.cursor)
				m.printAllowancePage()
			} else {
				// m.printStock()
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	// Return the updated model to Tea for processing
	// Note: Not returning command
	return m, nil
}

func (m model) AllowanceViewMain() {

}

// View implements tea.Model.
func (m model) View() string {
	// header
	s := "Select a menu option:\n"

	for i, choice := range m.choices {
		// Is the cursor here?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// is it selected?
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		// render the processing
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}

	// footer
	s += "\n Press q to quit.\n"

	// send to UI for rendering
	return s
}
