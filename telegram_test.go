package telegram

import "testing"

func TestNewTelegramWithoutMarkdown(t *testing.T) {
	telegram := New("139:AAH", "11788", false)

	expected := "https://api.telegram.org/bot139:AAH/sendMessage"
	actual := telegram.endpoint
	if actual != expected {
		t.Errorf("Expected endpoint to be '%s', but got '%s'", expected, actual)
	}
}

func TestNewTelegramWithMarkdown(t *testing.T) {
	telegram := New("139:AAH", "11788", true)

	expected := "https://api.telegram.org/bot139:AAH/sendMessage?parse_mode=Markdown"
	actual := telegram.endpoint
	if actual != expected {
		t.Errorf("Expected endpoint to be '%s', but got '%s'", expected, actual)
	}
}
