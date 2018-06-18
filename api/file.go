package api

import (
	"encoding/json"
)

func (self *Client) GetFileGeneral(filename string) (FileGeneral, error) {
	var data FileGeneral
	body, err := get_indicator_data(self.key, "file", filename, "general")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)

	return data, nil
}

func (self *Client) GetFileAnalysis(filename string) (FileAnalysis, error) {
	var data FileAnalysis
	body, err := get_indicator_data(self.key, "file", filename, "analysis")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)

	return data, nil
}
