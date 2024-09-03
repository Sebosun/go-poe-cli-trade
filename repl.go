package main

import (
	"fmt"
	"go-poe-trade/helpers"
	"os"
	"strings"
)

func replParse(text string, state *State) {
	if text == "" {
		return
	}

	exitCommands := []string{"exit", "e", "quit", "Exit", ":q", "close"}
	_, isExitCommand := helpers.FindStr(exitCommands, strings.ToLower(text))

	if isExitCommand {
		os.Exit(0)
		return
	}

	args := strings.Split(text, " ")
	cmd := args[0]

	switch cmd {
	case "test":
		state.items.ParseSharedNames()
	case "pc":
		priceCheck(text, state)
	case "all":
		state.printAllCurrency()
	case "export":
		state.exportToCSV()
	default:
		fmt.Println("Use pc to price check an item!")
	}

}
