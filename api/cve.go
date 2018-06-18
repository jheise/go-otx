package api

import (
	"encoding/json"
	"fmt"
)

func (self *Client) GetCveGeneral(cve string) (CveGeneral, error) {
	var data CveGeneral
	body, err := get_indicator_data(self.key, "cve", cve, "general")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}
