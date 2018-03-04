package tgbot

import (
	"errors"
	"fmt"
)

type getMeResp struct {
	Ok          bool   `json:"ok"`
	Result      User   `json:"result"`
	Description string `json:"description"`
}

func GetMe(token string) (*User, error) {
	r := getMeResp{}
	url := fmt.Sprintf("%s/bot%s/%s", config.apiBase, token, "getMe")
	err := getRequest(url, &r)
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, errors.New(r.Description)
	}
	return &r.Result, nil
}
