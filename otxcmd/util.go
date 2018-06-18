package main

import (
	"fmt"
	"os"
	"time"

	// local
	"github.com/jheise/go-otx/api"
)

// make a number of variables
var (
	gen         = "general"
	geo         = "geo"
	dns         = "dns"
	passive_dns = "passive_dns"
	malware     = "malware"
	ipv4        = "IPv4"
	domain      = "domain"
	hostname    = "hostname"
)

func notImpl(report string, target string) {
	fmt.Printf("%s: Section %s not implemented\n", report, target)
}

func checkError(report string, section string, err error) {
	if err != nil {
		fmt.Printf("Unable to retrieve %s section %s. %v", report, section, err)
		os.Exit(1)
	}
}

func printPulseInfo(data api.PulseInfo) {
	fmt.Println("Pulse Info:")
	fmt.Printf("  Count: %d\n", data.Count)
	fmt.Printf("  Pulses:\n")
	for x, pulse := range data.Pulses {
		printPulse(pulse, x)
	}
}

func printPulse(pulse api.Pulse, count int) {
	fmt.Printf("    Pulse %d:\n", count+1)
	fmt.Printf("      - %+v\n", pulse)
}

func printSections(sections []string) {
	fmt.Println("Sections:")
	for _, section := range sections {
		fmt.Printf("  %s\n", section)
	}
}

func printValidation(validations []api.Validation) {
	fmt.Println("Validation:")
	for _, validation := range validations {
		fmt.Printf("  Message:  %s\n", validation.Message)
		fmt.Printf("  Name:  %s\n", validation.Name)
		fmt.Printf("  Source:  %s\n", validation.Source)
	}
}

func printGeoData(data api.Geo) {
	fmt.Printf("Area Code:\t%d\n", data.AreaCode)
	fmt.Printf("ASN:\t%s\n", data.Asn)
	fmt.Printf("Charset:\t%d\n", data.Charset)
	fmt.Printf("City:\t%s\n", data.City)
	fmt.Printf("City Data:\t%t\n", data.CityData)
	fmt.Printf("Continent Code:\t%s\n", data.ContinentCode)
	fmt.Printf("Country Code:\t%s\n", data.CountryCode)
	fmt.Printf("Country Code(3):\t%s\n", data.CountryCode3)
	fmt.Printf("Country Name:\t%s\n", data.CountryName)
	fmt.Printf("DMA Code:\t%d\n", data.DmaCode)
	fmt.Printf("Flag Title:\t%s\n", data.FlagTitle)
	fmt.Printf("Flag Url:\t%s\n", data.FlagUrl)
	fmt.Printf("Latitude:\t%f\n", data.Latitude)
	fmt.Printf("Longitude:\t%f\n", data.Longitude)
	fmt.Printf("Postal Code:\t%s\n", data.PostalCode)
	fmt.Printf("Region:\t%s\n", data.Region)
}

func printPassiveDns(data api.GenPassiveDns) {
	fmt.Printf("Count:\t%d\n", data.Count)
	for _, entry := range data.PassiveDns {
		fmt.Printf("  Address:  %s\n", entry.Address)
		fmt.Printf("  Asset Type:  %s\n", entry.AssetType)
		fmt.Printf("  First: %s\n", entry.First)
		fmt.Printf("  Hostname: %s\n", entry.Hostname)
		fmt.Printf("  Indicator: %s\n", entry.IndicatorLink)
		fmt.Printf("  Last: %s\n", entry.Last)
		fmt.Println()
	}
}

func printBaseIndicator(data api.BaseIndicator) {
	fmt.Println(data)
}

func printMalwareCSV(data []api.MalwareData) {
	for _, mwd := range data {
		fmt.Printf("%d,%s\n", mwd.DatetimeInt, mwd.Hash)
	}
}

func printMalwareTime(data []api.MalwareData) {
	for _, mwd := range data {
		fmt.Printf("%v - %s\n", time.Unix(int64(mwd.DatetimeInt), 0), mwd.Hash)
	}
}
