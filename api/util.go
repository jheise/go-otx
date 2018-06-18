package api

import (
	// standard
	"fmt"
	"io/ioutil"
	"net/http"
)

func get_indicator_data(api string, ioctype string, target string, section string) ([]byte, error) {
	url_dst := fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/%s/%s/%s", ioctype, target, section)
	client := new(http.Client)
	req, err := http.NewRequest("GET", url_dst, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("X-OTX-API-KEY", api)
	req.Header.Add("User-Agent", "Golang_go_otx/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
