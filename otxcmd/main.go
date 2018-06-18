package main

import (
	// standard
	"fmt"

	// local
	"github.com/jheise/go-otx/api"

	// external
	"github.com/mitchellh/go-homedir"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	apipath  = kingpin.Flag("apikey", "Path to API key").Default("~/.otxcmd").String()
	section  = kingpin.Flag("section", "Section to query, default general").Default("general").String()
	report   = kingpin.Arg("report", "What to report on [ipv4, domain, hostname]").Required().String()
	target   = kingpin.Arg("target", "What to report on").Required().String()
	sections = [6]string{"general", "reputation", "geo", "malware", "url_list", "passive_dns"}
	dsection = [4]string{"geo", "malware", "url_list", "passive_dns"}
	client   *api.Client
)

func init() {
	kingpin.Parse()
	var err error

	if *apipath == "~/.otxcmd" {
		keypath, err := homedir.Expand(*apipath)
		if err != nil {
			panic(err)
		}

		client, err = api.NewClientFromFile(keypath)
		if err != nil {
			panic(err)
		}

	} else {
		client, err = api.NewClientFromFile(*apipath)
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	switch *report {
	case "ipv4":
		process_ipv4(*target, *section)
	case "ipv6":
		process_ipv6(*target, *section)
	case domain:
		process_domain(*target, *section)
	case hostname:
		process_hostname(*target, *section)
	case "cve":
		processCve(*target, *section)
	case "file":
		processFile(*target, *section)
	case "url":
		processUrl(*target, *section)
	case "nids":
		processNids(*target, *section)
	default:
		fmt.Printf("Report %s not implemented\n", *report)
	}
}
