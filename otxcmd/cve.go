package main

import (
	"fmt"
)

func processCve(target string, section string) {
	switch section {
	case gen:
		handleCveGen(target)
	default:
		notImpl("CVE", section)
	}
}

func handleCveGen(target string) {
	fmt.Println(client.GetCveGeneral(target))
}
