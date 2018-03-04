package tgbot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

func getRequest(url string, into interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, &into)
}
