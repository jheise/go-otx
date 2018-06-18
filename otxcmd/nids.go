package main

import (
	"fmt"
)

func processNids(target string, section string) {
	switch section {
	case gen:
		handleNidsGen(target)
	default:
		notImpl("Nids", target)
	}
}

func handleNidsGen(target string) {
	fmt.Println(client.GetNidsGeneral(target))
}
