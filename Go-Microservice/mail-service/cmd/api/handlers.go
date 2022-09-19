package main

import (
	"fmt"
	"net/http"
)

func (app *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
		From    string `json:"from"`
	}

	var requestPayload mailMessage

	if err := app.readJSON(w, r, &requestPayload); err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
		From:    requestPayload.From,
	}

	if err := app.Mailer.SendSMTPMessage(msg); err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Mail sent to %s ", requestPayload.To),
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
