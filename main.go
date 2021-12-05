package discohook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Webhook struct {
	Url           string `json:"-"`
	Content       string `json:"content,omitempty"`
	ThreadId      string `json:"-"`
	Username      string `json:"username,omitempty"`
	AvatarURL     string `json:"avatar_url,omitempty"`
	Tts           bool   `json:"tts,omitempty"`
	AllowEveryone bool   `json:"-"`
	// THIS FIELD IS OVERWRITTEN, i have no idea how to export private fields to json.
	// If you do please help! cry about it lmao
	AllowedMentions map[string][]string `json:"allowed_mentions,omitempty"`
}

// Sends the webhook
func (w *Webhook) Send() error {
	if w.ThreadId != "" {
		w.Url = w.Url + "?thread_id=" + w.ThreadId
	}
	if w.AllowEveryone {
		w.AllowedMentions = nil
	} else {
		w.AllowedMentions = map[string][]string{"parse": {""}}
	}

	jsonString, err := json.Marshal(w)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", w.Url, bytes.NewBuffer(jsonString))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
