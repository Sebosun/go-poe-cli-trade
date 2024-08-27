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
	case "pc":
		priceCheck(text, state)
	case "all":
		printAllCurrency(state)
	case "export":
		exportToCSV(state)
	default:
		fmt.Println("Use pc to price check an item!")
	}

}
