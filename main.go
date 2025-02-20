package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var cleanInputs []string
	cleanInputs = strings.Fields(strings.ToLower(text))
	return cleanInputs
}