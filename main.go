package main // import "github.com/mroth/subtleist"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var webhookURL string

func slashUsage() string {
	return "Usage: `/socialrules [surprise|wellactually|backseat|subtle] [<@user>]`\n" +
		"Anonymously send a Recurse Center social rule either publicly to your current channel " +
		"or privately to a specific user."
}

func init() {
	webhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("SLACK_WEBHOOK_URL environment variable must be set!")
	}
}

// a message payload to send to a Slack incoming webhook
// https://api.slack.com/methods/chat.postMessage for field descriptions
// doesn't represent all possible fields, just the ones we are using
type slackMessage struct {
	Channel     string                   `json:"channel"`
	Text        string                   `json:"text,omitempty"`
	Attachments []slackMessageAttachment `json:"attachments,omitempty"`
	// we are intentionally not using the below, because we never want to override
	// what gets set as the default in the Slack Webhook configuration interface
	// Username    string                   `json:"username,omitempty"`
	// IconURL     string                   `json:"icon_url,omitempty"`
	// IconEmoji   string                   `json:"icon_emoji,omitempty"`
}

// a slack message attachment, see: https://api.slack.com/docs/attachments
// doesn't represent all possible attachment fields, just the ones we are using
type slackMessageAttachment struct {
	Pretext   string `json:"pretext,omitempty"`
	Text      string `json:"text,omitempty"`
	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`
}

func (r Rule) formatForPostingTo(destination string) []byte {
	var intro string
	if strings.HasPrefix(destination, "@") {
		intro = "Someone asked me to share this Recurse Center social rule with you."
	} else {
		intro = "Someone asked me to share this Recurse Center social rule with the group."
	}

	payload, _ := json.Marshal(slackMessage{
		Channel: destination,
		Attachments: []slackMessageAttachment{
			slackMessageAttachment{
				Pretext:   intro,
				Title:     r.title,
				TitleLink: r.uri(),
				Text:      r.description,
			},
		},
	})
	return payload
}

// Extract the user command sent and the destination for sending.
// command is always the first word
// destination is either user-specified, or the channel the user was in
func extractParams(req *http.Request) (cmd string, dest string) {
	textWords := strings.Split(req.FormValue("text"), " ")
	cmd = textWords[0]
	if len(textWords) >= 2 {
		dest = textWords[1]
	} else {
		dest = req.FormValue("channel_id")
	}
	return
}

// Handle incoming slack slash commands.
func slackHandler(rw http.ResponseWriter, req *http.Request) {
	cmd, dest := extractParams(req)

	// if command is absent or "help", echo back usage.
	if cmd == "" || cmd == "help" {
		fmt.Fprintln(rw, slashUsage())
		return
	}

	// otherwise, check against all possible rule triggers as commands
	for _, r := range rules {
		if strings.Contains(cmd, r.trigger) {
			// post to destination
			payload := r.formatForPostingTo(dest)
			req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				// if something goes wrong in slack webhook, body has descriptive error
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Fprintf(rw, "Something went wrong (%v): %v", resp.Status, body)
				log.Printf("WARNING: %v from Slack webhook: %v\n", resp.Status, body)
			}

			// echo success back if sent to user privately (so they know it worked)
			// only for private msgs since public messages are their own confirmation
			if strings.HasPrefix(dest, "@") {
				fmt.Fprintf(rw, `Okay, I sent "%v" to <%v>`, r.title, dest)
			}
			return
		}
	}

	// we didnt recognize any of the commands, remind user of usage
	fmt.Fprintf(rw, "Sorry I don't know about `%v`.\n%v", cmd, slashUsage())
}

func main() {
	http.HandleFunc("/slack_hook", slackHandler)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
