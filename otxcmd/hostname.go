package main

import (
	// standard
	"fmt"
)

func process_hostname(target string, section string) {
	switch section {
	case gen:
		handleHostGen(target)
	case geo:
		handleHostGeo(target)
	case passive_dns:
		fallthrough
	case dns:
		handleHostDns(target)
	case malware:
		handleHostMalware(target)
	default:
		notImpl("Hostname", section)
	}
}

func handleHostGen(target string) {
	data, err := client.GetHostnameGeneral(target)
	checkError("hostname", geo, err)

	printSections(data.Sections)
	fmt.Printf("Alexa:\t%s\n", data.Alexa)
	printBaseIndicator(data.BaseIndicator)
	fmt.Printf("Indicator:\t%s\n", data.Indicator)
	printPulseInfo(data.PulseInfo)
	printValidation(data.Validation)
	fmt.Printf("Whois:\t%s\n", data.Whois)

}

func handleHostGeo(target string) {
	data, err := client.GetHostnameGeo(target)
	checkError("hostname", geo, err)

	printGeoData(data)
}

func handleHostDns(target string) {
	data, err := client.GetHostnamePassiveDns(target)
	checkError("hostname", dns, err)

	printPassiveDns(data)
}

func handleHostMalware(target string) {
	data, err := client.GetHostnameMalware(target)
	checkError("hostname", malware, err)

	printMalwareCSV(data)
}
