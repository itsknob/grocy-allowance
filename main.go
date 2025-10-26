package main

import (
	"fmt"
	"os"

	"example.com/grocy-allowance/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dm := tui.GetInitialDepositModel()

	p := tea.NewProgram(dm)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
