# webgook
A simple discord go library to send webhooks.

# Example
```go
func main() {
package main

import "github.com/MoltenCoreDev/webgook"

	wh := Webhook{
		Url:      "https://discord.com/api/webhooks/id/token", // You can copy this from the channel settings, where you create the webhook
		Username: "Use discohook today!",
		Content:  "Now with attachment support!",
		Files:    []string{"input.gif"},
	}
	wh.Send()
}

```
