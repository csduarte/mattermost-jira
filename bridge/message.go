package bridge

import "encoding/json"

// Message structure for Mattermost JSON creation.
type Message struct {
	Text     string `json:"text"`
	Channel  string `json:"channel,omitempty"`
	Username string `json:"username"`
	IconURL  string `json:"icon_url"`
}

// NewMessageFromWebhook constructs for a basic message for uChat
func NewMessageFromWebhook(w *Webhook, b *Bridge, channel string) *Message {
	return &Message{
		Text:     w.String(),
		Channel:  channel,
		Username: b.UsernameOverride,
		IconURL:  b.IconURL,
	}
}

func (m *Message) toJSON() ([]byte, error) {
	return json.Marshal(m)
}
