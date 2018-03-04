package tgbot

import "fmt"

type WebhookInfo struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int      `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       int      `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

type getWebhookInfoResp struct {
	Ok          bool        `json:"ok"`
	Result      WebhookInfo `json:"result"`
	Description string      `json:"description"`
}

func GetWebhookInfo(token string) (*WebhookInfo, error) {
	r := getWebhookInfoResp{}
	url := fmt.Sprintf("%s/bot%s/%s", config.apiBase, token, "getWebhookInfo")
	err := getRequest(url, &r)
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, fmt.Errorf("GetWebhookInfo failed: %s\n", r.Description)
	}
	return &r.Result, nil
}

func SetWebhook(token, url string) error {
	r := response{}
	apiUrl := fmt.Sprintf("%s/bot%s/setWebhook?url=%s", config.apiBase, token, url)
	err := getRequest(apiUrl, &r)
	if err != nil {
		return err
	}
	if !r.Ok {
		return fmt.Errorf("SetWebhook failed: %s\n", r.Description)
	}
	return nil
}
