package api

import (
	// standard
	"encoding/json"
)

func (self *Client) GetHostnameGeneral(hostname string) (HostnameGeneral, error) {
	var data HostnameGeneral
	body, err := get_indicator_data(self.key, "hostname", hostname, "general")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetHostnameGeo(hostname string) (Geo, error) {
	var data Geo
	body, err := get_indicator_data(self.key, "hostname", hostname, "geo")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)

	return data, nil
}

func (self *Client) GetHostnameMalware(hostname string) ([]MalwareData, error) {
	data, err := self.GetMalware("hostname", hostname)

	return data, err
}

func (self *Client) GetHostnameUrlList(hostname string) (HostnameUrlList, error) {
	var data HostnameUrlList
	body, err := get_indicator_data(self.key, "hostname", hostname, "url_list")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)

	return data, nil
}

func (self *Client) GetHostnamePassiveDns(hostname string) (GenPassiveDns, error) {
	var data GenPassiveDns
	body, err := get_indicator_data(self.key, "hostname", hostname, "passive_dns")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)

	return data, nil
}

// general, geo, malware, url_list, passive_dns)
