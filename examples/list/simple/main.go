package main

import (
	"fmt"

	"github.com/yokecd/lipgloss/list"
)

func main() {
	l := list.New(
		"A",
		"B",
		"C",
		list.New(
			"D",
			"E",
			"F",
		).Enumerator(list.Roman),
		"G",
	)
	fmt.Println(l)
}
