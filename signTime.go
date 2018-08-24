package main

import (
	"encoding/json"
	//	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
)

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
		params := &slack.Msg{Text: s.Text}
		b, err := json.Marshal(params)
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
