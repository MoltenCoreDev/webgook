package webgook

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
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
	AvatarURL     string   `json:"avatar_url,omitempty"`
	Tts           bool     `json:"tts,omitempty"`
	Files         []string `json:"-"` // Paths to files to upload
	AllowEveryone bool     `json:"-"`
	// THIS FIELD IS GOING TO GET OVERWRITTEN cry about it lmao
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

	if w.Files == nil {
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

	} else {
		// TODO: Add support for multiple files
		f, _ := os.Open(w.Files[0])

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", filepath.Base(f.Name()))

		io.Copy(part, f)

		header := textproto.MIMEHeader{}
		header.Add("Content-Disposition", "form-data; name=payload_json")
		header.Add("Content-Type", "application-json")

		jsonWriter, _ := writer.CreatePart(header)
		jsonWriter.Write(jsonString)

		writer.Close()

		req, _ := http.NewRequest("POST", w.Url, body)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, _ := client.Do(req)
		resp.Body.Close()
	}
	return err
}

func main() {
	wh := Webhook{
		Url:      "https://discord.com/api/webhooks/917033935815475250/0twx5_h1S5E2hsiKwVpZ8iVmb--BKoOk8zSXX1YXfSwjbwtmYowHzfS2gH5q0gn8S7FJ",
		Username: "Use discohook today!",
		Content:  "Now with attachment support!",
		Files:    []string{"input.gif"},
	}
	wh.Send()
}
