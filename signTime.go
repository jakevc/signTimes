package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
	"os"
	"strings"
)

// Build a date signup attachments from parsed input
func buildAttachments(s *slack.SlashCommand) slack.Attachment {

	// parse dated between ;
	split := strings.Split(s.Text, ";")

	attachments := slack.Attachment{
		Text:     split[1],
		Fallback: "Woah dude, that don't work",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "Signup!",
				Text:  "Signup!",
				Style: "primary",
				Type:  "button",
			},
		},
	}
	return attachments
}

// body of slash handler built from nlopes/slack
func slashHandler(w http.ResponseWriter, r *http.Request) {
	verificationToken := "WSSPF87NgO5USa49IfYTVlnn"
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(verificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/signtime":

		// print user input
		fmt.Println(s.Text)

		// Initialize the msg
		params := &slack.Msg{Text: "Signup for a meeting time!"}

		// Build attachments
		attachments := buildAttachments(&s)

		fmt.Println(attachments)
		// Add attachments
		params.Attachments = []slack.Attachment{attachments}

		b, err := json.Marshal(params)
		fmt.Fprintf(os.Stdout, "%s", b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", slashHandler)
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":5038", nil)
}
