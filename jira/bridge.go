package jira

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// DefaultIconURL w
	DefaultIconURL = "https://raw.githubusercontent.com/csduarte/mattermost-jira/master/assets/logo.png"
	// DefaultUsername w
	DefaultUsername = "JIRA"
)

// Bridge sturcture will hold Jira Bridge data and settings
type Bridge struct {
	Client           *http.Client
	UsernameOverride string
	IconURL          string
}

// NewBridge generates a default bridge
func NewBridge() *Bridge {
	return &Bridge{
		Client:           &http.Client{},
		UsernameOverride: DefaultUsername,
		IconURL:          DefaultIconURL,
	}
}

// Handler will return the handler for use any ServerMux
func (b *Bridge) Handler(w http.ResponseWriter, r *http.Request) {
	mattermostHookURL := r.URL.Query().Get("mattermost_hook_url")
	channelOverride := r.URL.Query().Get("channel")

	if len(mattermostHookURL) < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request\n"))
		return
	}

	hook, err := NewWebhookfromJSON(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := NewMessageFromWebhook(hook, b, channelOverride).toJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := b.Client.Post(mattermostHookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	ioutil.ReadAll(res.Body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}
