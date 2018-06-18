package api

import (
	"encoding/json"
	"fmt"
)

func (self *Client) GetNidsGeneral(nids string) (NidsGeneral, error) {
	var data NidsGeneral
	body, err := get_indicator_data(self.key, "nids", nids, "general")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}
