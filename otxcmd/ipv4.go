package main

import (
	// standard
	"fmt"
)

func process_ipv4(target string, section string) {
	switch section {
	case gen:
		handleipv4gen(target)
	case geo:
		handleipv4geo(target)
	case passive_dns:
		fallthrough
	case dns:
		handleipv4dns(target)
	case malware:
		handleipv4malware(target)
	default:
		notImpl(ipv4, section)
	}
}

func handleipv4gen(target string) {
	data, err := client.GetIPv4General(target)
	checkError(ipv4, gen, err)

	printSections(data.Sections)
	fmt.Printf("ASN:\t%s\n", data.Asn)
	fmt.Printf("Base Indicator:\n")
	fmt.Printf("City:\t%s\n", data.City)
	fmt.Printf("Country Name:\t%s\n", data.CountryName)
	printPulseInfo(data.PulseInfo)
	fmt.Printf("Type:\t%s\n", data.Type)
	fmt.Printf("Type Title:\t%s\n", data.TypeTitle)
	printValidation(data.Validation)
	fmt.Printf("Whois:\t%s\n", data.Whois)
}

func handleipv4geo(target string) {
	data, err := client.GetIPv4Geo(target)
	checkError(ipv4, geo, err)

	printGeoData(data)
}

func handleipv4dns(target string) {
	data, err := client.GetIPv4PassiveDns(target)
	checkError(ipv4, dns, err)

	printPassiveDns(data)
}

func handleipv4malware(target string) {
	data, err := client.GetIPv4Malware(target)
	checkError(ipv4, malware, err)

	printMalwareCSV(data)
}
