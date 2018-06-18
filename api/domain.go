package api

import (
	"encoding/json"
	// "fmt"
)

func (self *Client) GetDomainGeneral(domain string) (DomainGeneral, error) {
	var domaindata DomainGeneral
	body, err := get_indicator_data(self.key, "domain", domain, "general")
	if err != nil {
		return domaindata, err
	}
	json.Unmarshal(body, &domaindata)

	return domaindata, nil
}

func (self *Client) GetDomainGeo(domain string) (Geo, error) {
	var domaindata Geo
	body, err := get_indicator_data(self.key, "domain", domain, "geo")
	if err != nil {
		return domaindata, err
	}

	json.Unmarshal(body, &domaindata)

	return domaindata, nil
}

func (self *Client) GetDomainMalware(domain string) ([]MalwareData, error) {
	data, err := self.GetMalware("domain", domain)

	return data, err
}

func (self *Client) GetDomainUrlList(domain string) (DomainUrlList, error) {
	var domaindata DomainUrlList
	body, err := get_indicator_data(self.key, "domain", domain, "url_list")
	if err != nil {
		return domaindata, err
	}

	json.Unmarshal(body, &domaindata)

	return domaindata, nil
}

func (self *Client) GetDomainPassiveDns(domain string) (GenPassiveDns, error) {
	var domaindata GenPassiveDns
	body, err := get_indicator_data(self.key, "domain", domain, "passive_dns")
	if err != nil {
		return domaindata, err
	}

	json.Unmarshal(body, &domaindata)

	return domaindata, nil
}

// general, geo, malware, url_list, passive_dns)
