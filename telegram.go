package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "https://api.telegram.org/bot%s/sendMessage"

type Telegram struct {
	Token    string
	ChatID   string
	Markdown bool
	endpoint string
}

func New(token, chatID string, markdown bool) *Telegram {
	e := fmt.Sprintf(endpoint, token)
	if markdown {
		e += "?parse_mode=Markdown"
	}
	return &Telegram{
		Token:    token,
		ChatID:   chatID,
		endpoint: e,
	}
}

func (t *Telegram) Send(text string) error {
	payload := struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatID: t.ChatID,
		Text:   text,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(t.endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification with status: %s", resp.Status)
	}
	return nil
}
