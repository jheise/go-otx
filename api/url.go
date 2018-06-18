package api

import (
	"encoding/json"
	"fmt"
)

func (self *Client) GetUrlGeneral(url string) (UrlGeneral, error) {
	var data UrlGeneral
	body, err := get_indicator_data(self.key, "url", url, "general")
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return data, nil
}

func (self *Client) GetUrlUrlList(url string) (UrlUrlList, error) {
	var data UrlUrlList
	body, err := get_indicator_data(self.key, "url", url, "url_list")
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)

	fmt.Println(string(body))
	return data, nil
}
