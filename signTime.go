package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
	"os"
	"strings"
)

func ParseInput(s *slack.SlashCommand) []string {
	split := strings.Split(s.Text, ";")
	for _, v := range split {
		fmt.Println(v)
	}
	return split
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
		fmt.Println(s.Text)
		ParseInput(&s)
		fmt.Println(s.Text)
		params := &slack.Msg{Text: "Signup for a meeting time!"}
		attachments := slack.Attachment{
			Text:     s.Text,
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
