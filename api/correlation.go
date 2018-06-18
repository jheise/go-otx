package api

import (
	"encoding/json"
	"fmt"
)

func (self *Client) GetCorrelationRuleGeneral(correl string) (CorrelationRuleGeneral, error) {
	var data CorrelationRuleGeneral
	body, err := get_indicator_data(self.key, "correlation-rule", correl, "general")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}
