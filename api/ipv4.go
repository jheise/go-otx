package api

import (
	"encoding/json"
	// "fmt"
)

func (self *Client) GetIPv4General(ipv4 string) (IPv4General, error) {
	var data IPv4General
	body, err := get_indicator_data(self.key, "IPv4", ipv4, "general")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPv4Geo(ipv4 string) (Geo, error) {
	var data Geo
	body, err := get_indicator_data(self.key, "IPv4", ipv4, "geo")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPv4Reputation(ipv4 string) (IPv4Reputation, error) {
	var data IPv4Reputation
	body, err := get_indicator_data(self.key, "IPv4", ipv4, "reputation")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPv4Malware(ipv4 string) ([]MalwareData, error) {
	var data []MalwareData
	data, err := self.GetMalware("IPv4", ipv4)
	if err != nil {
		return data, err
	}

	return data, err
}

func (self *Client) GetIPv4UrlList(ipv4 string) (IPv4UrlList, error) {
	var data IPv4UrlList
	body, err := get_indicator_data(self.key, "IPv4", ipv4, "url_list")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPv4PassiveDns(ipv4 string) (GenPassiveDns, error) {
	var data GenPassiveDns
	body, err := get_indicator_data(self.key, "IPv4", ipv4, "passive_dns")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}
