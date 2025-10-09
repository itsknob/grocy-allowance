package main

import (
	"fmt"
	"os"

	"example.com/grocy-allowance/tui"
	tea "github.com/charmbracelet/bubbletea"
)

type GrocyConfig struct {
	GROCY_URL string
}

func main() {
	// config := GrocyConfig{
	// 	os.Getenv("GROCY_URL"),
	// }

	// grocyClient := grocy.NewGrocyClient(config.GROCY_URL)

	// var items *[]grocy.Item
	// var err error
	// items, err = grocyClient.GetStock()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, item := range *items {
	// 	// fmt.Println(item.Product.Name)
	// 	fmt.Printf("| %36s | %36d |\r\n", item.Product.Name, item.Product.Id)
	//
	// }
	// initialChoices := []string{"Buy Carrots", "Buy Celery", "Buy Potatoes"}
	initialChoices := []string{"Get Stock"}
	p := tea.NewProgram(tui.InitModel(initialChoices))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error: %v", err)
		os.Exit(1)
	}
}
