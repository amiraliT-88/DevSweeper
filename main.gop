package main

import (
	"fmt"
	"os"

	"devsweeper/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not get user home dir: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(ui.InitialModel(home))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
