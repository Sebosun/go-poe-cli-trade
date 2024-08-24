package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func isPriceCheck(text string) bool {
	inputSlice := strings.Split(text, " ")

	if inputSlice[0] == "pc" {
		return true
	}

	return false
}

func replParse(text string, state *State) {
	if text == "" {
		return
	}
	exitCommands := []string{"exit", "e", "quit", "Exit", ":q", "close"}
	_, isExitCommand := slices.BinarySearch(exitCommands, text)

	if isExitCommand {
		os.Exit(0)
		return
	}

	if isPriceCheck(text) {
		parsePriceCheck(text, state)
		return
	}

	fmt.Println("# No valid commant provided")
}
