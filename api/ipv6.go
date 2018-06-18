package api

import (
	"encoding/json"
	"fmt"
)

func (self *Client) GetIPv6General(ipv6 string) (IPv6General, error) {
	var data IPv6General
	body, err := get_indicator_data(self.key, "IPv6", ipv6, "general")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPv6Geo(ipv6 string) (Geo, error) {
	var data Geo
	body, err := get_indicator_data(self.key, "IPv6", ipv6, "geo")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPv6Reputation(ipv6 string) (IPv6Reputation, error) {
	var data IPv6Reputation
	body, err := get_indicator_data(self.key, "IPv6", ipv6, "reputation")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}

func (self *Client) GetIPv6Malware(ipv6 string) (Malware, error) {
	var data Malware
	body, err := get_indicator_data(self.key, "IPv6", ipv6, "malware")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}

func (self *Client) GetIPv6UrlList(ipv6 string) (IPv6UrlList, error) {
	var data IPv6UrlList
	body, err := get_indicator_data(self.key, "IPv6", ipv6, "url_list")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}

func (self *Client) GetIPv6PassiveDns(ipv6 string) (GenPassiveDns, error) {
	var data GenPassiveDns
	body, err := get_indicator_data(self.key, "IPv6", ipv6, "passive_dns")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}
