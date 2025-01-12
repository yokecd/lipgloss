package main

import (
	"fmt"

	"github.com/yokecd/lipgloss"
	"github.com/yokecd/lipgloss/list"
)

func main() {
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255")).MarginRight(1)

	l := list.New(
		"Glossier",
		"Claire’s Boutique",
		"Nyx",
		"Mac",
		"Milk",
	).
		Enumerator(list.Roman).
		EnumeratorStyle(enumeratorStyle).
		ItemStyle(itemStyle)

	fmt.Println(l)
}
