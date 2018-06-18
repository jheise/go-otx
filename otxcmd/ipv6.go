package main

import (
	"fmt"
)

func process_ipv6(target string, section string) {
	switch section {
	case gen:
		handleIPv6Gen(target)
	default:
		notImpl("IPv6", section)
	}
}

func handleIPv6Gen(target string) {
	fmt.Println(client.GetIPv6General(target))
}
