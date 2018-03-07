package tgbot

import "fmt"

type sendMessageResponse struct {
	Ok          bool    `json:"ok"`
	Result      Message `json:"result"`
	Description string  `json:"description"`
}

func SendMessage(token string, opt SendMessageContent) (*Message, error) {
	u := fmt.Sprintf("%s/bot%s/%s", config.apiBase, token, "sendMessage")
	r := sendMessageResponse{}
	err := postRequest(u, opt, &r)
	if err != nil {
		return nil, fmt.Errorf("SendMessage failed: %s\n", err)
	}
	if !r.Ok {
		return nil, fmt.Errorf("SendMessage failed: %s\n", r.Description)
	}
	return &r.Result, nil
}
