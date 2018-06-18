package main

import (
	"fmt"
)

func processUrl(target string, section string) {
	switch section {
	case gen:
		handleUrlGen(target)
	default:
		notImpl("URL", section)
	}
}

func handleUrlGen(target string) {
	fmt.Println(client.GetUrlGeneral(target))
}
