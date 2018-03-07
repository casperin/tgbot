package tgbot

import (
	"bytes"
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
	if into == nil {
		return nil
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, &into)
}

func postRequest(url string, params interface{}, into interface{}) error {
	jsonValue, err := json.Marshal(params)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if into == nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &into)
}
