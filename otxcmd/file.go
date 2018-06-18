package main

import (
	"fmt"
)

func processFile(target string, section string) {
	switch section {
	case gen:
		handleFileGen(target)
	default:
		notImpl("File", section)
	}
}

func handleFileGen(target string) {
	data, err := client.GetFileGeneral(target)
	checkError("File", gen, err)

	printSections(data.Sections)
	printBaseIndicator(data.BaseIndicator)
	fmt.Printf("Indicator:\t%s\n", data.Indicator)
	printPulseInfo(data.PulseInfo)
	fmt.Printf("Type:\t%s\n", data.Type)
	fmt.Printf("Type Title:\t%s\n", data.TypeTitle)
	printValidation(data.Validation)
}
