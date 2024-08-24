package main

import (
	"fmt"
	"go-poe-trade/helpers"
	"os"
	"strings"
)

func isPriceCheck(text string) bool {
	inputSlice := strings.Split(text, " ")

	return inputSlice[0] == "pc"
}

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

	if isPriceCheck(text) {
		parsePriceCheck(text, state)
		return
	}

	if text == "all" {
		printAllCurrency(state)
		return
	}

	fmt.Println("Use pc to price check an item!")
}
