package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	state := State{}

	state.initState()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()

		text := scanner.Text()
		if len(text) != 0 {
			text = strings.Trim(text, " ")
			text = strings.ToLower(text)
			replParse(text, &state)
		}
	}
}
