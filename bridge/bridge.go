package bridge

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
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
	Log              *logrus.Logger
}

// New generates a default bridge
func New() *Bridge {
	return &Bridge{
		Client:           &http.Client{},
		UsernameOverride: DefaultUsername,
		IconURL:          DefaultIconURL,
		Log:              nil,
	}
}

// Handler will return the handler for use any ServerMux
func (b *Bridge) Handler(w http.ResponseWriter, r *http.Request) {
	hookURL := r.URL.Query().Get("mattermost_hook_url")
	channelOverride := r.URL.Query().Get("channel")

	if len(hookURL) < 0 {
		b.WriteError(w, r, "Missing mattermost_hook_url")
		return
	}

	hook, err := NewWebhookfromJSON(r.Body)
	if err != nil {
		b.WriteError(w, r, err.Error())
		return
	}

	data, err := NewMessageFromWebhook(hook, b, channelOverride).toJSON()
	if err != nil {
		b.WriteError(w, r, err.Error())
		return
	}

	res, err := b.Client.Post(hookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		b.WriteError(w, r, err.Error())
		return
	}

	defer res.Body.Close()
	ioutil.ReadAll(res.Body)
	b.WriteSuccess(w, r, "Webhook sent successfull")
}

// WriterError w
func (b *Bridge) WriteError(w http.ResponseWriter, r *http.Request, e string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request\n"))
}

// WriteSuccess w
func (b *Bridge) WriteSuccess(w http.ResponseWriter, r *http.Request, m string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}
