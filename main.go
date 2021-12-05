package discohook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// A webhook is a struct that contains the information needed to execute a webhook.
type Webhook struct {
	// You can get this from the channel integration setting menu.
	Url      string `json:"-"`
	Content  string `json:"content,omitempty"`
	ThreadId string `json:"-"`
	// Changes the shown username of webhook's message.
	Username string `json:"username,omitempty"`
	// Changes the shown avatar of webhook's message.
	AvatarURL     string `json:"avatar_url,omitempty"`
	Tts           bool   `json:"tts,omitempty"`
	AllowEveryone bool   `json:"-"`
	// THIS FIELD IS OVERWRITTEN, i have no idea how to export private fields to json.
	// If you do please help! cry about it lmao
	AllowedMentions map[string][]string `json:"allowed_mentions,omitempty"`
}

// Send() will execute the webhook, effectively sending a message to the channel it was created in.
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
