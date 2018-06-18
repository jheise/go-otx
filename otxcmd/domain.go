package main

import (
	// standard
	"fmt"
)

func process_domain(target string, section string) {
	switch section {
	case gen:
		handleDomainGen(target)
	case geo:
		handleDomainGeo(target)
	case passive_dns:
		fallthrough
	case dns:
		handleDomainDns(target)
	case malware:
		handleDomainMalware(target)
	default:
		notImpl(domain, section)
	}
}

func handleDomainGen(target string) {
	data, err := client.GetDomainGeneral(target)
	checkError(domain, gen, err)

	// print selected general fields
	printSections(data.Sections)
	fmt.Printf("Alexa:\t%s\n", data.Alexa)
	printBaseIndicator(data.BaseIndicator)
	fmt.Printf("Indicator:\t%s\n", data.Indicator)
	printPulseInfo(data.PulseInfo)
	printValidation(data.Validation)
	fmt.Printf("Whois:\t%s\n", data.Whois)
}

func handleDomainGeo(target string) {
	data, err := client.GetDomainGeo(target)
	checkError(domain, geo, err)

	printGeoData(data)
}

func handleDomainDns(target string) {
	data, err := client.GetDomainPassiveDns(target)
	checkError(domain, dns, err)

	printPassiveDns(data)
}

func handleDomainMalware(target string) {
	data, err := client.GetDomainMalware(target)
	checkError(domain, malware, err)

	printMalwareCSV(data)
}
