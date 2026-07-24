package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NotifyDiscord struct {
	webhookURL string
	client     *http.Client
}

type message struct {
	Content string `json:"content"`
}

func NewNotify(webhookURL string) Notify {
	return &NotifyDiscord{
		webhookURL: webhookURL,
		client:     &http.Client{},
	}
}

func (n *NotifyDiscord) Notify(content string) error {
	payload := message{
		Content: content,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	// discord通知
	resp, err := n.client.Post(
		n.webhookURL,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook failed: status = %d", resp.StatusCode)
	}
	return nil
}